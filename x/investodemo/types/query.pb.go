// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: investodemo/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # 3
type QueryGetCompanyRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetCompanyRequest) Reset()         { *m = QueryGetCompanyRequest{} }
func (m *QueryGetCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompanyRequest) ProtoMessage()    {}
func (*QueryGetCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dee429f6201b05a, []int{0}
}
func (m *QueryGetCompanyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompanyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompanyRequest.Merge(m, src)
}
func (m *QueryGetCompanyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompanyRequest proto.InternalMessageInfo

func (m *QueryGetCompanyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetCompanyResponse struct {
	Company *Company `protobuf:"bytes,1,opt,name=Company,proto3" json:"Company,omitempty"`
}

func (m *QueryGetCompanyResponse) Reset()         { *m = QueryGetCompanyResponse{} }
func (m *QueryGetCompanyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompanyResponse) ProtoMessage()    {}
func (*QueryGetCompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dee429f6201b05a, []int{1}
}
func (m *QueryGetCompanyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompanyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompanyResponse.Merge(m, src)
}
func (m *QueryGetCompanyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompanyResponse proto.InternalMessageInfo

func (m *QueryGetCompanyResponse) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

type QueryAllCompanyRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompanyRequest) Reset()         { *m = QueryAllCompanyRequest{} }
func (m *QueryAllCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompanyRequest) ProtoMessage()    {}
func (*QueryAllCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dee429f6201b05a, []int{2}
}
func (m *QueryAllCompanyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompanyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompanyRequest.Merge(m, src)
}
func (m *QueryAllCompanyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompanyRequest proto.InternalMessageInfo

func (m *QueryAllCompanyRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCompanyResponse struct {
	Company    []*Company          `protobuf:"bytes,1,rep,name=Company,proto3" json:"Company,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompanyResponse) Reset()         { *m = QueryAllCompanyResponse{} }
func (m *QueryAllCompanyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompanyResponse) ProtoMessage()    {}
func (*QueryAllCompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dee429f6201b05a, []int{3}
}
func (m *QueryAllCompanyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompanyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompanyResponse.Merge(m, src)
}
func (m *QueryAllCompanyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompanyResponse proto.InternalMessageInfo

func (m *QueryAllCompanyResponse) GetCompany() []*Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *QueryAllCompanyResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCompanyRequest)(nil), "faddat.investodemo.investodemo.QueryGetCompanyRequest")
	proto.RegisterType((*QueryGetCompanyResponse)(nil), "faddat.investodemo.investodemo.QueryGetCompanyResponse")
	proto.RegisterType((*QueryAllCompanyRequest)(nil), "faddat.investodemo.investodemo.QueryAllCompanyRequest")
	proto.RegisterType((*QueryAllCompanyResponse)(nil), "faddat.investodemo.investodemo.QueryAllCompanyResponse")
}

func init() { proto.RegisterFile("investodemo/query.proto", fileDescriptor_6dee429f6201b05a) }

var fileDescriptor_6dee429f6201b05a = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbd, 0x4e, 0xeb, 0x30,
	0x1c, 0xc5, 0xeb, 0x5c, 0xdd, 0x7b, 0x75, 0x7d, 0x25, 0x06, 0x0f, 0x14, 0x2a, 0x14, 0xa1, 0x0c,
	0xb4, 0x20, 0xb0, 0xd5, 0x82, 0x60, 0x2e, 0x48, 0x74, 0x61, 0x80, 0x8e, 0x88, 0x01, 0xa7, 0x31,
	0xc1, 0x52, 0x12, 0xa7, 0xb5, 0x5b, 0x51, 0x21, 0x16, 0x9e, 0x00, 0x89, 0x67, 0x60, 0x45, 0x3c,
	0x06, 0x63, 0x25, 0x16, 0x46, 0xd4, 0xb2, 0xf1, 0x12, 0xa8, 0xb1, 0x0b, 0xe9, 0x87, 0x68, 0x61,
	0x8b, 0xe2, 0x73, 0x8e, 0x7f, 0xc7, 0x7f, 0x1b, 0x66, 0x79, 0xd4, 0x62, 0x52, 0x09, 0x8f, 0x85,
	0x82, 0xd4, 0x9b, 0xac, 0xd1, 0xc6, 0x71, 0x43, 0x28, 0x81, 0xec, 0x33, 0xea, 0x79, 0x54, 0xe1,
	0xd4, 0x7a, 0xfa, 0x3b, 0xb7, 0xe4, 0x0b, 0xe1, 0x07, 0x8c, 0xd0, 0x98, 0x13, 0x1a, 0x45, 0x42,
	0x51, 0xc5, 0x45, 0x24, 0xb5, 0x3b, 0xb7, 0x56, 0x13, 0x32, 0x14, 0x92, 0xb8, 0x54, 0x32, 0x1d,
	0x4b, 0x5a, 0x45, 0x97, 0x29, 0x5a, 0x24, 0x31, 0xf5, 0x79, 0x94, 0x88, 0x8d, 0x76, 0x31, 0x8d,
	0x50, 0x13, 0x61, 0x4c, 0x23, 0x03, 0xe1, 0x14, 0xe0, 0xfc, 0x51, 0xdf, 0x5c, 0x61, 0x6a, 0x4f,
	0x2f, 0x54, 0x59, 0xbd, 0xc9, 0xa4, 0x42, 0x73, 0xd0, 0xe2, 0xde, 0x02, 0x58, 0x06, 0x85, 0x7f,
	0x55, 0x8b, 0x7b, 0xce, 0x09, 0xcc, 0x8e, 0x29, 0x65, 0x2c, 0x22, 0xc9, 0x50, 0x19, 0xfe, 0x35,
	0xbf, 0x12, 0xfd, 0xff, 0x52, 0x1e, 0x7f, 0xdd, 0x0d, 0x0f, 0x12, 0x06, 0x3e, 0xe7, 0xd4, 0x70,
	0x94, 0x83, 0x60, 0x84, 0x63, 0x1f, 0xc2, 0xcf, 0x42, 0x26, 0x7f, 0x05, 0xeb, 0xf6, 0xb8, 0xdf,
	0x1e, 0xeb, 0x43, 0x35, 0xed, 0xf1, 0x21, 0xf5, 0x99, 0xf1, 0x56, 0x53, 0x4e, 0xe7, 0x0e, 0x98,
	0x02, 0xe9, 0x2d, 0x26, 0x15, 0xf8, 0xf5, 0x93, 0x02, 0xa8, 0x32, 0x84, 0x69, 0x99, 0x63, 0x98,
	0x86, 0xa9, 0xf7, 0x4f, 0x73, 0x96, 0xde, 0x2c, 0xf8, 0x3b, 0xe1, 0x44, 0x0f, 0xe0, 0x03, 0x0b,
	0x6d, 0x4f, 0x03, 0x9a, 0x3c, 0xc5, 0xdc, 0xce, 0xb7, 0x7d, 0x1a, 0xc9, 0xd9, 0xba, 0x7e, 0x7a,
	0xbd, 0xb5, 0x30, 0x5a, 0x27, 0x3a, 0x80, 0xa4, 0xef, 0xd0, 0x84, 0xfb, 0x44, 0x2e, 0xb9, 0x77,
	0x85, 0xee, 0x01, 0x84, 0x26, 0xa9, 0x1c, 0x04, 0x33, 0x52, 0x8f, 0xcd, 0x7c, 0x46, 0xea, 0xf1,
	0x41, 0x3a, 0x24, 0xa1, 0x5e, 0x45, 0xf9, 0x19, 0xa9, 0x77, 0x0f, 0x1e, 0xbb, 0x36, 0xe8, 0x74,
	0x6d, 0xf0, 0xd2, 0xb5, 0xc1, 0x4d, 0xcf, 0xce, 0x74, 0x7a, 0x76, 0xe6, 0xb9, 0x67, 0x67, 0x8e,
	0x4b, 0x3e, 0x57, 0xe7, 0x4d, 0x17, 0xd7, 0x44, 0x38, 0x12, 0xb6, 0x91, 0x24, 0x5c, 0x0c, 0xe5,
	0xa9, 0x76, 0xcc, 0xa4, 0xfb, 0x27, 0x79, 0x54, 0x9b, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65,
	0xfd, 0x23, 0x70, 0xf4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	Company(ctx context.Context, in *QueryGetCompanyRequest, opts ...grpc.CallOption) (*QueryGetCompanyResponse, error)
	CompanyAll(ctx context.Context, in *QueryAllCompanyRequest, opts ...grpc.CallOption) (*QueryAllCompanyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Company(ctx context.Context, in *QueryGetCompanyRequest, opts ...grpc.CallOption) (*QueryGetCompanyResponse, error) {
	out := new(QueryGetCompanyResponse)
	err := c.cc.Invoke(ctx, "/faddat.investodemo.investodemo.Query/Company", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CompanyAll(ctx context.Context, in *QueryAllCompanyRequest, opts ...grpc.CallOption) (*QueryAllCompanyResponse, error) {
	out := new(QueryAllCompanyResponse)
	err := c.cc.Invoke(ctx, "/faddat.investodemo.investodemo.Query/CompanyAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Company(context.Context, *QueryGetCompanyRequest) (*QueryGetCompanyResponse, error)
	CompanyAll(context.Context, *QueryAllCompanyRequest) (*QueryAllCompanyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Company(ctx context.Context, req *QueryGetCompanyRequest) (*QueryGetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Company not implemented")
}
func (*UnimplementedQueryServer) CompanyAll(ctx context.Context, req *QueryAllCompanyRequest) (*QueryAllCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Company_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Company(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faddat.investodemo.investodemo.Query/Company",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Company(ctx, req.(*QueryGetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CompanyAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CompanyAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faddat.investodemo.investodemo.Query/CompanyAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CompanyAll(ctx, req.(*QueryAllCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faddat.investodemo.investodemo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Company",
			Handler:    _Query_Company_Handler,
		},
		{
			MethodName: "CompanyAll",
			Handler:    _Query_CompanyAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "investodemo/query.proto",
}

func (m *QueryGetCompanyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompanyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompanyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCompanyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompanyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompanyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Company != nil {
		{
			size, err := m.Company.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCompanyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompanyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompanyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCompanyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompanyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompanyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Company) > 0 {
		for iNdEx := len(m.Company) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Company[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCompanyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCompanyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Company != nil {
		l = m.Company.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCompanyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCompanyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Company) > 0 {
		for _, e := range m.Company {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCompanyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCompanyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompanyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCompanyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCompanyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompanyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Company", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Company == nil {
				m.Company = &Company{}
			}
			if err := m.Company.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCompanyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCompanyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompanyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCompanyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCompanyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompanyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Company", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Company = append(m.Company, &Company{})
			if err := m.Company[len(m.Company)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
