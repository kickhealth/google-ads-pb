// Copyright 2023 Google LLC
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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newGoogleAdsClientHook clientHook

// GoogleAdsCallOptions contains the retry settings for each method of GoogleAdsClient.
type GoogleAdsCallOptions struct {
	Search       []gax.CallOption
	SearchStream []gax.CallOption
	Mutate       []gax.CallOption
}

func defaultGoogleAdsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultGoogleAdsCallOptions() *GoogleAdsCallOptions {
	return &GoogleAdsCallOptions{
		Search: []gax.CallOption{
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
		SearchStream: []gax.CallOption{
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
		Mutate: []gax.CallOption{
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

// internalGoogleAdsClient is an interface that defines the methods available from Google Ads API.
type internalGoogleAdsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Search(context.Context, *servicespb.SearchGoogleAdsRequest, ...gax.CallOption) *GoogleAdsRowIterator
	SearchStream(context.Context, *servicespb.SearchGoogleAdsStreamRequest, ...gax.CallOption) (servicespb.GoogleAdsService_SearchStreamClient, error)
	Mutate(context.Context, *servicespb.MutateGoogleAdsRequest, ...gax.CallOption) (*servicespb.MutateGoogleAdsResponse, error)
}

// GoogleAdsClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch data and metrics across resources.
type GoogleAdsClient struct {
	// The internal transport-dependent client.
	internalClient internalGoogleAdsClient

	// The call options for this service.
	CallOptions *GoogleAdsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GoogleAdsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GoogleAdsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *GoogleAdsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Search returns all rows that match the search query.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ChangeEventError (at )
// ChangeStatusError (at )
// ClickViewError (at )
// HeaderError (at )
// InternalError (at )
// QueryError (at )
// QuotaError (at )
// RequestError (at )
func (c *GoogleAdsClient) Search(ctx context.Context, req *servicespb.SearchGoogleAdsRequest, opts ...gax.CallOption) *GoogleAdsRowIterator {
	return c.internalClient.Search(ctx, req, opts...)
}

// SearchStream returns all rows that match the search stream query.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ChangeEventError (at )
// ChangeStatusError (at )
// ClickViewError (at )
// HeaderError (at )
// InternalError (at )
// QueryError (at )
// QuotaError (at )
// RequestError (at )
func (c *GoogleAdsClient) SearchStream(ctx context.Context, req *servicespb.SearchGoogleAdsStreamRequest, opts ...gax.CallOption) (servicespb.GoogleAdsService_SearchStreamClient, error) {
	return c.internalClient.SearchStream(ctx, req, opts...)
}

// Mutate creates, updates, or removes resources. This method supports atomic
// transactions with multiple types of resources. For example, you can
// atomically create a campaign and a campaign budget, or perform up to
// thousands of mutates atomically.
//
// This method is essentially a wrapper around a series of mutate methods. The
// only features it offers over calling those methods directly are:
//
//	Atomic transactions
//
//	Temp resource names (described below)
//
//	Somewhat reduced latency over making a series of mutate calls
//
// Note: Only resources that support atomic transactions are included, so this
// method can’t replace all calls to individual services.
//
// Atomic Transaction BenefitsAtomicity makes error handling much easier. If you’re making a series of
// changes and one fails, it can leave your account in an inconsistent state.
// With atomicity, you either reach the chosen state directly, or the request
// fails and you can retry.
//
// Temp Resource NamesTemp resource names are a special type of resource name used to create a
// resource and reference that resource in the same request. For example, if a
// campaign budget is created with resource_name equal to
// customers/123/campaignBudgets/-1, that resource name can be reused in
// the Campaign.budget field in the same request. That way, the two
// resources are created and linked atomically.
//
// To create a temp resource name, put a negative number in the part of the
// name that the server would normally allocate.
//
// Note:
//
//	Resources must be created with a temp name before the name can be reused.
//	For example, the previous CampaignBudget+Campaign example would fail if
//	the mutate order was reversed.
//
//	Temp names are not remembered across requests.
//
//	There’s no limit to the number of temp names in a request.
//
//	Each temp name must use a unique negative number, even if the resource
//	types differ.
//
// LatencyIt’s important to group mutates by resource type or the request may time
// out and fail. Latency is roughly equal to a series of calls to individual
// mutate methods, where each change in resource type is a new call. For
// example, mutating 10 campaigns then 10 ad groups is like 2 calls, while
// mutating 1 campaign, 1 ad group, 1 campaign, 1 ad group is like 4 calls.
//
// List of thrown errors:
// AdCustomizerError (at )
// AdError (at )
// AdGroupAdError (at )
// AdGroupCriterionError (at )
// AdGroupError (at )
// AssetError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// BiddingError (at )
// CampaignBudgetError (at )
// CampaignCriterionError (at )
// CampaignError (at )
// CampaignExperimentError (at )
// CampaignSharedSetError (at )
// CollectionSizeError (at )
// ContextError (at )
// ConversionActionError (at )
// CriterionError (at )
// CustomerFeedError (at )
// DatabaseError (at )
// DateError (at )
// DateRangeError (at )
// DistinctError (at )
// ExtensionFeedItemError (at )
// ExtensionSettingError (at )
// FeedAttributeReferenceError (at )
// FeedError (at )
// FeedItemError (at )
// FeedItemSetError (at )
// FieldError (at )
// FieldMaskError (at )
// FunctionParsingError (at )
// HeaderError (at )
// ImageError (at )
// InternalError (at )
// KeywordPlanAdGroupKeywordError (at )
// KeywordPlanCampaignError (at )
// KeywordPlanError (at )
// LabelError (at )
// ListOperationError (at )
// MediaUploadError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// PolicyFindingError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SettingError (at )
// SharedSetError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
// UserListError (at )
// YoutubeVideoRegistrationError (at )
func (c *GoogleAdsClient) Mutate(ctx context.Context, req *servicespb.MutateGoogleAdsRequest, opts ...gax.CallOption) (*servicespb.MutateGoogleAdsResponse, error) {
	return c.internalClient.Mutate(ctx, req, opts...)
}

// googleAdsGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type googleAdsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing GoogleAdsClient
	CallOptions **GoogleAdsCallOptions

	// The gRPC API client.
	googleAdsClient servicespb.GoogleAdsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewGoogleAdsClient creates a new google ads service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch data and metrics across resources.
func NewGoogleAdsClient(ctx context.Context, opts ...option.ClientOption) (*GoogleAdsClient, error) {
	clientOpts := defaultGoogleAdsGRPCClientOptions()
	if newGoogleAdsClientHook != nil {
		hookOpts, err := newGoogleAdsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := GoogleAdsClient{CallOptions: defaultGoogleAdsCallOptions()}

	c := &googleAdsGRPCClient{
		connPool:        connPool,
		googleAdsClient: servicespb.NewGoogleAdsServiceClient(connPool),
		CallOptions:     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *googleAdsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *googleAdsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *googleAdsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *googleAdsGRPCClient) Search(ctx context.Context, req *servicespb.SearchGoogleAdsRequest, opts ...gax.CallOption) *GoogleAdsRowIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Search[0:len((*c.CallOptions).Search):len((*c.CallOptions).Search)], opts...)
	it := &GoogleAdsRowIterator{}
	req = proto.Clone(req).(*servicespb.SearchGoogleAdsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*servicespb.GoogleAdsRow, string, error) {
		resp := &servicespb.SearchGoogleAdsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.googleAdsClient.Search(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *googleAdsGRPCClient) SearchStream(ctx context.Context, req *servicespb.SearchGoogleAdsStreamRequest, opts ...gax.CallOption) (servicespb.GoogleAdsService_SearchStreamClient, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SearchStream[0:len((*c.CallOptions).SearchStream):len((*c.CallOptions).SearchStream)], opts...)
	var resp servicespb.GoogleAdsService_SearchStreamClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.googleAdsClient.SearchStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *googleAdsGRPCClient) Mutate(ctx context.Context, req *servicespb.MutateGoogleAdsRequest, opts ...gax.CallOption) (*servicespb.MutateGoogleAdsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Mutate[0:len((*c.CallOptions).Mutate):len((*c.CallOptions).Mutate)], opts...)
	var resp *servicespb.MutateGoogleAdsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.googleAdsClient.Mutate(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
