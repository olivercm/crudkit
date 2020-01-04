// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package service_customer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserListReq struct {
	Current              int64    `protobuf:"varint,1,opt,name=current,proto3" json:"current"`
	PageSize             int64    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListReq) Reset()         { *m = GetUserListReq{} }
func (m *GetUserListReq) String() string { return proto.CompactTextString(m) }
func (*GetUserListReq) ProtoMessage()    {}
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListReq.Unmarshal(m, b)
}
func (m *GetUserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListReq.Marshal(b, m, deterministic)
}
func (m *GetUserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListReq.Merge(m, src)
}
func (m *GetUserListReq) XXX_Size() int {
	return xxx_messageInfo_GetUserListReq.Size(m)
}
func (m *GetUserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListReq proto.InternalMessageInfo

func (m *GetUserListReq) GetCurrent() int64 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *GetUserListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GetUserListResp struct {
	Result               []*UserListData `protobuf:"bytes,1,rep,name=result,proto3" json:"result"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetUserListResp) Reset()         { *m = GetUserListResp{} }
func (m *GetUserListResp) String() string { return proto.CompactTextString(m) }
func (*GetUserListResp) ProtoMessage()    {}
func (*GetUserListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListResp.Unmarshal(m, b)
}
func (m *GetUserListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListResp.Marshal(b, m, deterministic)
}
func (m *GetUserListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListResp.Merge(m, src)
}
func (m *GetUserListResp) XXX_Size() int {
	return xxx_messageInfo_GetUserListResp.Size(m)
}
func (m *GetUserListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListResp proto.InternalMessageInfo

func (m *GetUserListResp) GetResult() []*UserListData {
	if m != nil {
		return m.Result
	}
	return nil
}

type UserListData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime"`
	UpdateTime           int64    `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime"`
	DeleteTime           int64    `protobuf:"varint,6,opt,name=deleteTime,proto3" json:"deleteTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListData) Reset()         { *m = UserListData{} }
func (m *UserListData) String() string { return proto.CompactTextString(m) }
func (*UserListData) ProtoMessage()    {}
func (*UserListData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserListData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListData.Unmarshal(m, b)
}
func (m *UserListData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListData.Marshal(b, m, deterministic)
}
func (m *UserListData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListData.Merge(m, src)
}
func (m *UserListData) XXX_Size() int {
	return xxx_messageInfo_UserListData.Size(m)
}
func (m *UserListData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListData.DiscardUnknown(m)
}

var xxx_messageInfo_UserListData proto.InternalMessageInfo

func (m *UserListData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserListData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserListData) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserListData) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UserListData) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *UserListData) GetDeleteTime() int64 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

type CreateUserReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime"`
	UpdateTime           int64    `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime"`
	DeleteTime           int64    `protobuf:"varint,6,opt,name=deleteTime,proto3" json:"deleteTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserReq) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *CreateUserReq) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CreateUserReq) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *CreateUserReq) GetDeleteTime() int64 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

type CreateUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

type UpdateUserReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime"`
	UpdateTime           int64    `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime"`
	DeleteTime           int64    `protobuf:"varint,6,opt,name=deleteTime,proto3" json:"deleteTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserReq) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UpdateUserReq) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UpdateUserReq) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *UpdateUserReq) GetDeleteTime() int64 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

type UpdateUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResp) Reset()         { *m = UpdateUserResp{} }
func (m *UpdateUserResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResp) ProtoMessage()    {}
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UpdateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResp.Unmarshal(m, b)
}
func (m *UpdateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResp.Merge(m, src)
}
func (m *UpdateUserResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResp.Size(m)
}
func (m *UpdateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResp proto.InternalMessageInfo

type DeleteUserReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserReq) Reset()         { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()    {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *DeleteUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserReq.Unmarshal(m, b)
}
func (m *DeleteUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserReq.Marshal(b, m, deterministic)
}
func (m *DeleteUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserReq.Merge(m, src)
}
func (m *DeleteUserReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserReq.Size(m)
}
func (m *DeleteUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserReq proto.InternalMessageInfo

func (m *DeleteUserReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (m *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(m, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetUserListReq)(nil), "service.customer.GetUserListReq")
	proto.RegisterType((*GetUserListResp)(nil), "service.customer.GetUserListResp")
	proto.RegisterType((*UserListData)(nil), "service.customer.UserListData")
	proto.RegisterType((*CreateUserReq)(nil), "service.customer.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "service.customer.CreateUserResp")
	proto.RegisterType((*UpdateUserReq)(nil), "service.customer.UpdateUserReq")
	proto.RegisterType((*UpdateUserResp)(nil), "service.customer.UpdateUserResp")
	proto.RegisterType((*DeleteUserReq)(nil), "service.customer.DeleteUserReq")
	proto.RegisterType((*DeleteUserResp)(nil), "service.customer.DeleteUserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x4e, 0xf3, 0x30,
	0x14, 0x54, 0x92, 0x7e, 0xfd, 0xe0, 0x95, 0xfe, 0xe8, 0xad, 0xa2, 0x2e, 0xda, 0x92, 0x55, 0x57,
	0x59, 0x14, 0x89, 0x0b, 0xb4, 0x02, 0x21, 0x21, 0x21, 0xa5, 0xf4, 0x00, 0x21, 0x79, 0xaa, 0x2c,
	0xf5, 0xc7, 0xd8, 0x0e, 0x0b, 0xce, 0x03, 0x77, 0xe2, 0x38, 0xc8, 0x4e, 0x4a, 0x6c, 0x1a, 0x60,
	0xdb, 0x9d, 0xfd, 0x66, 0x3c, 0x99, 0x4c, 0x26, 0x06, 0x28, 0x24, 0x89, 0x98, 0x8b, 0xbd, 0xda,
	0xe3, 0x40, 0x92, 0x78, 0x61, 0x19, 0xc5, 0x59, 0x21, 0xd5, 0x7e, 0x4b, 0x22, 0xba, 0x81, 0xde,
	0x2d, 0xa9, 0x95, 0x24, 0x71, 0xcf, 0xa4, 0x4a, 0xe8, 0x19, 0x43, 0xf8, 0x9f, 0x15, 0x42, 0xd0,
	0x4e, 0x85, 0xde, 0xc4, 0x9b, 0x06, 0xc9, 0x61, 0x8b, 0x43, 0x38, 0xe3, 0xe9, 0x9a, 0x96, 0xec,
	0x95, 0x42, 0xdf, 0x40, 0x5f, 0xfb, 0xe8, 0x0e, 0xfa, 0x8e, 0x8e, 0xe4, 0x78, 0x0d, 0x6d, 0x41,
	0xb2, 0xd8, 0x68, 0x9d, 0x60, 0xda, 0x99, 0x8d, 0xe2, 0xef, 0x4f, 0x8f, 0x0f, 0xfc, 0x45, 0xaa,
	0xd2, 0xa4, 0x62, 0x47, 0x6f, 0x1e, 0x5c, 0xd8, 0x00, 0xf6, 0xc0, 0x67, 0x79, 0x65, 0xc6, 0x67,
	0x39, 0x22, 0xb4, 0x76, 0xe9, 0xb6, 0xf4, 0x70, 0x9e, 0x98, 0x35, 0x0e, 0x20, 0x48, 0xd7, 0x14,
	0x06, 0x86, 0xa4, 0x97, 0x38, 0x02, 0xc8, 0x04, 0xa5, 0x8a, 0x1e, 0xd9, 0x96, 0xc2, 0x96, 0x01,
	0xac, 0x89, 0xc6, 0x0b, 0x9e, 0x1f, 0xf0, 0x7f, 0x25, 0x5e, 0x4f, 0x34, 0x9e, 0xd3, 0x86, 0x2a,
	0xbc, 0x5d, 0xe2, 0xf5, 0x24, 0x7a, 0xf7, 0xa0, 0x3b, 0x37, 0x72, 0xda, 0xac, 0x4e, 0xee, 0x34,
	0x7d, 0x0e, 0xa0, 0x67, 0xdb, 0x94, 0xdc, 0x38, 0x5f, 0x19, 0x81, 0x93, 0x77, 0x6e, 0xdb, 0x94,
	0x3c, 0x1a, 0x43, 0x77, 0x61, 0xf0, 0x1f, 0x8c, 0xeb, 0x23, 0x36, 0x41, 0xf2, 0xd9, 0x87, 0x0f,
	0xfd, 0x79, 0xd5, 0xb7, 0x65, 0xd9, 0x3f, 0x4c, 0xa0, 0x63, 0x95, 0x15, 0x27, 0xc7, 0xc5, 0x74,
	0xff, 0x89, 0xe1, 0xe5, 0x1f, 0x0c, 0xc9, 0xf1, 0x01, 0xa0, 0x8e, 0x19, 0xc7, 0xc7, 0x07, 0x9c,
	0xae, 0x0c, 0x27, 0xbf, 0x13, 0x4a, 0xc1, 0xfa, 0xed, 0x9b, 0x04, 0x9d, 0x4f, 0xd8, 0x24, 0xe8,
	0x86, 0xa7, 0x05, 0xeb, 0x6c, 0x9a, 0x04, 0x9d, 0x68, 0x9b, 0x04, 0xdd, 0x68, 0x9f, 0xda, 0xe6,
	0x52, 0xb9, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x02, 0x5d, 0x45, 0x62, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error) {
	out := new(GetUserListResp)
	err := c.cc.Invoke(ctx, "/service.customer.CustomerService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/service.customer.CustomerService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/service.customer.CustomerService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	out := new(DeleteUserResp)
	err := c.cc.Invoke(ctx, "/service.customer.CustomerService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) GetUserList(ctx context.Context, req *GetUserListReq) (*GetUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedCustomerServiceServer) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedCustomerServiceServer) UpdateUser(ctx context.Context, req *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedCustomerServiceServer) DeleteUser(ctx context.Context, req *DeleteUserReq) (*DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.customer.CustomerService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.customer.CustomerService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.customer.CustomerService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.customer.CustomerService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _CustomerService_GetUserList_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _CustomerService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _CustomerService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _CustomerService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
