// Code generated by protoc-gen-go.
// source: master.proto
// DO NOT EDIT!

/*
Package master_pb is a generated protocol buffer package.

It is generated from these files:
	master.proto

It has these top-level messages:
	Heartbeat
	HeartbeatResponse
	VolumeInformationMessage
	Empty
	SuperBlockExtra
	ClientListenRequest
	VolumeLocation
	LookupVolumeRequest
	LookupVolumeResponse
	Location
*/
package master_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Heartbeat struct {
	Ip             string                      `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port           uint32                      `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	PublicUrl      string                      `protobuf:"bytes,3,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
	MaxVolumeCount uint32                      `protobuf:"varint,4,opt,name=max_volume_count,json=maxVolumeCount" json:"max_volume_count,omitempty"`
	MaxFileKey     uint64                      `protobuf:"varint,5,opt,name=max_file_key,json=maxFileKey" json:"max_file_key,omitempty"`
	DataCenter     string                      `protobuf:"bytes,6,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Rack           string                      `protobuf:"bytes,7,opt,name=rack" json:"rack,omitempty"`
	AdminPort      uint32                      `protobuf:"varint,8,opt,name=admin_port,json=adminPort" json:"admin_port,omitempty"`
	Volumes        []*VolumeInformationMessage `protobuf:"bytes,9,rep,name=volumes" json:"volumes,omitempty"`
	// delta volume ids
	NewVids     []uint32 `protobuf:"varint,10,rep,packed,name=new_vids,json=newVids" json:"new_vids,omitempty"`
	DeletedVids []uint32 `protobuf:"varint,11,rep,packed,name=deleted_vids,json=deletedVids" json:"deleted_vids,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Heartbeat) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Heartbeat) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Heartbeat) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func (m *Heartbeat) GetMaxVolumeCount() uint32 {
	if m != nil {
		return m.MaxVolumeCount
	}
	return 0
}

func (m *Heartbeat) GetMaxFileKey() uint64 {
	if m != nil {
		return m.MaxFileKey
	}
	return 0
}

func (m *Heartbeat) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Heartbeat) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Heartbeat) GetAdminPort() uint32 {
	if m != nil {
		return m.AdminPort
	}
	return 0
}

func (m *Heartbeat) GetVolumes() []*VolumeInformationMessage {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Heartbeat) GetNewVids() []uint32 {
	if m != nil {
		return m.NewVids
	}
	return nil
}

func (m *Heartbeat) GetDeletedVids() []uint32 {
	if m != nil {
		return m.DeletedVids
	}
	return nil
}

type HeartbeatResponse struct {
	VolumeSizeLimit uint64 `protobuf:"varint,1,opt,name=volumeSizeLimit" json:"volumeSizeLimit,omitempty"`
	SecretKey       string `protobuf:"bytes,2,opt,name=secretKey" json:"secretKey,omitempty"`
	Leader          string `protobuf:"bytes,3,opt,name=leader" json:"leader,omitempty"`
}

func (m *HeartbeatResponse) Reset()                    { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()               {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HeartbeatResponse) GetVolumeSizeLimit() uint64 {
	if m != nil {
		return m.VolumeSizeLimit
	}
	return 0
}

func (m *HeartbeatResponse) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *HeartbeatResponse) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

type VolumeInformationMessage struct {
	Id               uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Size             uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Collection       string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
	FileCount        uint64 `protobuf:"varint,4,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	DeleteCount      uint64 `protobuf:"varint,5,opt,name=delete_count,json=deleteCount" json:"delete_count,omitempty"`
	DeletedByteCount uint64 `protobuf:"varint,6,opt,name=deleted_byte_count,json=deletedByteCount" json:"deleted_byte_count,omitempty"`
	ReadOnly         bool   `protobuf:"varint,7,opt,name=read_only,json=readOnly" json:"read_only,omitempty"`
	ReplicaPlacement uint32 `protobuf:"varint,8,opt,name=replica_placement,json=replicaPlacement" json:"replica_placement,omitempty"`
	Version          uint32 `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	Ttl              uint32 `protobuf:"varint,10,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *VolumeInformationMessage) Reset()                    { *m = VolumeInformationMessage{} }
func (m *VolumeInformationMessage) String() string            { return proto.CompactTextString(m) }
func (*VolumeInformationMessage) ProtoMessage()               {}
func (*VolumeInformationMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VolumeInformationMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VolumeInformationMessage) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VolumeInformationMessage) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *VolumeInformationMessage) GetFileCount() uint64 {
	if m != nil {
		return m.FileCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetDeleteCount() uint64 {
	if m != nil {
		return m.DeleteCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetDeletedByteCount() uint64 {
	if m != nil {
		return m.DeletedByteCount
	}
	return 0
}

func (m *VolumeInformationMessage) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *VolumeInformationMessage) GetReplicaPlacement() uint32 {
	if m != nil {
		return m.ReplicaPlacement
	}
	return 0
}

func (m *VolumeInformationMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeInformationMessage) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SuperBlockExtra struct {
	ErasureCoding *SuperBlockExtra_ErasureCoding `protobuf:"bytes,1,opt,name=erasure_coding,json=erasureCoding" json:"erasure_coding,omitempty"`
}

func (m *SuperBlockExtra) Reset()                    { *m = SuperBlockExtra{} }
func (m *SuperBlockExtra) String() string            { return proto.CompactTextString(m) }
func (*SuperBlockExtra) ProtoMessage()               {}
func (*SuperBlockExtra) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SuperBlockExtra) GetErasureCoding() *SuperBlockExtra_ErasureCoding {
	if m != nil {
		return m.ErasureCoding
	}
	return nil
}

type SuperBlockExtra_ErasureCoding struct {
	Data      uint32   `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
	Parity    uint32   `protobuf:"varint,2,opt,name=parity" json:"parity,omitempty"`
	VolumeIds []uint32 `protobuf:"varint,3,rep,packed,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *SuperBlockExtra_ErasureCoding) Reset()         { *m = SuperBlockExtra_ErasureCoding{} }
func (m *SuperBlockExtra_ErasureCoding) String() string { return proto.CompactTextString(m) }
func (*SuperBlockExtra_ErasureCoding) ProtoMessage()    {}
func (*SuperBlockExtra_ErasureCoding) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *SuperBlockExtra_ErasureCoding) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *SuperBlockExtra_ErasureCoding) GetParity() uint32 {
	if m != nil {
		return m.Parity
	}
	return 0
}

func (m *SuperBlockExtra_ErasureCoding) GetVolumeIds() []uint32 {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

type ClientListenRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClientListenRequest) Reset()                    { *m = ClientListenRequest{} }
func (m *ClientListenRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientListenRequest) ProtoMessage()               {}
func (*ClientListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientListenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VolumeLocation struct {
	Url         string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	PublicUrl   string   `protobuf:"bytes,2,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
	NewVids     []uint32 `protobuf:"varint,3,rep,packed,name=new_vids,json=newVids" json:"new_vids,omitempty"`
	DeletedVids []uint32 `protobuf:"varint,4,rep,packed,name=deleted_vids,json=deletedVids" json:"deleted_vids,omitempty"`
}

func (m *VolumeLocation) Reset()                    { *m = VolumeLocation{} }
func (m *VolumeLocation) String() string            { return proto.CompactTextString(m) }
func (*VolumeLocation) ProtoMessage()               {}
func (*VolumeLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VolumeLocation) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VolumeLocation) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func (m *VolumeLocation) GetNewVids() []uint32 {
	if m != nil {
		return m.NewVids
	}
	return nil
}

func (m *VolumeLocation) GetDeletedVids() []uint32 {
	if m != nil {
		return m.DeletedVids
	}
	return nil
}

type LookupVolumeRequest struct {
	VolumeIds  []string `protobuf:"bytes,1,rep,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
	Collection string   `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
}

func (m *LookupVolumeRequest) Reset()                    { *m = LookupVolumeRequest{} }
func (m *LookupVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupVolumeRequest) ProtoMessage()               {}
func (*LookupVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LookupVolumeRequest) GetVolumeIds() []string {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

func (m *LookupVolumeRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type LookupVolumeResponse struct {
	VolumeIdLocations []*LookupVolumeResponse_VolumeIdLocation `protobuf:"bytes,1,rep,name=volume_id_locations,json=volumeIdLocations" json:"volume_id_locations,omitempty"`
}

func (m *LookupVolumeResponse) Reset()                    { *m = LookupVolumeResponse{} }
func (m *LookupVolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupVolumeResponse) ProtoMessage()               {}
func (*LookupVolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LookupVolumeResponse) GetVolumeIdLocations() []*LookupVolumeResponse_VolumeIdLocation {
	if m != nil {
		return m.VolumeIdLocations
	}
	return nil
}

type LookupVolumeResponse_VolumeIdLocation struct {
	VolumeId  string      `protobuf:"bytes,1,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	Locations []*Location `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
	Error     string      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *LookupVolumeResponse_VolumeIdLocation) Reset()         { *m = LookupVolumeResponse_VolumeIdLocation{} }
func (m *LookupVolumeResponse_VolumeIdLocation) String() string { return proto.CompactTextString(m) }
func (*LookupVolumeResponse_VolumeIdLocation) ProtoMessage()    {}
func (*LookupVolumeResponse_VolumeIdLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *LookupVolumeResponse_VolumeIdLocation) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *LookupVolumeResponse_VolumeIdLocation) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *LookupVolumeResponse_VolumeIdLocation) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Location struct {
	Url       string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	PublicUrl string `protobuf:"bytes,2,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Location) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Location) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Heartbeat)(nil), "master_pb.Heartbeat")
	proto.RegisterType((*HeartbeatResponse)(nil), "master_pb.HeartbeatResponse")
	proto.RegisterType((*VolumeInformationMessage)(nil), "master_pb.VolumeInformationMessage")
	proto.RegisterType((*Empty)(nil), "master_pb.Empty")
	proto.RegisterType((*SuperBlockExtra)(nil), "master_pb.SuperBlockExtra")
	proto.RegisterType((*SuperBlockExtra_ErasureCoding)(nil), "master_pb.SuperBlockExtra.ErasureCoding")
	proto.RegisterType((*ClientListenRequest)(nil), "master_pb.ClientListenRequest")
	proto.RegisterType((*VolumeLocation)(nil), "master_pb.VolumeLocation")
	proto.RegisterType((*LookupVolumeRequest)(nil), "master_pb.LookupVolumeRequest")
	proto.RegisterType((*LookupVolumeResponse)(nil), "master_pb.LookupVolumeResponse")
	proto.RegisterType((*LookupVolumeResponse_VolumeIdLocation)(nil), "master_pb.LookupVolumeResponse.VolumeIdLocation")
	proto.RegisterType((*Location)(nil), "master_pb.Location")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Seaweed service

type SeaweedClient interface {
	SendHeartbeat(ctx context.Context, opts ...grpc.CallOption) (Seaweed_SendHeartbeatClient, error)
	KeepConnected(ctx context.Context, opts ...grpc.CallOption) (Seaweed_KeepConnectedClient, error)
	LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error)
}

type seaweedClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedClient(cc *grpc.ClientConn) SeaweedClient {
	return &seaweedClient{cc}
}

func (c *seaweedClient) SendHeartbeat(ctx context.Context, opts ...grpc.CallOption) (Seaweed_SendHeartbeatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Seaweed_serviceDesc.Streams[0], c.cc, "/master_pb.Seaweed/SendHeartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedSendHeartbeatClient{stream}
	return x, nil
}

type Seaweed_SendHeartbeatClient interface {
	Send(*Heartbeat) error
	Recv() (*HeartbeatResponse, error)
	grpc.ClientStream
}

type seaweedSendHeartbeatClient struct {
	grpc.ClientStream
}

func (x *seaweedSendHeartbeatClient) Send(m *Heartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedSendHeartbeatClient) Recv() (*HeartbeatResponse, error) {
	m := new(HeartbeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedClient) KeepConnected(ctx context.Context, opts ...grpc.CallOption) (Seaweed_KeepConnectedClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Seaweed_serviceDesc.Streams[1], c.cc, "/master_pb.Seaweed/KeepConnected", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedKeepConnectedClient{stream}
	return x, nil
}

type Seaweed_KeepConnectedClient interface {
	Send(*ClientListenRequest) error
	Recv() (*VolumeLocation, error)
	grpc.ClientStream
}

type seaweedKeepConnectedClient struct {
	grpc.ClientStream
}

func (x *seaweedKeepConnectedClient) Send(m *ClientListenRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedKeepConnectedClient) Recv() (*VolumeLocation, error) {
	m := new(VolumeLocation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedClient) LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error) {
	out := new(LookupVolumeResponse)
	err := grpc.Invoke(ctx, "/master_pb.Seaweed/LookupVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Seaweed service

type SeaweedServer interface {
	SendHeartbeat(Seaweed_SendHeartbeatServer) error
	KeepConnected(Seaweed_KeepConnectedServer) error
	LookupVolume(context.Context, *LookupVolumeRequest) (*LookupVolumeResponse, error)
}

func RegisterSeaweedServer(s *grpc.Server, srv SeaweedServer) {
	s.RegisterService(&_Seaweed_serviceDesc, srv)
}

func _Seaweed_SendHeartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedServer).SendHeartbeat(&seaweedSendHeartbeatServer{stream})
}

type Seaweed_SendHeartbeatServer interface {
	Send(*HeartbeatResponse) error
	Recv() (*Heartbeat, error)
	grpc.ServerStream
}

type seaweedSendHeartbeatServer struct {
	grpc.ServerStream
}

func (x *seaweedSendHeartbeatServer) Send(m *HeartbeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedSendHeartbeatServer) Recv() (*Heartbeat, error) {
	m := new(Heartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Seaweed_KeepConnected_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedServer).KeepConnected(&seaweedKeepConnectedServer{stream})
}

type Seaweed_KeepConnectedServer interface {
	Send(*VolumeLocation) error
	Recv() (*ClientListenRequest, error)
	grpc.ServerStream
}

type seaweedKeepConnectedServer struct {
	grpc.ServerStream
}

func (x *seaweedKeepConnectedServer) Send(m *VolumeLocation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedKeepConnectedServer) Recv() (*ClientListenRequest, error) {
	m := new(ClientListenRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Seaweed_LookupVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedServer).LookupVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master_pb.Seaweed/LookupVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedServer).LookupVolume(ctx, req.(*LookupVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Seaweed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master_pb.Seaweed",
	HandlerType: (*SeaweedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupVolume",
			Handler:    _Seaweed_LookupVolume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendHeartbeat",
			Handler:       _Seaweed_SendHeartbeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "KeepConnected",
			Handler:       _Seaweed_KeepConnected_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "master.proto",
}

func init() { proto.RegisterFile("master.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6f, 0xdc, 0x44,
	0x14, 0xae, 0xbd, 0x9b, 0xec, 0xfa, 0x6d, 0x36, 0xdd, 0x4c, 0x22, 0xe4, 0x2e, 0xa5, 0x5d, 0xcc,
	0xc5, 0x08, 0x14, 0x95, 0x70, 0x44, 0x5c, 0xba, 0x0a, 0x22, 0x4a, 0x50, 0x83, 0x03, 0x3d, 0x70,
	0x31, 0xb3, 0xf6, 0x6b, 0x35, 0xca, 0x78, 0x6c, 0x66, 0x66, 0x93, 0x75, 0x2f, 0xfc, 0x67, 0x9c,
	0xf8, 0x6f, 0xb8, 0x71, 0xe3, 0x86, 0x66, 0x3c, 0xde, 0x1f, 0x6e, 0x00, 0x89, 0xdb, 0x9b, 0x6f,
	0xde, 0x8f, 0x6f, 0xde, 0xf7, 0x9e, 0x0d, 0x07, 0x05, 0x55, 0x1a, 0xe5, 0x69, 0x25, 0x4b, 0x5d,
	0x92, 0xa0, 0x39, 0xa5, 0xd5, 0x22, 0xfa, 0xc3, 0x87, 0xe0, 0x5b, 0xa4, 0x52, 0x2f, 0x90, 0x6a,
	0x72, 0x08, 0x3e, 0xab, 0x42, 0x6f, 0xe6, 0xc5, 0x41, 0xe2, 0xb3, 0x8a, 0x10, 0xe8, 0x57, 0xa5,
	0xd4, 0xa1, 0x3f, 0xf3, 0xe2, 0x71, 0x62, 0x6d, 0xf2, 0x11, 0x40, 0xb5, 0x5c, 0x70, 0x96, 0xa5,
	0x4b, 0xc9, 0xc3, 0x9e, 0xf5, 0x0d, 0x1a, 0xe4, 0x47, 0xc9, 0x49, 0x0c, 0x93, 0x82, 0xae, 0xd2,
	0xbb, 0x92, 0x2f, 0x0b, 0x4c, 0xb3, 0x72, 0x29, 0x74, 0xd8, 0xb7, 0xe1, 0x87, 0x05, 0x5d, 0xbd,
	0xb6, 0xf0, 0xdc, 0xa0, 0x64, 0x66, 0x58, 0xad, 0xd2, 0x37, 0x8c, 0x63, 0x7a, 0x8b, 0x75, 0xb8,
	0x37, 0xf3, 0xe2, 0x7e, 0x02, 0x05, 0x5d, 0x7d, 0xc3, 0x38, 0x5e, 0x62, 0x4d, 0x9e, 0xc3, 0x28,
	0xa7, 0x9a, 0xa6, 0x19, 0x0a, 0x8d, 0x32, 0xdc, 0xb7, 0xb5, 0xc0, 0x40, 0x73, 0x8b, 0x18, 0x7e,
	0x92, 0x66, 0xb7, 0xe1, 0xc0, 0xde, 0x58, 0xdb, 0xf0, 0xa3, 0x79, 0xc1, 0x44, 0x6a, 0x99, 0x0f,
	0x6d, 0xe9, 0xc0, 0x22, 0xd7, 0x86, 0xfe, 0xd7, 0x30, 0x68, 0xb8, 0xa9, 0x30, 0x98, 0xf5, 0xe2,
	0xd1, 0xd9, 0x27, 0xa7, 0xeb, 0x6e, 0x9c, 0x36, 0xf4, 0x2e, 0xc4, 0x9b, 0x52, 0x16, 0x54, 0xb3,
	0x52, 0x7c, 0x87, 0x4a, 0xd1, 0xb7, 0x98, 0xb4, 0x31, 0xe4, 0x09, 0x0c, 0x05, 0xde, 0xa7, 0x77,
	0x2c, 0x57, 0x21, 0xcc, 0x7a, 0xf1, 0x38, 0x19, 0x08, 0xbc, 0x7f, 0xcd, 0x72, 0x45, 0x3e, 0x86,
	0x83, 0x1c, 0x39, 0x6a, 0xcc, 0x9b, 0xeb, 0x91, 0xbd, 0x1e, 0x39, 0xcc, 0xb8, 0x44, 0x0a, 0x8e,
	0xd6, 0xcd, 0x4e, 0x50, 0x55, 0xa5, 0x50, 0x48, 0x62, 0x78, 0xdc, 0x64, 0xbf, 0x61, 0xef, 0xf0,
	0x8a, 0x15, 0x4c, 0x5b, 0x05, 0xfa, 0x49, 0x17, 0x26, 0x4f, 0x21, 0x50, 0x98, 0x49, 0xd4, 0x97,
	0x58, 0x5b, 0x4d, 0x82, 0x64, 0x03, 0x90, 0x0f, 0x60, 0x9f, 0x23, 0xcd, 0x51, 0x3a, 0x51, 0xdc,
	0x29, 0xfa, 0xdd, 0x87, 0xf0, 0x9f, 0x1e, 0x66, 0x15, 0xcf, 0x6d, 0xbd, 0x71, 0xe2, 0xb3, 0xdc,
	0x74, 0x54, 0xb1, 0x77, 0x68, 0xb3, 0xf7, 0x13, 0x6b, 0x93, 0x67, 0x00, 0x59, 0xc9, 0x39, 0x66,
	0x26, 0xd0, 0x25, 0xdf, 0x42, 0x4c, 0xc7, 0xad, 0x88, 0x1b, 0xb1, 0xfb, 0x49, 0x60, 0x90, 0x46,
	0xe7, 0x75, 0x5f, 0x9c, 0x43, 0xa3, 0xb3, 0xeb, 0x4b, 0xe3, 0xf2, 0x39, 0x90, 0xb6, 0x75, 0x8b,
	0x7a, 0xed, 0xb8, 0x6f, 0x1d, 0x27, 0xee, 0xe6, 0x65, 0xdd, 0x7a, 0x7f, 0x08, 0x81, 0x44, 0x9a,
	0xa7, 0xa5, 0xe0, 0xb5, 0x95, 0x7e, 0x98, 0x0c, 0x0d, 0xf0, 0x4a, 0xf0, 0x9a, 0x7c, 0x06, 0x47,
	0x12, 0x2b, 0xce, 0x32, 0x9a, 0x56, 0x9c, 0x66, 0x58, 0xa0, 0x68, 0xa7, 0x60, 0xe2, 0x2e, 0xae,
	0x5b, 0x9c, 0x84, 0x30, 0xb8, 0x43, 0xa9, 0xcc, 0xb3, 0x02, 0xeb, 0xd2, 0x1e, 0xc9, 0x04, 0x7a,
	0x5a, 0xf3, 0x10, 0x2c, 0x6a, 0xcc, 0x68, 0x00, 0x7b, 0xe7, 0x45, 0xa5, 0xeb, 0xe8, 0x37, 0x0f,
	0x1e, 0xdf, 0x2c, 0x2b, 0x94, 0x2f, 0x79, 0x99, 0xdd, 0x9e, 0xaf, 0xb4, 0xa4, 0xe4, 0x15, 0x1c,
	0xa2, 0xa4, 0x6a, 0x29, 0x0d, 0xf7, 0x9c, 0x89, 0xb7, 0xb6, 0xa5, 0xa3, 0xb3, 0x78, 0x6b, 0xb8,
	0x3a, 0x31, 0xa7, 0xe7, 0x4d, 0xc0, 0xdc, 0xfa, 0x27, 0x63, 0xdc, 0x3e, 0x4e, 0x7f, 0x82, 0xf1,
	0xce, 0xbd, 0x11, 0xc6, 0x0c, 0xbe, 0x93, 0xca, 0xda, 0x46, 0xf1, 0x8a, 0x4a, 0xa6, 0x6b, 0xb7,
	0xa0, 0xee, 0x64, 0x04, 0x71, 0xfb, 0x67, 0xe6, 0xb0, 0x67, 0xe7, 0x30, 0x68, 0x90, 0x8b, 0x5c,
	0x45, 0x9f, 0xc2, 0xf1, 0x9c, 0x33, 0x14, 0xfa, 0x8a, 0x29, 0x8d, 0x22, 0xc1, 0x5f, 0x96, 0xa8,
	0xb4, 0xa9, 0x20, 0x68, 0x81, 0x6e, 0xfd, 0xad, 0x1d, 0xfd, 0x0a, 0x87, 0xcd, 0xe8, 0x5c, 0x95,
	0x99, 0x9d, 0x1b, 0xd3, 0x18, 0xb3, 0xf7, 0x8d, 0x93, 0x31, 0x3b, 0x1f, 0x04, 0xbf, 0xfb, 0x41,
	0xd8, 0xde, 0x98, 0xde, 0xbf, 0x6f, 0x4c, 0xff, 0xfd, 0x8d, 0xf9, 0x01, 0x8e, 0xaf, 0xca, 0xf2,
	0x76, 0x59, 0x35, 0x34, 0x5a, 0xae, 0xbb, 0x2f, 0xf4, 0x66, 0x3d, 0x53, 0x73, 0xfd, 0xc2, 0xce,
	0xc4, 0xfa, 0xdd, 0x89, 0x8d, 0xfe, 0xf4, 0xe0, 0x64, 0x37, 0xad, 0xdb, 0xc5, 0x9f, 0xe1, 0x78,
	0x9d, 0x37, 0xe5, 0xee, 0xcd, 0x4d, 0x81, 0xd1, 0xd9, 0x8b, 0x2d, 0x31, 0x1f, 0x8a, 0x6e, 0x3f,
	0x1f, 0x79, 0xdb, 0xac, 0xe4, 0xe8, 0xae, 0x83, 0xa8, 0xe9, 0x0a, 0x26, 0x5d, 0x37, 0x33, 0xd0,
	0xeb, 0xaa, 0xae, 0xb3, 0xc3, 0x36, 0x92, 0x7c, 0x01, 0xc1, 0x86, 0x88, 0x6f, 0x89, 0x1c, 0xef,
	0x10, 0x71, 0xb5, 0x36, 0x5e, 0xe4, 0x04, 0xf6, 0x50, 0xca, 0xb2, 0xfd, 0x10, 0x34, 0x87, 0xe8,
	0x2b, 0x18, 0xfe, 0x6f, 0x15, 0xcf, 0xfe, 0xf2, 0x60, 0x70, 0x83, 0xf4, 0x1e, 0x31, 0x27, 0x17,
	0x30, 0xbe, 0x41, 0x91, 0x6f, 0x7e, 0x1b, 0x27, 0x5b, 0x7c, 0xd6, 0xe8, 0xf4, 0xe9, 0x43, 0x68,
	0xdb, 0xab, 0xe8, 0x51, 0xec, 0xbd, 0xf0, 0xc8, 0x35, 0x8c, 0x2f, 0x11, 0xab, 0x79, 0x29, 0x04,
	0x66, 0x1a, 0x73, 0xf2, 0x6c, 0x2b, 0xe8, 0x81, 0x21, 0x9d, 0x3e, 0x79, 0xef, 0x6b, 0xdd, 0xbe,
	0xc9, 0x65, 0xfc, 0x1e, 0x0e, 0xb6, 0xb5, 0xd9, 0x49, 0xf8, 0xc0, 0x24, 0x4d, 0x9f, 0xff, 0x87,
	0xa8, 0xd1, 0xa3, 0xc5, 0xbe, 0xfd, 0x6b, 0x7e, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0x9f, 0x87, 0x2e, 0x45, 0x07, 0x00, 0x00,
}
