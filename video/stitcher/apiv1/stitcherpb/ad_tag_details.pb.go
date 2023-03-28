// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/video/stitcher/v1/ad_tag_details.proto

package stitcherpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Container for a live session's ad tag detail.
type LiveAdTagDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name in the form of
	// `projects/{project}/locations/{location}/liveSessions/{live_session}/liveAdTagDetails/{id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of ad requests.
	AdRequests []*AdRequest `protobuf:"bytes,2,rep,name=ad_requests,json=adRequests,proto3" json:"ad_requests,omitempty"`
}

func (x *LiveAdTagDetail) Reset() {
	*x = LiveAdTagDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveAdTagDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveAdTagDetail) ProtoMessage() {}

func (x *LiveAdTagDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveAdTagDetail.ProtoReflect.Descriptor instead.
func (*LiveAdTagDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP(), []int{0}
}

func (x *LiveAdTagDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LiveAdTagDetail) GetAdRequests() []*AdRequest {
	if x != nil {
		return x.AdRequests
	}
	return nil
}

// Information related to the details for one ad tag.
type VodAdTagDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the ad tag detail for the specified VOD session, in the form of
	// `projects/{project}/locations/{location}/vodSessions/{vod_session_id}/vodAdTagDetails/{id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of ad requests for one ad tag.
	AdRequests []*AdRequest `protobuf:"bytes,2,rep,name=ad_requests,json=adRequests,proto3" json:"ad_requests,omitempty"`
}

func (x *VodAdTagDetail) Reset() {
	*x = VodAdTagDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodAdTagDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodAdTagDetail) ProtoMessage() {}

func (x *VodAdTagDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodAdTagDetail.ProtoReflect.Descriptor instead.
func (*VodAdTagDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP(), []int{1}
}

func (x *VodAdTagDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VodAdTagDetail) GetAdRequests() []*AdRequest {
	if x != nil {
		return x.AdRequests
	}
	return nil
}

// Details of an ad request to an ad server.
type AdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ad tag URI processed with integrated macros.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The request metadata used to make the ad request.
	RequestMetadata *RequestMetadata `protobuf:"bytes,2,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	// The response metadata received from the ad request.
	ResponseMetadata *ResponseMetadata `protobuf:"bytes,3,opt,name=response_metadata,json=responseMetadata,proto3" json:"response_metadata,omitempty"`
}

func (x *AdRequest) Reset() {
	*x = AdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdRequest) ProtoMessage() {}

func (x *AdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdRequest.ProtoReflect.Descriptor instead.
func (*AdRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP(), []int{2}
}

func (x *AdRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *AdRequest) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

func (x *AdRequest) GetResponseMetadata() *ResponseMetadata {
	if x != nil {
		return x.ResponseMetadata
	}
	return nil
}

// Metadata for an ad request.
type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP headers of the ad request.
	Headers *structpb.Struct `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP(), []int{3}
}

func (x *RequestMetadata) GetHeaders() *structpb.Struct {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Metadata for the response of an ad request.
type ResponseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error message received when making the ad request.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Headers from the response.
	Headers *structpb.Struct `protobuf:"bytes,2,opt,name=headers,proto3" json:"headers,omitempty"`
	// Status code for the response.
	StatusCode string `protobuf:"bytes,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// Size in bytes of the response.
	SizeBytes int32 `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// Total time elapsed for the response.
	Duration *durationpb.Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// The body of the response.
	Body string `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ResponseMetadata) Reset() {
	*x = ResponseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMetadata) ProtoMessage() {}

func (x *ResponseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMetadata.ProtoReflect.Descriptor instead.
func (*ResponseMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMetadata) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ResponseMetadata) GetHeaders() *structpb.Struct {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ResponseMetadata) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *ResponseMetadata) GetSizeBytes() int32 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *ResponseMetadata) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *ResponseMetadata) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_google_cloud_video_stitcher_v1_ad_tag_details_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x64, 0x54, 0x61, 0x67, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x3a, 0x9d, 0x01, 0xea, 0x41, 0x99, 0x01, 0x0a, 0x2c, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x64, 0x54, 0x61,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x69, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x69, 0x76,
	0x65, 0x41, 0x64, 0x54, 0x61, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x7d, 0x22, 0x8b, 0x02, 0x0a, 0x0e, 0x56, 0x6f, 0x64, 0x41, 0x64, 0x54, 0x61, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x64, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x98, 0x01, 0xea, 0x41, 0x94, 0x01, 0x0a, 0x2b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x64, 0x41, 0x64,
	0x54, 0x61, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x65, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x76, 0x6f, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x76, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x6f, 0x64,
	0x41, 0x64, 0x54, 0x61, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x76, 0x6f,
	0x64, 0x5f, 0x61, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x7d,
	0x22, 0xd8, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x12, 0x5a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x11,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x0f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31,
	0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x22, 0xe6, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x41, 0x64, 0x54, 0x61, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescData = file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_video_stitcher_v1_ad_tag_details_proto_goTypes = []interface{}{
	(*LiveAdTagDetail)(nil),     // 0: google.cloud.video.stitcher.v1.LiveAdTagDetail
	(*VodAdTagDetail)(nil),      // 1: google.cloud.video.stitcher.v1.VodAdTagDetail
	(*AdRequest)(nil),           // 2: google.cloud.video.stitcher.v1.AdRequest
	(*RequestMetadata)(nil),     // 3: google.cloud.video.stitcher.v1.RequestMetadata
	(*ResponseMetadata)(nil),    // 4: google.cloud.video.stitcher.v1.ResponseMetadata
	(*structpb.Struct)(nil),     // 5: google.protobuf.Struct
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_google_cloud_video_stitcher_v1_ad_tag_details_proto_depIdxs = []int32{
	2, // 0: google.cloud.video.stitcher.v1.LiveAdTagDetail.ad_requests:type_name -> google.cloud.video.stitcher.v1.AdRequest
	2, // 1: google.cloud.video.stitcher.v1.VodAdTagDetail.ad_requests:type_name -> google.cloud.video.stitcher.v1.AdRequest
	3, // 2: google.cloud.video.stitcher.v1.AdRequest.request_metadata:type_name -> google.cloud.video.stitcher.v1.RequestMetadata
	4, // 3: google.cloud.video.stitcher.v1.AdRequest.response_metadata:type_name -> google.cloud.video.stitcher.v1.ResponseMetadata
	5, // 4: google.cloud.video.stitcher.v1.RequestMetadata.headers:type_name -> google.protobuf.Struct
	5, // 5: google.cloud.video.stitcher.v1.ResponseMetadata.headers:type_name -> google.protobuf.Struct
	6, // 6: google.cloud.video.stitcher.v1.ResponseMetadata.duration:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_ad_tag_details_proto_init() }
func file_google_cloud_video_stitcher_v1_ad_tag_details_proto_init() {
	if File_google_cloud_video_stitcher_v1_ad_tag_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveAdTagDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodAdTagDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_ad_tag_details_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_ad_tag_details_proto_depIdxs,
		MessageInfos:      file_google_cloud_video_stitcher_v1_ad_tag_details_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_ad_tag_details_proto = out.File
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_ad_tag_details_proto_depIdxs = nil
}
