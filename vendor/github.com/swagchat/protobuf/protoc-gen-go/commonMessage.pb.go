// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commonMessage.proto

package protoc_gen_go

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Order int32

const (
	Order_Asc  Order = 0
	Order_Desc Order = 1
)

var Order_name = map[int32]string{
	0: "Asc",
	1: "Desc",
}
var Order_value = map[string]int32{
	"Asc":  0,
	"Desc": 1,
}

func (x Order) String() string {
	return proto.EnumName(Order_name, int32(x))
}
func (Order) EnumDescriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{0} }

type RoomIds struct {
	RoomIDs []string `protobuf:"bytes,1,rep,name=room_ids,json=roomIds" json:"roomIds,omitempty"`
}

func (m *RoomIds) Reset()                    { *m = RoomIds{} }
func (m *RoomIds) String() string            { return proto.CompactTextString(m) }
func (*RoomIds) ProtoMessage()               {}
func (*RoomIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{0} }

func (m *RoomIds) GetRoomIDs() []string {
	if m != nil {
		return m.RoomIDs
	}
	return nil
}

type UserIds struct {
	UserIDs []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds" json:"userIds,omitempty"`
}

func (m *UserIds) Reset()                    { *m = UserIds{} }
func (m *UserIds) String() string            { return proto.CompactTextString(m) }
func (*UserIds) ProtoMessage()               {}
func (*UserIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{1} }

func (m *UserIds) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type RoleIds struct {
	RoleIDs []int32 `protobuf:"varint,1,rep,packed,name=role_ids,json=roleIds" json:"roleIds"`
}

func (m *RoleIds) Reset()                    { *m = RoleIds{} }
func (m *RoleIds) String() string            { return proto.CompactTextString(m) }
func (*RoleIds) ProtoMessage()               {}
func (*RoleIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{2} }

func (m *RoleIds) GetRoleIDs() []int32 {
	if m != nil {
		return m.RoleIDs
	}
	return nil
}

type ErrorResponse struct {
	Message          string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	DeveloperMessage string          `protobuf:"bytes,2,opt,name=developer_message,json=developerMessage,proto3" json:"developerMessage,omitempty"`
	Info             string          `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	InvalidParams    []*InvalidParam `protobuf:"bytes,5,rep,name=invalid_params,json=invalidParams" json:"invalidParams,omitempty"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{3} }

func (m *ErrorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorResponse) GetDeveloperMessage() string {
	if m != nil {
		return m.DeveloperMessage
	}
	return ""
}

func (m *ErrorResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ErrorResponse) GetInvalidParams() []*InvalidParam {
	if m != nil {
		return m.InvalidParams
	}
	return nil
}

type InvalidParam struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *InvalidParam) Reset()                    { *m = InvalidParam{} }
func (m *InvalidParam) String() string            { return proto.CompactTextString(m) }
func (*InvalidParam) ProtoMessage()               {}
func (*InvalidParam) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{4} }

func (m *InvalidParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvalidParam) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type OrderInfo struct {
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order Order  `protobuf:"varint,2,opt,name=order,proto3,enum=swagchat.protobuf.Order" json:"order,omitempty"`
}

func (m *OrderInfo) Reset()                    { *m = OrderInfo{} }
func (m *OrderInfo) String() string            { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()               {}
func (*OrderInfo) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{5} }

func (m *OrderInfo) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *OrderInfo) GetOrder() Order {
	if m != nil {
		return m.Order
	}
	return Order_Asc
}

func init() {
	proto.RegisterType((*RoomIds)(nil), "swagchat.protobuf.RoomIds")
	proto.RegisterType((*UserIds)(nil), "swagchat.protobuf.UserIds")
	proto.RegisterType((*RoleIds)(nil), "swagchat.protobuf.RoleIds")
	proto.RegisterType((*ErrorResponse)(nil), "swagchat.protobuf.ErrorResponse")
	proto.RegisterType((*InvalidParam)(nil), "swagchat.protobuf.InvalidParam")
	proto.RegisterType((*OrderInfo)(nil), "swagchat.protobuf.OrderInfo")
	proto.RegisterEnum("swagchat.protobuf.Order", Order_name, Order_value)
}

func init() { proto.RegisterFile("commonMessage.proto", fileDescriptorCommonMessage) }

