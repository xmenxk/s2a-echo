// Package echo provides libraries to ping the Bifrost Echo service.
package main

import (
	"context"
	"github.com/google/s2a-go"
	"google.golang.org/appengine"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/proto"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

const (
	echoServiceAddress = "dns:///bifrost-echo-c2p.mtls.googleapis.com:443"
	mdsAddress         = "metadata.google.internal:80"

	connectionTimeout = time.Duration(10) * time.Second

	// appEngineRootCert MUST be kept in sync with the "google-appengine-root-ca"
	// CA certificate in the "zatar-private-cas" GCP project:
	//  https://pantheon.corp.google.com/home/dashboard?project=zatar-private-cas
	appEngineRootCert = "-----BEGIN CERTIFICATE-----\n" +
		"MIIB2TCCAX+gAwIBAgIUAPEHXaekBcPkuxZuQBXrL/KUoeEwCgYIKoZIzj0EAwIw\n" +
		"OjEVMBMGA1UEChMMR29vZ2xlIENsb3VkMSEwHwYDVQQDExhHb29nbGUgQXBwRW5n\n" +
		"aW5lIFJvb3QgQ0EwHhcNMjEwODEyMDQzMDE5WhcNMzEwODEyMTQzNzU5WjA6MRUw\n" +
		"EwYDVQQKEwxHb29nbGUgQ2xvdWQxITAfBgNVBAMTGEdvb2dsZSBBcHBFbmdpbmUg\n" +
		"Um9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIHBqSc2YtCSmlG7ZLGm\n" +
		"KPUyabAjh+xsZj+Hbts7Zgu4sgf53PevAvAiJw6JPndlg692z4lDK4rfqOb71eLL\n" +
		"B5GjYzBhMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQW\n" +
		"BBRNCEC2JLHR5Q1HzEq104zgkZ2MNzAfBgNVHSMEGDAWgBRNCEC2JLHR5Q1HzEq1\n" +
		"04zgkZ2MNzAKBggqhkjOPQQDAgNIADBFAiEAxxALPFvTgyYyhW0jc/nE07gOzdR4\n" +
		"VyDXFH67f0lvhP0CIDlCynvmNQa4sIK8m+9DfEqD3OGNEOSX6ILQEGJHRWHb\n" +
		"-----END CERTIFICATE-----"
)

// DoEcho establishes an mTLS connection to the Bifrost Echo service using
// S2A-in-GAE, and pings the service.
func DoEcho(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	queryGoogleApis(ctx, mdsAddress, false /*returnAfterConnectionAttempt*/, w, r, false /*enableV2=*/)
}

// DoEchoV2 is exactly like DoEcho, except using S2Av2-in-GAE.
func DoEchoV2(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	queryGoogleApis(ctx, mdsAddress, false /*returnAfterConnectionAttempt*/, w, r, true /*enableV2*/)
}

func queryGoogleApis(ctx context.Context, s2aAddress string, returnAfterConnectionAttempt bool, w http.ResponseWriter, r *http.Request, enableV2 bool) {

	// Set up the client-side S2A transport credentials.
	var ensureProcessSessionTickets sync.WaitGroup
	var clientOpts *s2a.ClientOptions
	if enableV2 {
		clientOpts = &s2a.ClientOptions{
			S2AAddress: s2aAddress,
			EnableV2:   enableV2,
		}
	} else {
		clientOpts = &s2a.ClientOptions{
			S2AAddress:                  s2aAddress,
			EnsureProcessSessionTickets: &ensureProcessSessionTickets,
			EnableV2:                    enableV2,
		}
	}
	creds, err := s2a.NewClientCreds(clientOpts)
	if err != nil {
		grpclog.Errorf("Failed to create s2a client credentials: %v", err)
		http.Error(w, "Failed to create s2a client credentials: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set up a connection to the server, using S2A-in-GAE for authentication.
	perRPCCreds, err := oauth.NewApplicationDefault(ctx,
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/bifrost")
	if err != nil {
		grpclog.Errorf("Failed to get per-RPC credentials: %v", err)
		http.Error(w, "Failed to get per-RPC credentials: "+err.Error(), http.StatusBadRequest)
		return
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRPCCreds),
		grpc.WithReturnConnectionError(),
		grpc.WithBlock(),
	}
	ctx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, echoServiceAddress, opts...)
	if err != nil {
		grpclog.Errorf("Failed to connect to Bifrost Echo service: %v", err)
		http.Error(w, "Failed to connect to Bifrost Echo service: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if !enableV2 {
			ensureProcessSessionTickets.Wait()
		}
		conn.Close()
	}()
	w.Write([]byte("Successfully connected!"))
	if returnAfterConnectionAttempt {
		return
	}

	// Query the Bifrost Echo service.
	client := NewProxyClient(conn)
	ctx, cancel = context.WithTimeout(ctx, connectionTimeout)
	defer cancel()
	response, err := client.Info(ctx, &InfoRequest{Request: proto.String("echo")})
	if err != nil {
		grpclog.Errorf("Error while calling Bifrost Echo service: %v", err)
		http.Error(w, "Error while calling Bifrost Echo service: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Return the response to the caller.
	grpclog.Infof("Received response from Bifrost Echo service: %q", response)
	buffer, err := proto.Marshal(response)
	if err != nil {
		grpclog.Errorf("Failed to marshal response: %v", err)
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buffer)
	w.Write([]byte("\n"))
}
