//
//     Copyright 2022 The Dragonfly Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pkg/apis/dfdaemon/v1/dfdaemon.proto

package dfdaemon

import (
	v1 "d7y.io/api/v2/pkg/apis/common/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identify one downloading, the framework will fill it automatically.
	// Deprecated
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Download file from the url, not only for http.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Pieces will be written to output path directly,
	// at the same time, dfdaemon workspace also makes soft link to the output.
	Output string `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	// Timeout duration.
	Timeout uint64 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Rate limit in bytes per second.
	Limit float64 `protobuf:"fixed64,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// Disable back-to-source.
	DisableBackSource bool `protobuf:"varint,6,opt,name=disable_back_source,json=disableBackSource,proto3" json:"disable_back_source,omitempty"`
	// URL meta info.
	UrlMeta *v1.UrlMeta `protobuf:"bytes,7,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
	// User id.
	Uid int64 `protobuf:"varint,10,opt,name=uid,proto3" json:"uid,omitempty"`
	// Group id.
	Gid int64 `protobuf:"varint,11,opt,name=gid,proto3" json:"gid,omitempty"`
	// Keep original offset, used for ranged request, only available for hard link, otherwise will failed.
	KeepOriginalOffset bool `protobuf:"varint,12,opt,name=keep_original_offset,json=keepOriginalOffset,proto3" json:"keep_original_offset,omitempty"`
	// Recursive download, when enabled, daemon will call resource list api to list and then recursive download each object.
	Recursive bool `protobuf:"varint,13,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *DownRequest) Reset() {
	*x = DownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownRequest) ProtoMessage() {}

func (x *DownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownRequest.ProtoReflect.Descriptor instead.
func (*DownRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{0}
}

func (x *DownRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DownRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DownRequest) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *DownRequest) GetTimeout() uint64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *DownRequest) GetLimit() float64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *DownRequest) GetDisableBackSource() bool {
	if x != nil {
		return x.DisableBackSource
	}
	return false
}

func (x *DownRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

func (x *DownRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DownRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *DownRequest) GetKeepOriginalOffset() bool {
	if x != nil {
		return x.KeepOriginalOffset
	}
	return false
}

func (x *DownRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type DownResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task id.
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Peer id.
	PeerId string `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// Task has completed length.
	CompletedLength uint64 `protobuf:"varint,4,opt,name=completed_length,json=completedLength,proto3" json:"completed_length,omitempty"`
	// Task has been completed.
	Done bool `protobuf:"varint,5,opt,name=done,proto3" json:"done,omitempty"`
	// Task output. Used in recursive download.
	Output string `protobuf:"bytes,6,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *DownResult) Reset() {
	*x = DownResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownResult) ProtoMessage() {}

func (x *DownResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownResult.ProtoReflect.Descriptor instead.
func (*DownResult) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{1}
}

func (x *DownResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *DownResult) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *DownResult) GetCompletedLength() uint64 {
	if x != nil {
		return x.CompletedLength
	}
	return 0
}

func (x *DownResult) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *DownResult) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type StatTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Download url.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// URL meta info.
	UrlMeta *v1.UrlMeta `protobuf:"bytes,2,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
	// Check local cache only.
	LocalOnly bool `protobuf:"varint,3,opt,name=local_only,json=localOnly,proto3" json:"local_only,omitempty"`
}

func (x *StatTaskRequest) Reset() {
	*x = StatTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatTaskRequest) ProtoMessage() {}

func (x *StatTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatTaskRequest.ProtoReflect.Descriptor instead.
func (*StatTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{2}
}

func (x *StatTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StatTaskRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

func (x *StatTaskRequest) GetLocalOnly() bool {
	if x != nil {
		return x.LocalOnly
	}
	return false
}

type ImportTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Download url.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// URL meta info.
	UrlMeta *v1.UrlMeta `protobuf:"bytes,2,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
	// File to be imported.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// Task type.
	Type v1.TaskType `protobuf:"varint,4,opt,name=type,proto3,enum=common.TaskType" json:"type,omitempty"`
}

