//
//  Copyright 2020 Docker Compose CLI authors

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: protos/compose/v1/compose.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ComposeUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string   `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	WorkDir     string   `protobuf:"bytes,2,opt,name=workDir,proto3" json:"workDir,omitempty"`
	Files       []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ComposeUpRequest) Reset() {
	*x = ComposeUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUpRequest) ProtoMessage() {}

func (x *ComposeUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUpRequest.ProtoReflect.Descriptor instead.
func (*ComposeUpRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{0}
}

func (x *ComposeUpRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ComposeUpRequest) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *ComposeUpRequest) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type ComposeUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
}

func (x *ComposeUpResponse) Reset() {
	*x = ComposeUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUpResponse) ProtoMessage() {}

func (x *ComposeUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUpResponse.ProtoReflect.Descriptor instead.
func (*ComposeUpResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeUpResponse) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ComposeDownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string   `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	WorkDir     string   `protobuf:"bytes,2,opt,name=workDir,proto3" json:"workDir,omitempty"`
	Files       []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ComposeDownRequest) Reset() {
	*x = ComposeDownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeDownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeDownRequest) ProtoMessage() {}

func (x *ComposeDownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeDownRequest.ProtoReflect.Descriptor instead.
func (*ComposeDownRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{2}
}

func (x *ComposeDownRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ComposeDownRequest) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *ComposeDownRequest) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type ComposeDownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
}

func (x *ComposeDownResponse) Reset() {
	*x = ComposeDownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeDownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeDownResponse) ProtoMessage() {}

func (x *ComposeDownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeDownResponse.ProtoReflect.Descriptor instead.
func (*ComposeDownResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{3}
}

func (x *ComposeDownResponse) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ComposeStacksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
}

func (x *ComposeStacksRequest) Reset() {
	*x = ComposeStacksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeStacksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeStacksRequest) ProtoMessage() {}

func (x *ComposeStacksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeStacksRequest.ProtoReflect.Descriptor instead.
func (*ComposeStacksRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{4}
}

func (x *ComposeStacksRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ComposeStacksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stacks []*Stack `protobuf:"bytes,1,rep,name=stacks,proto3" json:"stacks,omitempty"`
}

func (x *ComposeStacksResponse) Reset() {
	*x = ComposeStacksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeStacksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeStacksResponse) ProtoMessage() {}

func (x *ComposeStacksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeStacksResponse.ProtoReflect.Descriptor instead.
func (*ComposeStacksResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{5}
}

func (x *ComposeStacksResponse) GetStacks() []*Stack {
	if x != nil {
		return x.Stacks
	}
	return nil
}

type Stack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Reason string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Stack) Reset() {
	*x = Stack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stack) ProtoMessage() {}

func (x *Stack) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stack.ProtoReflect.Descriptor instead.
func (*Stack) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{6}
}

