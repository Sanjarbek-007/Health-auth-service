// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: lifeStyle.proto

package health

import (
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

type AddLifeStyleDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType  string `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue string `protobuf:"bytes,3,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
}

func (x *AddLifeStyleDataReq) Reset() {
	*x = AddLifeStyleDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLifeStyleDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLifeStyleDataReq) ProtoMessage() {}

func (x *AddLifeStyleDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLifeStyleDataReq.ProtoReflect.Descriptor instead.
func (*AddLifeStyleDataReq) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{0}
}

func (x *AddLifeStyleDataReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddLifeStyleDataReq) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *AddLifeStyleDataReq) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

type AddLifeStyleDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddLifeStyleDataRes) Reset() {
	*x = AddLifeStyleDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLifeStyleDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLifeStyleDataRes) ProtoMessage() {}

func (x *AddLifeStyleDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLifeStyleDataRes.ProtoReflect.Descriptor instead.
func (*AddLifeStyleDataRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{1}
}

func (x *AddLifeStyleDataRes) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

type GetLifeStyleDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *GetLifeStyleDataReq) Reset() {
	*x = GetLifeStyleDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyleDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyleDataReq) ProtoMessage() {}

func (x *GetLifeStyleDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyleDataReq.ProtoReflect.Descriptor instead.
func (*GetLifeStyleDataReq) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{2}
}

func (x *GetLifeStyleDataReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetLifeStyleDataReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetLifeStyleDataReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GetLifeStyle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DataType     string `protobuf:"bytes,4,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue    string `protobuf:"bytes,5,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	RecordedDate string `protobuf:"bytes,6,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
}

func (x *GetLifeStyle) Reset() {
	*x = GetLifeStyle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyle) ProtoMessage() {}

func (x *GetLifeStyle) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyle.ProtoReflect.Descriptor instead.
func (*GetLifeStyle) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{3}
}

func (x *GetLifeStyle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLifeStyle) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetLifeStyle) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetLifeStyle) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *GetLifeStyle) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *GetLifeStyle) GetRecordedDate() string {
	if x != nil {
		return x.RecordedDate
	}
	return ""
}

type GetLifeStyleDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LifeStyle []*GetLifeStyle `protobuf:"bytes,1,rep,name=life_style,json=lifeStyle,proto3" json:"life_style,omitempty"`
}

func (x *GetLifeStyleDataRes) Reset() {
	*x = GetLifeStyleDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyleDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyleDataRes) ProtoMessage() {}

func (x *GetLifeStyleDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyleDataRes.ProtoReflect.Descriptor instead.
func (*GetLifeStyleDataRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{4}
}

func (x *GetLifeStyleDataRes) GetLifeStyle() []*GetLifeStyle {
	if x != nil {
		return x.LifeStyle
	}
	return nil
}

type GetLifeStyleDataByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLifeStyleDataByIdReq) Reset() {
	*x = GetLifeStyleDataByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyleDataByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyleDataByIdReq) ProtoMessage() {}

func (x *GetLifeStyleDataByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyleDataByIdReq.ProtoReflect.Descriptor instead.
func (*GetLifeStyleDataByIdReq) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{5}
}