var fileDescriptorCommonMessage = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x9b, 0x26, 0xae, 0x9b, 0x81, 0x46, 0xc9, 0x52, 0x8a, 0x15, 0x24, 0x1c, 0x71, 0x40,
	0x51, 0xd5, 0x3a, 0xa2, 0xdc, 0xe0, 0x02, 0x56, 0x39, 0x44, 0x50, 0x01, 0x96, 0xb8, 0x70, 0x89,
	0x1c, 0x7b, 0xe2, 0x5a, 0xf2, 0x7a, 0xac, 0x5d, 0xa7, 0x88, 0xb7, 0x81, 0x37, 0xe2, 0x09, 0x7c,
	0xe0, 0x98, 0xa7, 0x40, 0xfb, 0xc7, 0x95, 0x69, 0xb9, 0xcd, 0x37, 0xdf, 0x37, 0x3f, 0xcd, 0xce,
	0xc2, 0xa3, 0x84, 0x38, 0xa7, 0xf2, 0x0a, 0xa5, 0x8c, 0x33, 0x0c, 0x2a, 0x41, 0x35, 0xb1, 0x89,
	0xfc, 0x1e, 0x67, 0xc9, 0x75, 0x5c, 0x1b, 0xbd, 0xde, 0x6e, 0xa6, 0xc7, 0x19, 0x65, 0xa4, 0xd5,
	0x42, 0x55, 0xc6, 0x78, 0xfe, 0x11, 0xdc, 0x88, 0x88, 0x2f, 0x53, 0xc9, 0xde, 0xc0, 0xa1, 0x20,
	0xe2, 0xab, 0x3c, 0x95, 0x5e, 0x6f, 0xd6, 0x9f, 0x0f, 0xc3, 0xd9, 0x9f, 0xc6, 0x37, 0xf6, 0xa5,
	0xdc, 0x35, 0xfe, 0x44, 0x98, 0xe4, 0x19, 0xf1, 0xbc, 0x46, 0x5e, 0xd5, 0x3f, 0x22, 0xd7, 0xb6,
	0x5e, 0x0f, 0x7e, 0xff, 0xf2, 0xf7, 0x14, 0xed, 0xab, 0x44, 0x61, 0x69, 0x5b, 0x89, 0xe2, 0x2e,
	0x4d, 0xdb, 0x86, 0xb6, 0x35, 0xc9, 0x2e, 0xcd, 0xb6, 0x2c, 0x2d, 0x54, 0xbb, 0x15, 0xa8, 0x68,
	0x2f, 0xd5, 0x6e, 0x05, 0xde, 0xd2, 0x9c, 0xf0, 0xc4, 0xec, 0x56, 0xa0, 0xa1, 0xb9, 0xc2, 0x24,
	0xa3, 0xb6, 0xb0, 0x8c, 0x9f, 0xfb, 0x70, 0xf4, 0x5e, 0x08, 0x12, 0x11, 0xca, 0x8a, 0x4a, 0x89,
	0x6c, 0x01, 0x2e, 0x37, 0xb7, 0xf2, 0x7a, 0xb3, 0xde, 0x7c, 0x18, 0x3e, 0x56, 0xcb, 0xd8, 0x56,
	0x77, 0x19, 0xdb, 0x62, 0x1f, 0x60, 0x92, 0xe2, 0x0d, 0x16, 0x54, 0xa1, 0x58, 0xb5, 0xa3, 0xfb,
	0x7a, 0xf4, 0xd9, 0xae, 0xf1, 0xa7, 0xb7, 0xe6, 0xd5, 0x3d, 0xc6, 0xf8, 0xae, 0xc7, 0x5e, 0xc0,
	0x20, 0x2f, 0x37, 0xe4, 0xf5, 0xf5, 0x3c, 0xdb, 0x35, 0xfe, 0x48, 0xe9, 0xce, 0x8c, 0xf6, 0x59,
	0x02, 0xa3, 0xbc, 0xbc, 0x89, 0x8b, 0x3c, 0x5d, 0x55, 0xb1, 0x88, 0xb9, 0xf4, 0x9c, 0x59, 0x7f,
	0xfe, 0xe0, 0xc2, 0x0f, 0xee, 0xfd, 0x6c, 0xb0, 0x34, 0xc1, 0xcf, 0x2a, 0x17, 0x3e, 0xdd, 0x35,
	0xfe, 0x93, 0xbc, 0xd3, 0xe9, 0x1e, 0xf8, 0xe8, 0x1f, 0xc3, 0x9e, 0xe8, 0x2d, 0x3c, 0xec, 0x12,
	0x18, 0x83, 0x41, 0x19, 0x73, 0x7b, 0x9d, 0x48, 0xd7, 0xec, 0x04, 0x0e, 0x04, 0xc6, 0x92, 0x4a,
	0xf3, 0xf0, 0xc8, 0x2a, 0x4b, 0xf8, 0x02, 0xc3, 0x4f, 0x22, 0x45, 0xb1, 0x54, 0x9b, 0x1f, 0x83,
	0xb3, 0xc9, 0xb1, 0x48, 0xed, 0xbc, 0x11, 0x2c, 0x00, 0x87, 0x54, 0x44, 0xcf, 0x8f, 0x2e, 0xbc,
	0xff, 0x3c, 0x43, 0x23, 0x22, 0x13, 0x3b, 0x9d, 0x82, 0xa3, 0x35, 0x73, 0xa1, 0xff, 0x4e, 0x26,
	0xe3, 0x3d, 0x76, 0x08, 0x83, 0x4b, 0x94, 0xc9, 0xb8, 0x17, 0x9e, 0x7d, 0x3b, 0xcd, 0xf2, 0xfa,
	0x7a, 0xbb, 0x0e, 0x12, 0xe2, 0x8b, 0x16, 0xb4, 0x68, 0x41, 0xa6, 0x48, 0xce, 0x33, 0x2c, 0xcf,
	0x33, 0x5a, 0x1f, 0x68, 0xf9, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xc0, 0x02, 0xeb,
	0x28, 0x03, 0x00, 0x00,
}
