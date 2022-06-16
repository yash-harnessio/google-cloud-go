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

package dataflow

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newJobsV1Beta3ClientHook clientHook

// JobsV1Beta3CallOptions contains the retry settings for each method of JobsV1Beta3Client.
type JobsV1Beta3CallOptions struct {
	CreateJob          []gax.CallOption
	GetJob             []gax.CallOption
	UpdateJob          []gax.CallOption
	ListJobs           []gax.CallOption
	AggregatedListJobs []gax.CallOption
	CheckActiveJobs    []gax.CallOption
	SnapshotJob        []gax.CallOption
}

func defaultJobsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultJobsV1Beta3CallOptions() *JobsV1Beta3CallOptions {
	return &JobsV1Beta3CallOptions{
		CreateJob:          []gax.CallOption{},
		GetJob:             []gax.CallOption{},
		UpdateJob:          []gax.CallOption{},
		ListJobs:           []gax.CallOption{},
		AggregatedListJobs: []gax.CallOption{},
		CheckActiveJobs:    []gax.CallOption{},
		SnapshotJob:        []gax.CallOption{},
	}
}

// internalJobsV1Beta3Client is an interface that defines the methods available from Dataflow API.
type internalJobsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateJob(context.Context, *dataflowpb.CreateJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	GetJob(context.Context, *dataflowpb.GetJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	UpdateJob(context.Context, *dataflowpb.UpdateJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	ListJobs(context.Context, *dataflowpb.ListJobsRequest, ...gax.CallOption) *JobIterator
	AggregatedListJobs(context.Context, *dataflowpb.ListJobsRequest, ...gax.CallOption) *JobIterator
	CheckActiveJobs(context.Context, *dataflowpb.CheckActiveJobsRequest, ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error)
	SnapshotJob(context.Context, *dataflowpb.SnapshotJobRequest, ...gax.CallOption) (*dataflowpb.Snapshot, error)
}

// JobsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides a method to create and modify Google Cloud Dataflow jobs.
// A Job is a multi-stage computation graph run by the Cloud Dataflow service.
type JobsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalJobsV1Beta3Client

	// The call options for this service.
	CallOptions *JobsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *JobsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *JobsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *JobsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateJob creates a Cloud Dataflow job.
//
// To create a job, we recommend using projects.locations.jobs.create with a
// [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.create is not recommended, as your job will always start
// in us-central1.
func (c *JobsV1Beta3Client) CreateJob(ctx context.Context, req *dataflowpb.CreateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.CreateJob(ctx, req, opts...)
}

// GetJob gets the state of the specified Cloud Dataflow job.
//
// To get the state of a job, we recommend using projects.locations.jobs.get
// with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.get is not recommended, as you can only get the state of
// jobs that are running in us-central1.
func (c *JobsV1Beta3Client) GetJob(ctx context.Context, req *dataflowpb.GetJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.GetJob(ctx, req, opts...)
}

// UpdateJob updates the state of an existing Cloud Dataflow job.
//
// To update the state of an existing job, we recommend using
// projects.locations.jobs.update with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.update is not recommended, as you can only update the state
// of jobs that are running in us-central1.
func (c *JobsV1Beta3Client) UpdateJob(ctx context.Context, req *dataflowpb.UpdateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.UpdateJob(ctx, req, opts...)
}

// ListJobs list the jobs of a project.
//
// To list the jobs of a project in a region, we recommend using
// projects.locations.jobs.list with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). To
// list the all jobs across all regions, use projects.jobs.aggregated. Using
// projects.jobs.list is not recommended, as you can only get the list of
// jobs that are running in us-central1.
func (c *JobsV1Beta3Client) ListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	return c.internalClient.ListJobs(ctx, req, opts...)
}

// AggregatedListJobs list the jobs of a project across all regions.
func (c *JobsV1Beta3Client) AggregatedListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	return c.internalClient.AggregatedListJobs(ctx, req, opts...)
}

// CheckActiveJobs check for existence of active jobs in the given project across all regions.
func (c *JobsV1Beta3Client) CheckActiveJobs(ctx context.Context, req *dataflowpb.CheckActiveJobsRequest, opts ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error) {
	return c.internalClient.CheckActiveJobs(ctx, req, opts...)
}

// SnapshotJob snapshot the state of a streaming job.
func (c *JobsV1Beta3Client) SnapshotJob(ctx context.Context, req *dataflowpb.SnapshotJobRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	return c.internalClient.SnapshotJob(ctx, req, opts...)
}

// jobsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type jobsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing JobsV1Beta3Client
	CallOptions **JobsV1Beta3CallOptions

	// The gRPC API client.
	jobsV1Beta3Client dataflowpb.JobsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewJobsV1Beta3Client creates a new jobs v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides a method to create and modify Google Cloud Dataflow jobs.
// A Job is a multi-stage computation graph run by the Cloud Dataflow service.
func NewJobsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*JobsV1Beta3Client, error) {
	clientOpts := defaultJobsV1Beta3GRPCClientOptions()
	if newJobsV1Beta3ClientHook != nil {
		hookOpts, err := newJobsV1Beta3ClientHook(ctx, clientHookParams{})
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
	client := JobsV1Beta3Client{CallOptions: defaultJobsV1Beta3CallOptions()}

	c := &jobsV1Beta3GRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		jobsV1Beta3Client: dataflowpb.NewJobsV1Beta3Client(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *jobsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *jobsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *jobsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *jobsV1Beta3GRPCClient) CreateJob(ctx context.Context, req *dataflowpb.CreateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateJob[0:len((*c.CallOptions).CreateJob):len((*c.CallOptions).CreateJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.CreateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) GetJob(ctx context.Context, req *dataflowpb.GetJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetJob[0:len((*c.CallOptions).GetJob):len((*c.CallOptions).GetJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) UpdateJob(ctx context.Context, req *dataflowpb.UpdateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateJob[0:len((*c.CallOptions).UpdateJob):len((*c.CallOptions).UpdateJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) ListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListJobs[0:len((*c.CallOptions).ListJobs):len((*c.CallOptions).ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
			resp, err = c.jobsV1Beta3Client.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

func (c *jobsV1Beta3GRPCClient) AggregatedListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AggregatedListJobs[0:len((*c.CallOptions).AggregatedListJobs):len((*c.CallOptions).AggregatedListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
			resp, err = c.jobsV1Beta3Client.AggregatedListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

func (c *jobsV1Beta3GRPCClient) CheckActiveJobs(ctx context.Context, req *dataflowpb.CheckActiveJobsRequest, opts ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CheckActiveJobs[0:len((*c.CallOptions).CheckActiveJobs):len((*c.CallOptions).CheckActiveJobs)], opts...)
	var resp *dataflowpb.CheckActiveJobsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.CheckActiveJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) SnapshotJob(ctx context.Context, req *dataflowpb.SnapshotJobRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SnapshotJob[0:len((*c.CallOptions).SnapshotJob):len((*c.CallOptions).SnapshotJob)], opts...)
	var resp *dataflowpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.SnapshotJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// JobIterator manages a stream of *dataflowpb.Job.
type JobIterator struct {
	items    []*dataflowpb.Job
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*dataflowpb.Job, error) {
	var item *dataflowpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
