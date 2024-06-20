// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package clients

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	servicespb "github.com/kickhealth/google-ads-pb/services"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newCustomerUserAccessClientHook clientHook

// CustomerUserAccessCallOptions contains the retry settings for each method of CustomerUserAccessClient.
type CustomerUserAccessCallOptions struct {
	MutateCustomerUserAccess []gax.CallOption
}

func defaultCustomerUserAccessGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("googleads.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCustomerUserAccessCallOptions() *CustomerUserAccessCallOptions {
	return &CustomerUserAccessCallOptions{
		MutateCustomerUserAccess: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalCustomerUserAccessClient is an interface that defines the methods available from Google Ads API.
type internalCustomerUserAccessClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomerUserAccess(context.Context, *servicespb.MutateCustomerUserAccessRequest, ...gax.CallOption) (*servicespb.MutateCustomerUserAccessResponse, error)
}

// CustomerUserAccessClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service manages the permissions of a user on a given customer.
type CustomerUserAccessClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerUserAccessClient

	// The call options for this service.
	CallOptions *CustomerUserAccessCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerUserAccessClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerUserAccessClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CustomerUserAccessClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomerUserAccess updates, removes permission of a user on a given customer. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CustomerUserAccessError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerUserAccessClient) MutateCustomerUserAccess(ctx context.Context, req *servicespb.MutateCustomerUserAccessRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerUserAccessResponse, error) {
	return c.internalClient.MutateCustomerUserAccess(ctx, req, opts...)
}

// customerUserAccessGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerUserAccessGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CustomerUserAccessClient
	CallOptions **CustomerUserAccessCallOptions

	// The gRPC API client.
	customerUserAccessClient servicespb.CustomerUserAccessServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCustomerUserAccessClient creates a new customer user access service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service manages the permissions of a user on a given customer.
func NewCustomerUserAccessClient(ctx context.Context, opts ...option.ClientOption) (*CustomerUserAccessClient, error) {
	clientOpts := defaultCustomerUserAccessGRPCClientOptions()
	if newCustomerUserAccessClientHook != nil {
		hookOpts, err := newCustomerUserAccessClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomerUserAccessClient{CallOptions: defaultCustomerUserAccessCallOptions()}

	c := &customerUserAccessGRPCClient{
		connPool:                 connPool,
		customerUserAccessClient: servicespb.NewCustomerUserAccessServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *customerUserAccessGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerUserAccessGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerUserAccessGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerUserAccessGRPCClient) MutateCustomerUserAccess(ctx context.Context, req *servicespb.MutateCustomerUserAccessRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerUserAccessResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCustomerUserAccess[0:len((*c.CallOptions).MutateCustomerUserAccess):len((*c.CallOptions).MutateCustomerUserAccess)], opts...)
	var resp *servicespb.MutateCustomerUserAccessResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerUserAccessClient.MutateCustomerUserAccess(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