func (x *ImportTaskRequest) Reset() {
	*x = ImportTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportTaskRequest) ProtoMessage() {}

func (x *ImportTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportTaskRequest.ProtoReflect.Descriptor instead.
func (*ImportTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{3}
}

func (x *ImportTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImportTaskRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

func (x *ImportTaskRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ImportTaskRequest) GetType() v1.TaskType {
	if x != nil {
		return x.Type
	}
	return v1.TaskType(0)
}

type ExportTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Download url.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Output path of downloaded file.
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	// Timeout duration.
	Timeout uint64 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Rate limit in bytes per second.
	Limit float64 `protobuf:"fixed64,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// URL meta info.
	UrlMeta *v1.UrlMeta `protobuf:"bytes,5,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
	// User id.
	Uid int64 `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
	// Group id.
	Gid int64 `protobuf:"varint,8,opt,name=gid,proto3" json:"gid,omitempty"`
	// Only export from local storage.
	LocalOnly bool `protobuf:"varint,9,opt,name=local_only,json=localOnly,proto3" json:"local_only,omitempty"`
}

func (x *ExportTaskRequest) Reset() {
	*x = ExportTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTaskRequest) ProtoMessage() {}

func (x *ExportTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTaskRequest.ProtoReflect.Descriptor instead.
func (*ExportTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{4}
}

func (x *ExportTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ExportTaskRequest) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ExportTaskRequest) GetTimeout() uint64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *ExportTaskRequest) GetLimit() float64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ExportTaskRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

func (x *ExportTaskRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ExportTaskRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *ExportTaskRequest) GetLocalOnly() bool {
	if x != nil {
		return x.LocalOnly
	}
	return false
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Download url.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// URL meta info.
	UrlMeta *v1.UrlMeta `protobuf:"bytes,2,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DeleteTaskRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

var File_pkg_apis_dfdaemon_v1_dfdaemon_proto protoreflect.FileDescriptor

var file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x66, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x1a,
	0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01,
	0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28,
	0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09,
	0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x72, 0x6c, 0x4d,
	0x65, 0x74, 0x61, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x75,
	0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6b, 0x65,
	0x65, 0x70, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00,
	0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x81, 0x01,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x08,
	0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x6e, 0x6c,
	0x79, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x72,
	0x6c, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x75, 0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x11,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x24, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x55, 0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x22,
	0x64, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x34, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x72, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x75, 0x72,
	0x6c, 0x4d, 0x65, 0x74, 0x61, 0x32, 0xcc, 0x04, 0x0a, 0x06, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x64,
	0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0e, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x2e, 0x64, 0x66,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41,
	0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x64,
	0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x41, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x1b, 0x2e, 0x64, 0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x64, 0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x64, 0x37, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64,
	0x66, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x66, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescOnce sync.Once
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescData = file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDesc
)

func file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescGZIP() []byte {
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescOnce.Do(func() {
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescData)
	})
	return file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDescData
}