func (x *Stack) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stack) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Stack) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ComposeServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string   `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	WorkDir     string   `protobuf:"bytes,2,opt,name=workDir,proto3" json:"workDir,omitempty"`
	Files       []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ComposeServicesRequest) Reset() {
	*x = ComposeServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServicesRequest) ProtoMessage() {}

func (x *ComposeServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServicesRequest.ProtoReflect.Descriptor instead.
func (*ComposeServicesRequest) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{7}
}

func (x *ComposeServicesRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ComposeServicesRequest) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *ComposeServicesRequest) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type ComposeServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ComposeServicesResponse) Reset() {
	*x = ComposeServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeServicesResponse) ProtoMessage() {}

func (x *ComposeServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeServicesResponse.ProtoReflect.Descriptor instead.
func (*ComposeServicesResponse) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{8}
}

func (x *ComposeServicesResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Replicas uint32   `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Desired  uint32   `protobuf:"varint,4,opt,name=desired,proto3" json:"desired,omitempty"`
	Ports    []string `protobuf:"bytes,5,rep,name=Ports,proto3" json:"Ports,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_compose_v1_compose_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_protos_compose_v1_compose_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_protos_compose_v1_compose_proto_rawDescGZIP(), []int{9}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Service) GetDesired() uint32 {
	if x != nil {
		return x.Desired
	}
	return 0
}

func (x *Service) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

var File_protos_compose_v1_compose_proto protoreflect.FileDescriptor

var file_protos_compose_v1_compose_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x44, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x44, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x66, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x44, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44,
	0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x38, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x15, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x5b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x6a, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x60,
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x22, 0x79, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65,
	0x73, 0x69, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x32, 0xe9, 0x03, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x02, 0x55, 0x70, 0x12, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x04, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x34,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_compose_v1_compose_proto_rawDescOnce sync.Once
	file_protos_compose_v1_compose_proto_rawDescData = file_protos_compose_v1_compose_proto_rawDesc
)

func file_protos_compose_v1_compose_proto_rawDescGZIP() []byte {
	file_protos_compose_v1_compose_proto_rawDescOnce.Do(func() {
		file_protos_compose_v1_compose_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_compose_v1_compose_proto_rawDescData)
	})
	return file_protos_compose_v1_compose_proto_rawDescData
}

var file_protos_compose_v1_compose_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_compose_v1_compose_proto_goTypes = []interface{}{
	(*ComposeUpRequest)(nil),        // 0: com.docker.api.protos.compose.v1.ComposeUpRequest
	(*ComposeUpResponse)(nil),       // 1: com.docker.api.protos.compose.v1.ComposeUpResponse
	(*ComposeDownRequest)(nil),      // 2: com.docker.api.protos.compose.v1.ComposeDownRequest
	(*ComposeDownResponse)(nil),     // 3: com.docker.api.protos.compose.v1.ComposeDownResponse
	(*ComposeStacksRequest)(nil),    // 4: com.docker.api.protos.compose.v1.ComposeStacksRequest
	(*ComposeStacksResponse)(nil),   // 5: com.docker.api.protos.compose.v1.ComposeStacksResponse
	(*Stack)(nil),                   // 6: com.docker.api.protos.compose.v1.Stack
	(*ComposeServicesRequest)(nil),  // 7: com.docker.api.protos.compose.v1.ComposeServicesRequest
	(*ComposeServicesResponse)(nil), // 8: com.docker.api.protos.compose.v1.ComposeServicesResponse
	(*Service)(nil),                 // 9: com.docker.api.protos.compose.v1.Service
}
var file_protos_compose_v1_compose_proto_depIdxs = []int32{
	6, // 0: com.docker.api.protos.compose.v1.ComposeStacksResponse.stacks:type_name -> com.docker.api.protos.compose.v1.Stack
	9, // 1: com.docker.api.protos.compose.v1.ComposeServicesResponse.services:type_name -> com.docker.api.protos.compose.v1.Service
	0, // 2: com.docker.api.protos.compose.v1.Compose.Up:input_type -> com.docker.api.protos.compose.v1.ComposeUpRequest
	2, // 3: com.docker.api.protos.compose.v1.Compose.Down:input_type -> com.docker.api.protos.compose.v1.ComposeDownRequest
	4, // 4: com.docker.api.protos.compose.v1.Compose.Stacks:input_type -> com.docker.api.protos.compose.v1.ComposeStacksRequest
	7, // 5: com.docker.api.protos.compose.v1.Compose.Services:input_type -> com.docker.api.protos.compose.v1.ComposeServicesRequest
	1, // 6: com.docker.api.protos.compose.v1.Compose.Up:output_type -> com.docker.api.protos.compose.v1.ComposeUpResponse
	3, // 7: com.docker.api.protos.compose.v1.Compose.Down:output_type -> com.docker.api.protos.compose.v1.ComposeDownResponse
	5, // 8: com.docker.api.protos.compose.v1.Compose.Stacks:output_type -> com.docker.api.protos.compose.v1.ComposeStacksResponse
	8, // 9: com.docker.api.protos.compose.v1.Compose.Services:output_type -> com.docker.api.protos.compose.v1.ComposeServicesResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_compose_v1_compose_proto_init() }
func file_protos_compose_v1_compose_proto_init() {
	if File_protos_compose_v1_compose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_compose_v1_compose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUpRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUpResponse); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeDownRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeDownResponse); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeStacksRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeStacksResponse); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stack); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServicesRequest); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeServicesResponse); i {
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
		file_protos_compose_v1_compose_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_protos_compose_v1_compose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_compose_v1_compose_proto_goTypes,
		DependencyIndexes: file_protos_compose_v1_compose_proto_depIdxs,
		MessageInfos:      file_protos_compose_v1_compose_proto_msgTypes,
	}.Build()
	File_protos_compose_v1_compose_proto = out.File
	file_protos_compose_v1_compose_proto_rawDesc = nil
	file_protos_compose_v1_compose_proto_goTypes = nil
	file_protos_compose_v1_compose_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComposeClient is the client API for Compose service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComposeClient interface {
	Up(ctx context.Context, in *ComposeUpRequest, opts ...grpc.CallOption) (*ComposeUpResponse, error)
	Down(ctx context.Context, in *ComposeDownRequest, opts ...grpc.CallOption) (*ComposeDownResponse, error)
	Stacks(ctx context.Context, in *ComposeStacksRequest, opts ...grpc.CallOption) (*ComposeStacksResponse, error)
	Services(ctx context.Context, in *ComposeServicesRequest, opts ...grpc.CallOption) (*ComposeServicesResponse, error)
}

type composeClient struct {
	cc grpc.ClientConnInterface
}

func NewComposeClient(cc grpc.ClientConnInterface) ComposeClient {
	return &composeClient{cc}
}

func (c *composeClient) Up(ctx context.Context, in *ComposeUpRequest, opts ...grpc.CallOption) (*ComposeUpResponse, error) {
	out := new(ComposeUpResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Up", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *composeClient) Down(ctx context.Context, in *ComposeDownRequest, opts ...grpc.CallOption) (*ComposeDownResponse, error) {
	out := new(ComposeDownResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Down", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *composeClient) Stacks(ctx context.Context, in *ComposeStacksRequest, opts ...grpc.CallOption) (*ComposeStacksResponse, error) {
	out := new(ComposeStacksResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Stacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *composeClient) Services(ctx context.Context, in *ComposeServicesRequest, opts ...grpc.CallOption) (*ComposeServicesResponse, error) {
	out := new(ComposeServicesResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.compose.v1.Compose/Services", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComposeServer is the server API for Compose service.
type ComposeServer interface {
	Up(context.Context, *ComposeUpRequest) (*ComposeUpResponse, error)
	Down(context.Context, *ComposeDownRequest) (*ComposeDownResponse, error)
	Stacks(context.Context, *ComposeStacksRequest) (*ComposeStacksResponse, error)
	Services(context.Context, *ComposeServicesRequest) (*ComposeServicesResponse, error)
}

// UnimplementedComposeServer can be embedded to have forward compatible implementations.
type UnimplementedComposeServer struct {
}

func (*UnimplementedComposeServer) Up(context.Context, *ComposeUpRequest) (*ComposeUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (*UnimplementedComposeServer) Down(context.Context, *ComposeDownRequest) (*ComposeDownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Down not implemented")
}
func (*UnimplementedComposeServer) Stacks(context.Context, *ComposeStacksRequest) (*ComposeStacksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stacks not implemented")
}
func (*UnimplementedComposeServer) Services(context.Context, *ComposeServicesRequest) (*ComposeServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Services not implemented")
}

func RegisterComposeServer(s *grpc.Server, srv ComposeServer) {
	s.RegisterService(&_Compose_serviceDesc, srv)
}

func _Compose_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Up(ctx, req.(*ComposeUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compose_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Down(ctx, req.(*ComposeDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compose_Stacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeStacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Stacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Stacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Stacks(ctx, req.(*ComposeStacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compose_Services_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComposeServer).Services(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.compose.v1.Compose/Services",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComposeServer).Services(ctx, req.(*ComposeServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compose_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.protos.compose.v1.Compose",
	HandlerType: (*ComposeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _Compose_Up_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _Compose_Down_Handler,
		},
		{
			MethodName: "Stacks",
			Handler:    _Compose_Stacks_Handler,
		},
		{
			MethodName: "Services",
			Handler:    _Compose_Services_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/compose/v1/compose.proto",
}
