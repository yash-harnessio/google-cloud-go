// Copyright 2022 Google LLC
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

package storage

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newBigQueryWriteClientHook clientHook

// BigQueryWriteCallOptions contains the retry settings for each method of BigQueryWriteClient.
type BigQueryWriteCallOptions struct {
	CreateWriteStream       []gax.CallOption
	AppendRows              []gax.CallOption
	GetWriteStream          []gax.CallOption
	FinalizeWriteStream     []gax.CallOption
	BatchCommitWriteStreams []gax.CallOption
	FlushRows               []gax.CallOption
}

func defaultBigQueryWriteGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerystorage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerystorage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBigQueryWriteCallOptions() *BigQueryWriteCallOptions {
	return &BigQueryWriteCallOptions{
		CreateWriteStream: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AppendRows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetWriteStream: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		FinalizeWriteStream: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchCommitWriteStreams: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		FlushRows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalBigQueryWriteClient is an interface that defines the methods available from BigQuery Storage API.
type internalBigQueryWriteClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateWriteStream(context.Context, *storagepb.CreateWriteStreamRequest, ...gax.CallOption) (*storagepb.WriteStream, error)
	AppendRows(context.Context, ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error)
	GetWriteStream(context.Context, *storagepb.GetWriteStreamRequest, ...gax.CallOption) (*storagepb.WriteStream, error)
	FinalizeWriteStream(context.Context, *storagepb.FinalizeWriteStreamRequest, ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error)
	BatchCommitWriteStreams(context.Context, *storagepb.BatchCommitWriteStreamsRequest, ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error)
	FlushRows(context.Context, *storagepb.FlushRowsRequest, ...gax.CallOption) (*storagepb.FlushRowsResponse, error)
}

// BigQueryWriteClient is a client for interacting with BigQuery Storage API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BigQuery Write API.
//
// The Write API can be used to write data to BigQuery.
//
// For supplementary information about the Write API, see:
// https://cloud.google.com/bigquery/docs/write-api (at https://cloud.google.com/bigquery/docs/write-api)
type BigQueryWriteClient struct {
	// The internal transport-dependent client.
	internalClient internalBigQueryWriteClient

	// The call options for this service.
	CallOptions *BigQueryWriteCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BigQueryWriteClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BigQueryWriteClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BigQueryWriteClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateWriteStream creates a write stream to the given table.
// Additionally, every table has a special stream named ‘_default’
// to which data can be written. This stream doesn’t need to be created using
// CreateWriteStream. It is a stream that can be used simultaneously by any
// number of clients. Data written to this stream is considered committed as
// soon as an acknowledgement is received.
func (c *BigQueryWriteClient) CreateWriteStream(ctx context.Context, req *storagepb.CreateWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	return c.internalClient.CreateWriteStream(ctx, req, opts...)
}

// AppendRows appends data to the given stream.
//
// If offset is specified, the offset is checked against the end of
// stream. The server returns OUT_OF_RANGE in AppendRowsResponse if an
// attempt is made to append to an offset beyond the current end of the stream
// or ALREADY_EXISTS if user provides an offset that has already been
// written to. User can retry with adjusted offset within the same RPC
// connection. If offset is not specified, append happens at the end of the
// stream.
//
// The response contains an optional offset at which the append
// happened.  No offset information will be returned for appends to a
// default stream.
//
// Responses are received in the same order in which requests are sent.
// There will be one response for each successful inserted request.  Responses
// may optionally embed error information if the originating AppendRequest was
// not successfully processed.
//
// The specifics of when successfully appended data is made visible to the
// table are governed by the type of stream:
//
//   For COMMITTED streams (which includes the default stream), data is
//   visible immediately upon successful append.
//
//   For BUFFERED streams, data is made visible via a subsequent FlushRows
//   rpc which advances a cursor to a newer offset in the stream.
//
//   For PENDING streams, data is not made visible until the stream itself is
//   finalized (via the FinalizeWriteStream rpc), and the stream is explicitly
//   committed via the BatchCommitWriteStreams rpc.
//
// Note: For users coding against the gRPC api directly, it may be
// necessary to supply the x-goog-request-params system parameter
// with write_stream=<full_write_stream_name>.
//
// More information about system parameters:
// https://cloud.google.com/apis/docs/system-parameters (at https://cloud.google.com/apis/docs/system-parameters)
func (c *BigQueryWriteClient) AppendRows(ctx context.Context, opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error) {
	return c.internalClient.AppendRows(ctx, opts...)
}

// GetWriteStream gets information about a write stream.
func (c *BigQueryWriteClient) GetWriteStream(ctx context.Context, req *storagepb.GetWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	return c.internalClient.GetWriteStream(ctx, req, opts...)
}

// FinalizeWriteStream finalize a write stream so that no new data can be appended to the
// stream. Finalize is not supported on the ‘_default’ stream.
func (c *BigQueryWriteClient) FinalizeWriteStream(ctx context.Context, req *storagepb.FinalizeWriteStreamRequest, opts ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error) {
	return c.internalClient.FinalizeWriteStream(ctx, req, opts...)
}

// BatchCommitWriteStreams atomically commits a group of PENDING streams that belong to the same
// parent table.
//
// Streams must be finalized before commit and cannot be committed multiple
// times. Once a stream is committed, data in the stream becomes available
// for read operations.
func (c *BigQueryWriteClient) BatchCommitWriteStreams(ctx context.Context, req *storagepb.BatchCommitWriteStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error) {
	return c.internalClient.BatchCommitWriteStreams(ctx, req, opts...)
}

// FlushRows flushes rows to a BUFFERED stream.
//
// If users are appending rows to BUFFERED stream, flush operation is
// required in order for the rows to become available for reading. A
// Flush operation flushes up to any previously flushed offset in a BUFFERED
// stream, to the offset specified in the request.
//
// Flush is not supported on the _default stream, since it is not BUFFERED.
func (c *BigQueryWriteClient) FlushRows(ctx context.Context, req *storagepb.FlushRowsRequest, opts ...gax.CallOption) (*storagepb.FlushRowsResponse, error) {
	return c.internalClient.FlushRows(ctx, req, opts...)
}

// bigQueryWriteGRPCClient is a client for interacting with BigQuery Storage API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryWriteGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BigQueryWriteClient
	CallOptions **BigQueryWriteCallOptions

	// The gRPC API client.
	bigQueryWriteClient storagepb.BigQueryWriteClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBigQueryWriteClient creates a new big query write client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BigQuery Write API.
//
// The Write API can be used to write data to BigQuery.
//
// For supplementary information about the Write API, see:
// https://cloud.google.com/bigquery/docs/write-api (at https://cloud.google.com/bigquery/docs/write-api)
func NewBigQueryWriteClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryWriteClient, error) {
	clientOpts := defaultBigQueryWriteGRPCClientOptions()
	if newBigQueryWriteClientHook != nil {
		hookOpts, err := newBigQueryWriteClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BigQueryWriteClient{CallOptions: defaultBigQueryWriteCallOptions()}

	c := &bigQueryWriteGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		bigQueryWriteClient: storagepb.NewBigQueryWriteClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *bigQueryWriteGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryWriteGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryWriteGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *bigQueryWriteGRPCClient) CreateWriteStream(ctx context.Context, req *storagepb.CreateWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateWriteStream[0:len((*c.CallOptions).CreateWriteStream):len((*c.CallOptions).CreateWriteStream)], opts...)
	var resp *storagepb.WriteStream
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.CreateWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) AppendRows(ctx context.Context, opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp storagepb.BigQueryWrite_AppendRowsClient
	opts = append((*c.CallOptions).AppendRows[0:len((*c.CallOptions).AppendRows):len((*c.CallOptions).AppendRows)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.AppendRows(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) GetWriteStream(ctx context.Context, req *storagepb.GetWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetWriteStream[0:len((*c.CallOptions).GetWriteStream):len((*c.CallOptions).GetWriteStream)], opts...)
	var resp *storagepb.WriteStream
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.GetWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) FinalizeWriteStream(ctx context.Context, req *storagepb.FinalizeWriteStreamRequest, opts ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FinalizeWriteStream[0:len((*c.CallOptions).FinalizeWriteStream):len((*c.CallOptions).FinalizeWriteStream)], opts...)
	var resp *storagepb.FinalizeWriteStreamResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.FinalizeWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) BatchCommitWriteStreams(ctx context.Context, req *storagepb.BatchCommitWriteStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCommitWriteStreams[0:len((*c.CallOptions).BatchCommitWriteStreams):len((*c.CallOptions).BatchCommitWriteStreams)], opts...)
	var resp *storagepb.BatchCommitWriteStreamsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.BatchCommitWriteStreams(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) FlushRows(ctx context.Context, req *storagepb.FlushRowsRequest, opts ...gax.CallOption) (*storagepb.FlushRowsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "write_stream", url.QueryEscape(req.GetWriteStream())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FlushRows[0:len((*c.CallOptions).FlushRows):len((*c.CallOptions).FlushRows)], opts...)
	var resp *storagepb.FlushRowsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.FlushRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