func (x *GetLifeStyleDataByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLifeStyleByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DataType     string `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue    string `protobuf:"bytes,4,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
	RecordedDate string `protobuf:"bytes,5,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
}

func (x *GetLifeStyleByIdRes) Reset() {
	*x = GetLifeStyleByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyleByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyleByIdRes) ProtoMessage() {}

func (x *GetLifeStyleByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyleByIdRes.ProtoReflect.Descriptor instead.
func (*GetLifeStyleByIdRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{6}
}

func (x *GetLifeStyleByIdRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLifeStyleByIdRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetLifeStyleByIdRes) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *GetLifeStyleByIdRes) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *GetLifeStyleByIdRes) GetRecordedDate() string {
	if x != nil {
		return x.RecordedDate
	}
	return ""
}

type GetLifeStyleDataByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LifeStyle *GetLifeStyleByIdRes `protobuf:"bytes,1,opt,name=life_style,json=lifeStyle,proto3" json:"life_style,omitempty"`
}

func (x *GetLifeStyleDataByIdRes) Reset() {
	*x = GetLifeStyleDataByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLifeStyleDataByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLifeStyleDataByIdRes) ProtoMessage() {}

func (x *GetLifeStyleDataByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLifeStyleDataByIdRes.ProtoReflect.Descriptor instead.
func (*GetLifeStyleDataByIdRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{7}
}

func (x *GetLifeStyleDataByIdRes) GetLifeStyle() *GetLifeStyleByIdRes {
	if x != nil {
		return x.LifeStyle
	}
	return nil
}

type UpdateLifeStyleDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DataType  string `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	DataValue string `protobuf:"bytes,3,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
}

func (x *UpdateLifeStyleDataReq) Reset() {
	*x = UpdateLifeStyleDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLifeStyleDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLifeStyleDataReq) ProtoMessage() {}

func (x *UpdateLifeStyleDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLifeStyleDataReq.ProtoReflect.Descriptor instead.
func (*UpdateLifeStyleDataReq) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateLifeStyleDataReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLifeStyleDataReq) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *UpdateLifeStyleDataReq) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

type UpdateLifeStyleDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateLifeStyleDataRes) Reset() {
	*x = UpdateLifeStyleDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLifeStyleDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLifeStyleDataRes) ProtoMessage() {}

func (x *UpdateLifeStyleDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLifeStyleDataRes.ProtoReflect.Descriptor instead.
func (*UpdateLifeStyleDataRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateLifeStyleDataRes) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

type DeleteLifeStyleDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteLifeStyleDataReq) Reset() {
	*x = DeleteLifeStyleDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLifeStyleDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLifeStyleDataReq) ProtoMessage() {}

func (x *DeleteLifeStyleDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLifeStyleDataReq.ProtoReflect.Descriptor instead.
func (*DeleteLifeStyleDataReq) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteLifeStyleDataReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteLifeStyleDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteLifeStyleDataRes) Reset() {
	*x = DeleteLifeStyleDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lifeStyle_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLifeStyleDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLifeStyleDataRes) ProtoMessage() {}

func (x *DeleteLifeStyleDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_lifeStyle_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLifeStyleDataRes.ProtoReflect.Descriptor instead.
func (*DeleteLifeStyleDataRes) Descriptor() ([]byte, []int) {
	return file_lifeStyle_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteLifeStyleDataRes) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

var File_lifeStyle_proto protoreflect.FileDescriptor

var file_lifeStyle_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x22, 0x6a, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4c,
	0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66,
	0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x6c, 0x69,
	0x66, 0x65, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9f, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x58, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x6c, 0x69,
	0x66, 0x65, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x52, 0x09,
	0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x22, 0x64, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x32, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x66,
	0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xcd, 0x03, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12,
	0x52, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x22, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21,
	0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x6c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x21,
	0x2e, 0x6c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lifeStyle_proto_rawDescOnce sync.Once
	file_lifeStyle_proto_rawDescData = file_lifeStyle_proto_rawDesc
)

func file_lifeStyle_proto_rawDescGZIP() []byte {
	file_lifeStyle_proto_rawDescOnce.Do(func() {
		file_lifeStyle_proto_rawDescData = protoimpl.X.CompressGZIP(file_lifeStyle_proto_rawDescData)
	})
	return file_lifeStyle_proto_rawDescData
}

var file_lifeStyle_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_lifeStyle_proto_goTypes = []any{
	(*AddLifeStyleDataReq)(nil),     // 0: lifeStyle.AddLifeStyleDataReq
	(*AddLifeStyleDataRes)(nil),     // 1: lifeStyle.AddLifeStyleDataRes
	(*GetLifeStyleDataReq)(nil),     // 2: lifeStyle.GetLifeStyleDataReq
	(*GetLifeStyle)(nil),            // 3: lifeStyle.GetLifeStyle
	(*GetLifeStyleDataRes)(nil),     // 4: lifeStyle.GetLifeStyleDataRes
	(*GetLifeStyleDataByIdReq)(nil), // 5: lifeStyle.GetLifeStyleDataByIdReq
	(*GetLifeStyleByIdRes)(nil),     // 6: lifeStyle.GetLifeStyleByIdRes
	(*GetLifeStyleDataByIdRes)(nil), // 7: lifeStyle.GetLifeStyleDataByIdRes
	(*UpdateLifeStyleDataReq)(nil),  // 8: lifeStyle.UpdateLifeStyleDataReq
	(*UpdateLifeStyleDataRes)(nil),  // 9: lifeStyle.UpdateLifeStyleDataRes
	(*DeleteLifeStyleDataReq)(nil),  // 10: lifeStyle.DeleteLifeStyleDataReq
	(*DeleteLifeStyleDataRes)(nil),  // 11: lifeStyle.DeleteLifeStyleDataRes
}
var file_lifeStyle_proto_depIdxs = []int32{
	3,  // 0: lifeStyle.GetLifeStyleDataRes.life_style:type_name -> lifeStyle.GetLifeStyle
	6,  // 1: lifeStyle.GetLifeStyleDataByIdRes.life_style:type_name -> lifeStyle.GetLifeStyleByIdRes
	0,  // 2: lifeStyle.lifeStyle.AddLifeStyleData:input_type -> lifeStyle.AddLifeStyleDataReq
	2,  // 3: lifeStyle.lifeStyle.GetLifeStyleData:input_type -> lifeStyle.GetLifeStyleDataReq
	5,  // 4: lifeStyle.lifeStyle.GetLifeStyleDataById:input_type -> lifeStyle.GetLifeStyleDataByIdReq
	8,  // 5: lifeStyle.lifeStyle.UpdateLifeStyleData:input_type -> lifeStyle.UpdateLifeStyleDataReq
	10, // 6: lifeStyle.lifeStyle.DeleteLifeStyleData:input_type -> lifeStyle.DeleteLifeStyleDataReq
	1,  // 7: lifeStyle.lifeStyle.AddLifeStyleData:output_type -> lifeStyle.AddLifeStyleDataRes
	4,  // 8: lifeStyle.lifeStyle.GetLifeStyleData:output_type -> lifeStyle.GetLifeStyleDataRes
	7,  // 9: lifeStyle.lifeStyle.GetLifeStyleDataById:output_type -> lifeStyle.GetLifeStyleDataByIdRes
	9,  // 10: lifeStyle.lifeStyle.UpdateLifeStyleData:output_type -> lifeStyle.UpdateLifeStyleDataRes
	11, // 11: lifeStyle.lifeStyle.DeleteLifeStyleData:output_type -> lifeStyle.DeleteLifeStyleDataRes
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_lifeStyle_proto_init() }
func file_lifeStyle_proto_init() {
	if File_lifeStyle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lifeStyle_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddLifeStyleDataReq); i {
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
		file_lifeStyle_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddLifeStyleDataRes); i {
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
		file_lifeStyle_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyleDataReq); i {
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
		file_lifeStyle_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyle); i {
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
		file_lifeStyle_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyleDataRes); i {
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
		file_lifeStyle_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyleDataByIdReq); i {
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
		file_lifeStyle_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyleByIdRes); i {
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
		file_lifeStyle_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetLifeStyleDataByIdRes); i {
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
		file_lifeStyle_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateLifeStyleDataReq); i {
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
		file_lifeStyle_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateLifeStyleDataRes); i {
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
		file_lifeStyle_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteLifeStyleDataReq); i {
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
		file_lifeStyle_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteLifeStyleDataRes); i {
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
			RawDescriptor: file_lifeStyle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lifeStyle_proto_goTypes,
		DependencyIndexes: file_lifeStyle_proto_depIdxs,
		MessageInfos:      file_lifeStyle_proto_msgTypes,
	}.Build()
	File_lifeStyle_proto = out.File
	file_lifeStyle_proto_rawDesc = nil
	file_lifeStyle_proto_goTypes = nil
	file_lifeStyle_proto_depIdxs = nil
}