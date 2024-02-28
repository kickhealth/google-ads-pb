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

var newCampaignAssetClientHook clientHook

// CampaignAssetCallOptions contains the retry settings for each method of CampaignAssetClient.
type CampaignAssetCallOptions struct {
	MutateCampaignAssets []gax.CallOption
}

func defaultCampaignAssetGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignAssetCallOptions() *CampaignAssetCallOptions {
	return &CampaignAssetCallOptions{
		MutateCampaignAssets: []gax.CallOption{
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

// internalCampaignAssetClient is an interface that defines the methods available from Google Ads API.
type internalCampaignAssetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignAssets(context.Context, *servicespb.MutateCampaignAssetsRequest, ...gax.CallOption) (*servicespb.MutateCampaignAssetsResponse, error)
}

// CampaignAssetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign assets.
type CampaignAssetClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignAssetClient

	// The call options for this service.
	CallOptions *CampaignAssetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignAssetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignAssetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignAssetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignAssets creates, updates, or removes campaign assets. Operation statuses are
// returned.
//
// List of thrown errors:
// AssetLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// ContextError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// NotAllowlistedError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignAssetClient) MutateCampaignAssets(ctx context.Context, req *servicespb.MutateCampaignAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignAssetsResponse, error) {
	return c.internalClient.MutateCampaignAssets(ctx, req, opts...)
}

// campaignAssetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignAssetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignAssetClient
	CallOptions **CampaignAssetCallOptions

	// The gRPC API client.
	campaignAssetClient servicespb.CampaignAssetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCampaignAssetClient creates a new campaign asset service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign assets.
func NewCampaignAssetClient(ctx context.Context, opts ...option.ClientOption) (*CampaignAssetClient, error) {
	clientOpts := defaultCampaignAssetGRPCClientOptions()
	if newCampaignAssetClientHook != nil {
		hookOpts, err := newCampaignAssetClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignAssetClient{CallOptions: defaultCampaignAssetCallOptions()}

	c := &campaignAssetGRPCClient{
		connPool:            connPool,
		campaignAssetClient: servicespb.NewCampaignAssetServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *campaignAssetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignAssetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignAssetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignAssetGRPCClient) MutateCampaignAssets(ctx context.Context, req *servicespb.MutateCampaignAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignAssetsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignAssets[0:len((*c.CallOptions).MutateCampaignAssets):len((*c.CallOptions).MutateCampaignAssets)], opts...)
	var resp *servicespb.MutateCampaignAssetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignAssetClient.MutateCampaignAssets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