var file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_apis_dfdaemon_v1_dfdaemon_proto_goTypes = []interface{}{
	(*DownRequest)(nil),         // 0: dfdaemon.DownRequest
	(*DownResult)(nil),          // 1: dfdaemon.DownResult
	(*StatTaskRequest)(nil),     // 2: dfdaemon.StatTaskRequest
	(*ImportTaskRequest)(nil),   // 3: dfdaemon.ImportTaskRequest
	(*ExportTaskRequest)(nil),   // 4: dfdaemon.ExportTaskRequest
	(*DeleteTaskRequest)(nil),   // 5: dfdaemon.DeleteTaskRequest
	(*v1.UrlMeta)(nil),          // 6: common.UrlMeta
	(v1.TaskType)(0),            // 7: common.TaskType
	(*v1.PieceTaskRequest)(nil), // 8: common.PieceTaskRequest
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
	(*v1.PiecePacket)(nil),      // 10: common.PiecePacket
}
var file_pkg_apis_dfdaemon_v1_dfdaemon_proto_depIdxs = []int32{
	6,  // 0: dfdaemon.DownRequest.url_meta:type_name -> common.UrlMeta
	6,  // 1: dfdaemon.StatTaskRequest.url_meta:type_name -> common.UrlMeta
	6,  // 2: dfdaemon.ImportTaskRequest.url_meta:type_name -> common.UrlMeta
	7,  // 3: dfdaemon.ImportTaskRequest.type:type_name -> common.TaskType
	6,  // 4: dfdaemon.ExportTaskRequest.url_meta:type_name -> common.UrlMeta
	6,  // 5: dfdaemon.DeleteTaskRequest.url_meta:type_name -> common.UrlMeta
	0,  // 6: dfdaemon.Daemon.Download:input_type -> dfdaemon.DownRequest
	8,  // 7: dfdaemon.Daemon.GetPieceTasks:input_type -> common.PieceTaskRequest
	9,  // 8: dfdaemon.Daemon.CheckHealth:input_type -> google.protobuf.Empty
	8,  // 9: dfdaemon.Daemon.SyncPieceTasks:input_type -> common.PieceTaskRequest
	2,  // 10: dfdaemon.Daemon.StatTask:input_type -> dfdaemon.StatTaskRequest
	3,  // 11: dfdaemon.Daemon.ImportTask:input_type -> dfdaemon.ImportTaskRequest
	4,  // 12: dfdaemon.Daemon.ExportTask:input_type -> dfdaemon.ExportTaskRequest
	5,  // 13: dfdaemon.Daemon.DeleteTask:input_type -> dfdaemon.DeleteTaskRequest
	9,  // 14: dfdaemon.Daemon.LeaveHost:input_type -> google.protobuf.Empty
	1,  // 15: dfdaemon.Daemon.Download:output_type -> dfdaemon.DownResult
	10, // 16: dfdaemon.Daemon.GetPieceTasks:output_type -> common.PiecePacket
	9,  // 17: dfdaemon.Daemon.CheckHealth:output_type -> google.protobuf.Empty
	10, // 18: dfdaemon.Daemon.SyncPieceTasks:output_type -> common.PiecePacket
	9,  // 19: dfdaemon.Daemon.StatTask:output_type -> google.protobuf.Empty
	9,  // 20: dfdaemon.Daemon.ImportTask:output_type -> google.protobuf.Empty
	9,  // 21: dfdaemon.Daemon.ExportTask:output_type -> google.protobuf.Empty
	9,  // 22: dfdaemon.Daemon.DeleteTask:output_type -> google.protobuf.Empty
	9,  // 23: dfdaemon.Daemon.LeaveHost:output_type -> google.protobuf.Empty
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_apis_dfdaemon_v1_dfdaemon_proto_init() }
func file_pkg_apis_dfdaemon_v1_dfdaemon_proto_init() {
	if File_pkg_apis_dfdaemon_v1_dfdaemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownRequest); i {
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
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownResult); i {
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
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatTaskRequest); i {
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
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportTaskRequest); i {
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
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTaskRequest); i {
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
		file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
			RawDescriptor: file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_dfdaemon_v1_dfdaemon_proto_goTypes,
		DependencyIndexes: file_pkg_apis_dfdaemon_v1_dfdaemon_proto_depIdxs,
		MessageInfos:      file_pkg_apis_dfdaemon_v1_dfdaemon_proto_msgTypes,
	}.Build()
	File_pkg_apis_dfdaemon_v1_dfdaemon_proto = out.File
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_rawDesc = nil
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_goTypes = nil
	file_pkg_apis_dfdaemon_v1_dfdaemon_proto_depIdxs = nil
}
