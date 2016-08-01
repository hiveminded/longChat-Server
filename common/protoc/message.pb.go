// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package protoc is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	MessageReq
	OnlineReq
	GroupReq
*/
package protoc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageReq_Message_ContentType int32

const (
	MessageReq_Message_Text    MessageReq_Message_ContentType = 0
	MessageReq_Message_Pic     MessageReq_Message_ContentType = 1
	MessageReq_Message_Emotion MessageReq_Message_ContentType = 2
	MessageReq_Message_Voice   MessageReq_Message_ContentType = 3
)

var MessageReq_Message_ContentType_name = map[int32]string{
	0: "Text",
	1: "Pic",
	2: "Emotion",
	3: "Voice",
}
var MessageReq_Message_ContentType_value = map[string]int32{
	"Text":    0,
	"Pic":     1,
	"Emotion": 2,
	"Voice":   3,
}

func (x MessageReq_Message_ContentType) String() string {
	return proto.EnumName(MessageReq_Message_ContentType_name, int32(x))
}
func (MessageReq_Message_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type MessageReq struct {
	Messages []*MessageReq_Message `protobuf:"bytes,1,rep,name=Messages,json=messages" json:"Messages,omitempty"`
}

func (m *MessageReq) Reset()                    { *m = MessageReq{} }
func (m *MessageReq) String() string            { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()               {}
func (*MessageReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MessageReq) GetMessages() []*MessageReq_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type MessageReq_Message struct {
	Id             []byte                         `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	From           []byte                         `protobuf:"bytes,2,opt,name=From,json=from,proto3" json:"From,omitempty"`
	To             []byte                         `protobuf:"bytes,3,opt,name=To,json=to,proto3" json:"To,omitempty"`
	Content        string                         `protobuf:"bytes,4,opt,name=Content,json=content" json:"Content,omitempty"`
	Type           MessageReq_Message_ContentType `protobuf:"varint,5,opt,name=Type,json=type,enum=protoc.MessageReq_Message_ContentType" json:"Type,omitempty"`
	IsGroupMessage bool                           `protobuf:"varint,6,opt,name=IsGroupMessage,json=isGroupMessage" json:"IsGroupMessage,omitempty"`
}

func (m *MessageReq_Message) Reset()                    { *m = MessageReq_Message{} }
func (m *MessageReq_Message) String() string            { return proto.CompactTextString(m) }
func (*MessageReq_Message) ProtoMessage()               {}
func (*MessageReq_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type OnlineReq struct {
	Items []*OnlineReq_Item `protobuf:"bytes,1,rep,name=Items,json=items" json:"Items,omitempty"`
}

func (m *OnlineReq) Reset()                    { *m = OnlineReq{} }
func (m *OnlineReq) String() string            { return proto.CompactTextString(m) }
func (*OnlineReq) ProtoMessage()               {}
func (*OnlineReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OnlineReq) GetItems() []*OnlineReq_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type OnlineReq_Item struct {
	Id       []byte `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	IsOnline bool   `protobuf:"varint,2,opt,name=IsOnline,json=isOnline" json:"IsOnline,omitempty"`
	IsGroup  bool   `protobuf:"varint,3,opt,name=IsGroup,json=isGroup" json:"IsGroup,omitempty"`
}

func (m *OnlineReq_Item) Reset()                    { *m = OnlineReq_Item{} }
func (m *OnlineReq_Item) String() string            { return proto.CompactTextString(m) }
func (*OnlineReq_Item) ProtoMessage()               {}
func (*OnlineReq_Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type GroupReq struct {
	Groups []*GroupReq_Group `protobuf:"bytes,1,rep,name=Groups,json=groups" json:"Groups,omitempty"`
}

func (m *GroupReq) Reset()                    { *m = GroupReq{} }
func (m *GroupReq) String() string            { return proto.CompactTextString(m) }
func (*GroupReq) ProtoMessage()               {}
func (*GroupReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GroupReq) GetGroups() []*GroupReq_Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type GroupReq_Group struct {
	Id        []byte   `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=Title,json=title" json:"Title,omitempty"`
	Logo      string   `protobuf:"bytes,3,opt,name=Logo,json=logo" json:"Logo,omitempty"`
	Members   [][]byte `protobuf:"bytes,4,rep,name=Members,json=members,proto3" json:"Members,omitempty"`
	Introduce string   `protobuf:"bytes,5,opt,name=Introduce,json=introduce" json:"Introduce,omitempty"`
}

func (m *GroupReq_Group) Reset()                    { *m = GroupReq_Group{} }
func (m *GroupReq_Group) String() string            { return proto.CompactTextString(m) }
func (*GroupReq_Group) ProtoMessage()               {}
func (*GroupReq_Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func init() {
	proto.RegisterType((*MessageReq)(nil), "protoc.MessageReq")
	proto.RegisterType((*MessageReq_Message)(nil), "protoc.MessageReq.Message")
	proto.RegisterType((*OnlineReq)(nil), "protoc.OnlineReq")
	proto.RegisterType((*OnlineReq_Item)(nil), "protoc.OnlineReq.Item")
	proto.RegisterType((*GroupReq)(nil), "protoc.GroupReq")
	proto.RegisterType((*GroupReq_Group)(nil), "protoc.GroupReq.Group")
	proto.RegisterEnum("protoc.MessageReq_Message_ContentType", MessageReq_Message_ContentType_name, MessageReq_Message_ContentType_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xbe, 0x77, 0xc7, 0xad, 0x10, 0x4b, 0x29, 0x42, 0xf4, 0x50, 0x74, 0x30, 0x3e, 0x14,
	0x1d, 0x5c, 0x28, 0xa5, 0xd7, 0xd2, 0x16, 0x81, 0x4d, 0xc2, 0x62, 0x72, 0x8f, 0xe5, 0x8d, 0x59,
	0xb0, 0xb4, 0x8e, 0xb4, 0x86, 0xf8, 0x17, 0xe4, 0x17, 0xe5, 0x96, 0xbf, 0x95, 0x7b, 0xf6, 0x4b,
	0x4e, 0x82, 0xc9, 0x45, 0x9a, 0xf7, 0xf6, 0xcd, 0xe8, 0xed, 0x1b, 0xc1, 0xa7, 0x96, 0x0d, 0xc3,
	0xf5, 0x96, 0x55, 0xfb, 0x5e, 0x48, 0x41, 0x62, 0xf3, 0x6a, 0xca, 0x47, 0x1f, 0x60, 0x69, 0x4f,
	0x28, 0xbb, 0x25, 0x3f, 0x01, 0x39, 0x34, 0xe4, 0xde, 0xb7, 0x60, 0x36, 0x99, 0x17, 0xb6, 0xa1,
	0xa9, 0x5e, 0x54, 0xa7, 0x12, 0xb9, 0x99, 0x43, 0xf1, 0xe4, 0x41, 0xe2, 0x58, 0x92, 0x82, 0x5f,
	0x6f, 0x54, 0xb7, 0x37, 0xfb, 0x48, 0x7d, 0xbe, 0x21, 0x04, 0xc2, 0x7f, 0xbd, 0x68, 0x73, 0xdf,
	0x30, 0xe1, 0x8d, 0xaa, 0xb5, 0x66, 0x25, 0xf2, 0xc0, 0x6a, 0x94, 0x9d, 0x1c, 0x92, 0x3f, 0xa2,
	0x93, 0xac, 0x93, 0x79, 0xa8, 0x48, 0x4c, 0x93, 0xc6, 0x42, 0xf2, 0x1b, 0xc2, 0xd5, 0x71, 0xcf,
	0xf2, 0x48, 0xd1, 0xe9, 0x7c, 0xfa, 0xbe, 0x9b, 0xca, 0x0d, 0xd0, 0x6a, 0x1a, 0x4a, 0xf5, 0x24,
	0x53, 0x48, 0xeb, 0xe1, 0x7f, 0x2f, 0x0e, 0x7b, 0xa7, 0xc9, 0x63, 0x35, 0x05, 0xd1, 0x94, 0xbf,
	0x61, 0xcb, 0x5f, 0x30, 0x79, 0xd5, 0x4c, 0x90, 0xfa, 0x24, 0xbb, 0x93, 0xd9, 0x07, 0x92, 0x40,
	0x70, 0xc9, 0x9b, 0xcc, 0x23, 0x13, 0x48, 0xfe, 0xb6, 0x42, 0x72, 0xd1, 0x65, 0x3e, 0xc1, 0x10,
	0x5d, 0x09, 0xde, 0xb0, 0x2c, 0x28, 0xef, 0x3d, 0xc0, 0x17, 0xdd, 0x8e, 0x77, 0x26, 0xbd, 0xef,
	0x10, 0xd5, 0x92, 0xb5, 0x63, 0x74, 0x5f, 0x46, 0xb3, 0x27, 0x45, 0xa5, 0x8f, 0x69, 0xc4, 0xb5,
	0xa8, 0x58, 0x40, 0xa8, 0xe1, 0x59, 0x5e, 0x05, 0xa0, 0x7a, 0xb0, 0x2d, 0x26, 0x33, 0x44, 0x11,
	0x77, 0x58, 0xe7, 0xe4, 0x6e, 0x64, 0xc2, 0x43, 0x34, 0x71, 0x57, 0x29, 0x1f, 0x3c, 0x40, 0xa6,
	0xd2, 0x46, 0x2a, 0x88, 0x4d, 0x7d, 0xe6, 0x64, 0x54, 0xb8, 0x22, 0xde, 0x1a, 0x55, 0x71, 0x84,
	0xc8, 0x10, 0x67, 0x5e, 0x3e, 0x43, 0xb4, 0xe2, 0x72, 0x67, 0x8d, 0x60, 0x1a, 0x49, 0x0d, 0xf4,
	0x46, 0x17, 0x62, 0x6b, 0xf7, 0x87, 0x69, 0xb8, 0x53, 0xb5, 0x76, 0xb6, 0x64, 0xed, 0x9a, 0xf5,
	0x83, 0xda, 0x60, 0xa0, 0xda, 0x93, 0xd6, 0x42, 0xf2, 0x15, 0x70, 0xdd, 0xc9, 0x5e, 0x6c, 0x0e,
	0x8d, 0x5d, 0x23, 0xa6, 0x98, 0x8f, 0xc4, 0xda, 0xfe, 0x88, 0x3f, 0x9e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0x63, 0x9f, 0xd3, 0xa0, 0x02, 0x00, 0x00,
}
