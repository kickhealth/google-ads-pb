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
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newCustomAudienceClientHook clientHook

// CustomAudienceCallOptions contains the retry settings for each method of CustomAudienceClient.
type CustomAudienceCallOptions struct {
	MutateCustomAudiences []gax.CallOption
}

func defaultCustomAudienceGRPCClientOptions() []option.ClientOption {
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

func defaultCustomAudienceCallOptions() *CustomAudienceCallOptions {
	return &CustomAudienceCallOptions{
		MutateCustomAudiences: []gax.CallOption{
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

// internalCustomAudienceClient is an interface that defines the methods available from Google Ads API.
type internalCustomAudienceClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomAudiences(context.Context, *servicespb.MutateCustomAudiencesRequest, ...gax.CallOption) (*servicespb.MutateCustomAudiencesResponse, error)
}

// CustomAudienceClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage custom audiences.
type CustomAudienceClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomAudienceClient

	// The call options for this service.
	CallOptions *CustomAudienceCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomAudienceClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomAudienceClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CustomAudienceClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomAudiences creates or updates custom audiences. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CustomAudienceError (at )
// CustomInterestError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// OperationAccessDeniedError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomAudienceClient) MutateCustomAudiences(ctx context.Context, req *servicespb.MutateCustomAudiencesRequest, opts ...gax.CallOption) (*servicespb.MutateCustomAudiencesResponse, error) {
	return c.internalClient.MutateCustomAudiences(ctx, req, opts...)
}

// customAudienceGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customAudienceGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CustomAudienceClient
	CallOptions **CustomAudienceCallOptions

	// The gRPC API client.
	customAudienceClient servicespb.CustomAudienceServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCustomAudienceClient creates a new custom audience service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage custom audiences.
func NewCustomAudienceClient(ctx context.Context, opts ...option.ClientOption) (*CustomAudienceClient, error) {
	clientOpts := defaultCustomAudienceGRPCClientOptions()
	if newCustomAudienceClientHook != nil {
		hookOpts, err := newCustomAudienceClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomAudienceClient{CallOptions: defaultCustomAudienceCallOptions()}

	c := &customAudienceGRPCClient{
		connPool:             connPool,
		customAudienceClient: servicespb.NewCustomAudienceServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *customAudienceGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customAudienceGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customAudienceGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customAudienceGRPCClient) MutateCustomAudiences(ctx context.Context, req *servicespb.MutateCustomAudiencesRequest, opts ...gax.CallOption) (*servicespb.MutateCustomAudiencesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCustomAudiences[0:len((*c.CallOptions).MutateCustomAudiences):len((*c.CallOptions).MutateCustomAudiences)], opts...)
	var resp *servicespb.MutateCustomAudiencesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customAudienceClient.MutateCustomAudiences(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
