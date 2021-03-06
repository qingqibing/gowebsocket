// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/remarketing_action.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A remarketing action. A snippet of JavaScript code that will collect the
// product id and the type of page people visited (product page, shopping cart
// page, purchase page, general site visit) on an advertiser's website.
type RemarketingAction struct {
	// The resource name of the remarketing action.
	// Remarketing action resource names have the form:
	//
	// `customers/{customer_id}/remarketingActions/{remarketing_action_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the remarketing action.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the remarketing action.
	//
	// This field is required and should not be empty when creating new
	// remarketing actions.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The snippets used for tracking remarketing actions.
	TagSnippets          []*common.TagSnippet `protobuf:"bytes,4,rep,name=tag_snippets,json=tagSnippets,proto3" json:"tag_snippets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RemarketingAction) Reset()         { *m = RemarketingAction{} }
func (m *RemarketingAction) String() string { return proto.CompactTextString(m) }
func (*RemarketingAction) ProtoMessage()    {}
func (*RemarketingAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ce490472884744, []int{0}
}

func (m *RemarketingAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingAction.Unmarshal(m, b)
}
func (m *RemarketingAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingAction.Marshal(b, m, deterministic)
}
func (m *RemarketingAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingAction.Merge(m, src)
}
func (m *RemarketingAction) XXX_Size() int {
	return xxx_messageInfo_RemarketingAction.Size(m)
}
func (m *RemarketingAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingAction.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingAction proto.InternalMessageInfo

func (m *RemarketingAction) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *RemarketingAction) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RemarketingAction) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RemarketingAction) GetTagSnippets() []*common.TagSnippet {
	if m != nil {
		return m.TagSnippets
	}
	return nil
}

func init() {
	proto.RegisterType((*RemarketingAction)(nil), "google.ads.googleads.v1.resources.RemarketingAction")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/remarketing_action.proto", fileDescriptor_62ce490472884744)
}

var fileDescriptor_62ce490472884744 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xb1, 0x13, 0x06, 0x73, 0xb2, 0xc3, 0x7c, 0x18, 0x21, 0x0b, 0x23, 0xd9, 0x08, 0x84,
	0x0d, 0xa4, 0x38, 0x1b, 0x3b, 0xa8, 0x27, 0xe7, 0x12, 0x5a, 0x68, 0x09, 0x4e, 0xf1, 0xa1, 0x18,
	0x82, 0x62, 0xab, 0x42, 0x34, 0x96, 0x8c, 0xa4, 0xa4, 0x0f, 0xd0, 0x37, 0xe9, 0xb1, 0x8f, 0xd2,
	0x07, 0xe9, 0xa1, 0x4f, 0x51, 0x6c, 0xd9, 0x6a, 0x21, 0xa4, 0xbd, 0xfd, 0xed, 0xef, 0xf7, 0xff,
	0xf4, 0xff, 0x3e, 0xc9, 0x43, 0x54, 0x08, 0xba, 0x25, 0x10, 0x67, 0x0a, 0x1a, 0x59, 0xaa, 0x7d,
	0x00, 0x25, 0x51, 0x62, 0x27, 0x53, 0xa2, 0xa0, 0x24, 0x39, 0x96, 0x37, 0x44, 0x33, 0x4e, 0xd7,
	0x38, 0xd5, 0x4c, 0x70, 0x50, 0x48, 0xa1, 0x85, 0x3f, 0x32, 0x06, 0x80, 0x33, 0x05, 0xac, 0x17,
	0xec, 0x03, 0x60, 0xbd, 0xfd, 0xe9, 0xb1, 0xf6, 0xa9, 0xc8, 0x73, 0xc1, 0xa1, 0xc6, 0x74, 0xad,
	0x38, 0x2b, 0x0a, 0xa2, 0x4d, 0xd3, 0xfe, 0xa0, 0x71, 0x14, 0x0c, 0x62, 0xce, 0x85, 0xc6, 0xe5,
	0x89, 0xaa, 0xae, 0xfe, 0xa8, 0xab, 0xd5, 0xd7, 0x66, 0x77, 0x0d, 0x6f, 0x25, 0x2e, 0x0a, 0x22,
	0xeb, 0xfa, 0xcf, 0x27, 0xc7, 0xfb, 0x1a, 0xbd, 0xe6, 0x0d, 0xab, 0xb8, 0xfe, 0x2f, 0xef, 0x4b,
	0x13, 0x69, 0xcd, 0x71, 0x4e, 0x7a, 0xce, 0xd0, 0x99, 0x7c, 0x8e, 0xba, 0xcd, 0xcf, 0x0b, 0x9c,
	0x13, 0xff, 0x8f, 0xe7, 0xb2, 0xac, 0xe7, 0x0e, 0x9d, 0x49, 0x67, 0xf6, 0xbd, 0x9e, 0x07, 0x34,
	0xe7, 0x80, 0x53, 0xae, 0xff, 0xff, 0x8b, 0xf1, 0x76, 0x47, 0x22, 0x97, 0x65, 0xfe, 0xd4, 0x6b,
	0x57, 0x8d, 0x5a, 0x15, 0x3e, 0x38, 0xc0, 0x57, 0x5a, 0x32, 0x4e, 0x0d, 0x5f, 0x91, 0xfe, 0xb9,
	0xd7, 0x7d, 0x33, 0xac, 0xea, 0xb5, 0x87, 0xad, 0x49, 0x67, 0xf6, 0x1b, 0x1c, 0xdb, 0xa1, 0x59,
	0x10, 0xb8, 0xc4, 0x74, 0x65, 0x2c, 0x51, 0x47, 0x5b, 0xad, 0xe6, 0x77, 0xae, 0x37, 0x4e, 0x45,
	0x0e, 0x3e, 0xbc, 0x82, 0xf9, 0xb7, 0x83, 0x7d, 0x2c, 0xcb, 0x94, 0x4b, 0xe7, 0xea, 0xac, 0x36,
	0x53, 0xb1, 0xc5, 0x9c, 0x02, 0x21, 0x29, 0xa4, 0x84, 0x57, 0x33, 0x34, 0x97, 0x55, 0x30, 0xf5,
	0xce, 0xd3, 0x38, 0xb1, 0xea, 0xde, 0x6d, 0x2d, 0xc2, 0xf0, 0xc1, 0x1d, 0x2d, 0x4c, 0xcb, 0x30,
	0x53, 0xc0, 0xc8, 0x52, 0xc5, 0x01, 0x88, 0x1a, 0xf2, 0xb1, 0x61, 0x92, 0x30, 0x53, 0x89, 0x65,
	0x92, 0x38, 0x48, 0x2c, 0xf3, 0xec, 0x8e, 0x4d, 0x01, 0xa1, 0x30, 0x53, 0x08, 0x59, 0x0a, 0xa1,
	0x38, 0x40, 0xc8, 0x72, 0x9b, 0x4f, 0x55, 0xd8, 0xbf, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x14, 0xed, 0x2e, 0xc6, 0x02, 0x00, 0x00,
}
