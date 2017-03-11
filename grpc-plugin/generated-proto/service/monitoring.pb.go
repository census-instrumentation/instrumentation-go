// Code generated by protoc-gen-go.
// source: github.com/google/instrumentation-proto/service/monitoring.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/google/instrumentation-proto/service/monitoring.proto

It has these top-level messages:
	CanonicalRpcStats
	StatsRequest
	StatsResponse
	TraceRequest
	TraceResponse
	MonitoringDataGroup
	CustomMonitoringData
	ViewResponse
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_instrumentation "github.com/google/instrumentation-proto/stats"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

// Canonical RPC stats exported by gRPC.
type CanonicalRpcStats struct {
	RpcClientErrors            *CanonicalRpcStats_View `protobuf:"bytes,1,opt,name=rpc_client_errors,json=rpcClientErrors" json:"rpc_client_errors,omitempty"`
	RpcClientCompletedRpcs     *CanonicalRpcStats_View `protobuf:"bytes,2,opt,name=rpc_client_completed_rpcs,json=rpcClientCompletedRpcs" json:"rpc_client_completed_rpcs,omitempty"`
	RpcClientStartedRpcs       *CanonicalRpcStats_View `protobuf:"bytes,3,opt,name=rpc_client_started_rpcs,json=rpcClientStartedRpcs" json:"rpc_client_started_rpcs,omitempty"`
	RpcClientElapsedTime       *CanonicalRpcStats_View `protobuf:"bytes,4,opt,name=rpc_client_elapsed_time,json=rpcClientElapsedTime" json:"rpc_client_elapsed_time,omitempty"`
	RpcClientServerElapsedTime *CanonicalRpcStats_View `protobuf:"bytes,5,opt,name=rpc_client_server_elapsed_time,json=rpcClientServerElapsedTime" json:"rpc_client_server_elapsed_time,omitempty"`
	RpcClientRequestBytes      *CanonicalRpcStats_View `protobuf:"bytes,6,opt,name=rpc_client_request_bytes,json=rpcClientRequestBytes" json:"rpc_client_request_bytes,omitempty"`
	RpcClientResponseBytes     *CanonicalRpcStats_View `protobuf:"bytes,7,opt,name=rpc_client_response_bytes,json=rpcClientResponseBytes" json:"rpc_client_response_bytes,omitempty"`
	RpcClientRequestCount      *CanonicalRpcStats_View `protobuf:"bytes,8,opt,name=rpc_client_request_count,json=rpcClientRequestCount" json:"rpc_client_request_count,omitempty"`
	RpcClientResponseCount     *CanonicalRpcStats_View `protobuf:"bytes,9,opt,name=rpc_client_response_count,json=rpcClientResponseCount" json:"rpc_client_response_count,omitempty"`
	RpcServerErrors            *CanonicalRpcStats_View `protobuf:"bytes,10,opt,name=rpc_server_errors,json=rpcServerErrors" json:"rpc_server_errors,omitempty"`
	RpcServerCompletedRpcs     *CanonicalRpcStats_View `protobuf:"bytes,11,opt,name=rpc_server_completed_rpcs,json=rpcServerCompletedRpcs" json:"rpc_server_completed_rpcs,omitempty"`
	RpcServerServerElapsedTime *CanonicalRpcStats_View `protobuf:"bytes,12,opt,name=rpc_server_server_elapsed_time,json=rpcServerServerElapsedTime" json:"rpc_server_server_elapsed_time,omitempty"`
	RpcServerRequestBytes      *CanonicalRpcStats_View `protobuf:"bytes,13,opt,name=rpc_server_request_bytes,json=rpcServerRequestBytes" json:"rpc_server_request_bytes,omitempty"`
	RpcServerResponseBytes     *CanonicalRpcStats_View `protobuf:"bytes,14,opt,name=rpc_server_response_bytes,json=rpcServerResponseBytes" json:"rpc_server_response_bytes,omitempty"`
	RpcServerRequestCount      *CanonicalRpcStats_View `protobuf:"bytes,15,opt,name=rpc_server_request_count,json=rpcServerRequestCount" json:"rpc_server_request_count,omitempty"`
	RpcServerResponseCount     *CanonicalRpcStats_View `protobuf:"bytes,16,opt,name=rpc_server_response_count,json=rpcServerResponseCount" json:"rpc_server_response_count,omitempty"`
	RpcServerElapsedTime       *CanonicalRpcStats_View `protobuf:"bytes,17,opt,name=rpc_server_elapsed_time,json=rpcServerElapsedTime" json:"rpc_server_elapsed_time,omitempty"`
}

func (m *CanonicalRpcStats) Reset()                    { *m = CanonicalRpcStats{} }
func (m *CanonicalRpcStats) String() string            { return proto.CompactTextString(m) }
func (*CanonicalRpcStats) ProtoMessage()               {}
func (*CanonicalRpcStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CanonicalRpcStats) GetRpcClientErrors() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientErrors
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientCompletedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientCompletedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientStartedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientStartedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientServerElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientRequestBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientRequestBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientResponseBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientResponseBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientRequestCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientRequestCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcClientResponseCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcClientResponseCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerErrors() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerErrors
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerCompletedRpcs() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerCompletedRpcs
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerServerElapsedTime
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerRequestBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerRequestBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerResponseBytes() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerResponseBytes
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerRequestCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerRequestCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerResponseCount() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerResponseCount
	}
	return nil
}

func (m *CanonicalRpcStats) GetRpcServerElapsedTime() *CanonicalRpcStats_View {
	if m != nil {
		return m.RpcServerElapsedTime
	}
	return nil
}

// Wrapper combining View and ViewDescriptor.
type CanonicalRpcStats_View struct {
	MeasurementDescriptor *google_instrumentation.MeasurementDescriptor `protobuf:"bytes,1,opt,name=measurement_descriptor,json=measurementDescriptor" json:"measurement_descriptor,omitempty"`
	ViewDescriptor        *google_instrumentation.ViewDescriptor        `protobuf:"bytes,2,opt,name=view_descriptor,json=viewDescriptor" json:"view_descriptor,omitempty"`
	View                  *google_instrumentation.View                  `protobuf:"bytes,3,opt,name=view" json:"view,omitempty"`
}

func (m *CanonicalRpcStats_View) Reset()                    { *m = CanonicalRpcStats_View{} }
func (m *CanonicalRpcStats_View) String() string            { return proto.CompactTextString(m) }
func (*CanonicalRpcStats_View) ProtoMessage()               {}
func (*CanonicalRpcStats_View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CanonicalRpcStats_View) GetMeasurementDescriptor() *google_instrumentation.MeasurementDescriptor {
	if m != nil {
		return m.MeasurementDescriptor
	}
	return nil
}

func (m *CanonicalRpcStats_View) GetViewDescriptor() *google_instrumentation.ViewDescriptor {
	if m != nil {
		return m.ViewDescriptor
	}
	return nil
}

func (m *CanonicalRpcStats_View) GetView() *google_instrumentation.View {
	if m != nil {
		return m.View
	}
	return nil
}

// This message is sent when requesting a set of stats (Census Views) from
// a client system, as part of the MonitoringService API's.
// TODO(aveitch): should this be named ViewRequest?
type StatsRequest struct {
	// An optional set of ViewDescriptor names. Only Views using these
	// descriptors will be sent back in the response. If no names are provided,
	// then all Views present in the client system will be included in every
	// response. If measurement_names is also provided, then Views matching the
	// intersection of the two are returned.
	// TODO(aveitch): Consider making this a list of regexes or prefix matches in
	// the future.
	ViewNames []string `protobuf:"bytes,1,rep,name=view_names,json=viewNames" json:"view_names,omitempty"`
	// An optional set of MeasurementDescriptor names. Only Views using these
	// descriptors will be sent back in the response. If no names are provided,
	// then all Views present in the client system will be included in every
	// response. If view_names is also provided, then Views matching the
	// intersection of the two are returned.
	// TODO(aveitch): Consider making this a list of regexes or prefix matches in
	// the future.
	MeasurementNames []string `protobuf:"bytes,2,rep,name=measurement_names,json=measurementNames" json:"measurement_names,omitempty"`
	// By default, the MeasurementDescriptors and ViewDescriptors corresponding to
	// the Views that can returned in a StatsResponse will be included in the
	// first such response. Set this to true to have them not sent.
	DontIncludeDescriptorsInFirstResponse bool `protobuf:"varint,3,opt,name=dont_include_descriptors_in_first_response,json=dontIncludeDescriptorsInFirstResponse" json:"dont_include_descriptors_in_first_response,omitempty"`
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatsRequest) GetViewNames() []string {
	if m != nil {
		return m.ViewNames
	}
	return nil
}

func (m *StatsRequest) GetMeasurementNames() []string {
	if m != nil {
		return m.MeasurementNames
	}
	return nil
}

func (m *StatsRequest) GetDontIncludeDescriptorsInFirstResponse() bool {
	if m != nil {
		return m.DontIncludeDescriptorsInFirstResponse
	}
	return false
}

// This message is sent in response to a GetStats or WatchStats call.
type StatsResponse struct {
	// The set of Views corresponding to the StatsRequest.
	ViewResponses []*ViewResponse `protobuf:"bytes,1,rep,name=view_responses,json=viewResponses" json:"view_responses,omitempty"`
}

func (m *StatsResponse) Reset()                    { *m = StatsResponse{} }
func (m *StatsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()               {}
func (*StatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatsResponse) GetViewResponses() []*ViewResponse {
	if m != nil {
		return m.ViewResponses
	}
	return nil
}

type TraceRequest struct {
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type TraceResponse struct {
}

func (m *TraceResponse) Reset()                    { *m = TraceResponse{} }
func (m *TraceResponse) String() string            { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()               {}
func (*TraceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type MonitoringDataGroup struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *MonitoringDataGroup) Reset()                    { *m = MonitoringDataGroup{} }
func (m *MonitoringDataGroup) String() string            { return proto.CompactTextString(m) }
func (*MonitoringDataGroup) ProtoMessage()               {}
func (*MonitoringDataGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MonitoringDataGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The wrapper for custom monitoring data.
type CustomMonitoringData struct {
	// can be any application specific monitoring data
	Contents *google_protobuf.Any `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
}

func (m *CustomMonitoringData) Reset()                    { *m = CustomMonitoringData{} }
func (m *CustomMonitoringData) String() string            { return proto.CompactTextString(m) }
func (*CustomMonitoringData) ProtoMessage()               {}
func (*CustomMonitoringData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CustomMonitoringData) GetContents() *google_protobuf.Any {
	if m != nil {
		return m.Contents
	}
	return nil
}

// This message contains all information relevant to a single View. It is used
// in StatsResponse.
type ViewResponse struct {
	// Each ViewResponse can optionally contain the MeasurementDescriptor and
	// ViewDescriptor for the View. These will be sent in the first WatchStats
	// response, or all GetStats responses. This is disabled if
	// dont_include_descriptors_in_first_response is set to true in the
	// StatsRequest.
	MeasurementDescriptor *google_instrumentation.MeasurementDescriptor `protobuf:"bytes,1,opt,name=measurement_descriptor,json=measurementDescriptor" json:"measurement_descriptor,omitempty"`
	ViewDescriptor        *google_instrumentation.ViewDescriptor        `protobuf:"bytes,2,opt,name=view_descriptor,json=viewDescriptor" json:"view_descriptor,omitempty"`
	// The View data.
	View *google_instrumentation.View `protobuf:"bytes,3,opt,name=view" json:"view,omitempty"`
}

func (m *ViewResponse) Reset()                    { *m = ViewResponse{} }
func (m *ViewResponse) String() string            { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()               {}
func (*ViewResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ViewResponse) GetMeasurementDescriptor() *google_instrumentation.MeasurementDescriptor {
	if m != nil {
		return m.MeasurementDescriptor
	}
	return nil
}

func (m *ViewResponse) GetViewDescriptor() *google_instrumentation.ViewDescriptor {
	if m != nil {
		return m.ViewDescriptor
	}
	return nil
}

func (m *ViewResponse) GetView() *google_instrumentation.View {
	if m != nil {
		return m.View
	}
	return nil
}

func init() {
	proto.RegisterType((*CanonicalRpcStats)(nil), "google.instrumentation.CanonicalRpcStats")
	proto.RegisterType((*CanonicalRpcStats_View)(nil), "google.instrumentation.CanonicalRpcStats.View")
	proto.RegisterType((*StatsRequest)(nil), "google.instrumentation.StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "google.instrumentation.StatsResponse")
	proto.RegisterType((*TraceRequest)(nil), "google.instrumentation.TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "google.instrumentation.TraceResponse")
	proto.RegisterType((*MonitoringDataGroup)(nil), "google.instrumentation.MonitoringDataGroup")
	proto.RegisterType((*CustomMonitoringData)(nil), "google.instrumentation.CustomMonitoringData")
	proto.RegisterType((*ViewResponse)(nil), "google.instrumentation.ViewResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Monitoring service

type MonitoringClient interface {
	// Return canonical RPC stats
	GetCanonicalRpcStats(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CanonicalRpcStats, error)
	// Query the server for specific stats
	GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	// Request the server to stream back snapshots of the requested stats
	WatchStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (Monitoring_WatchStatsClient, error)
	// Return request traces.
	GetRequestTraces(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error)
	// Return application-defined groups of monitoring data.
	// This is a low level facility to allow extension of the monitoring API to
	// application-specific monitoring data. Frameworks may use this to define
	// additional groups of monitoring data made available by servers.
	GetCustomMonitoringData(ctx context.Context, in *MonitoringDataGroup, opts ...grpc.CallOption) (*CustomMonitoringData, error)
}

type monitoringClient struct {
	cc *grpc.ClientConn
}

func NewMonitoringClient(cc *grpc.ClientConn) MonitoringClient {
	return &monitoringClient{cc}
}

func (c *monitoringClient) GetCanonicalRpcStats(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CanonicalRpcStats, error) {
	out := new(CanonicalRpcStats)
	err := grpc.Invoke(ctx, "/google.instrumentation.Monitoring/GetCanonicalRpcStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := grpc.Invoke(ctx, "/google.instrumentation.Monitoring/GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) WatchStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (Monitoring_WatchStatsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Monitoring_serviceDesc.Streams[0], c.cc, "/google.instrumentation.Monitoring/WatchStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringWatchStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Monitoring_WatchStatsClient interface {
	Recv() (*StatsResponse, error)
	grpc.ClientStream
}

type monitoringWatchStatsClient struct {
	grpc.ClientStream
}

func (x *monitoringWatchStatsClient) Recv() (*StatsResponse, error) {
	m := new(StatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitoringClient) GetRequestTraces(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceResponse, error) {
	out := new(TraceResponse)
	err := grpc.Invoke(ctx, "/google.instrumentation.Monitoring/GetRequestTraces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetCustomMonitoringData(ctx context.Context, in *MonitoringDataGroup, opts ...grpc.CallOption) (*CustomMonitoringData, error) {
	out := new(CustomMonitoringData)
	err := grpc.Invoke(ctx, "/google.instrumentation.Monitoring/GetCustomMonitoringData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitoring service

type MonitoringServer interface {
	// Return canonical RPC stats
	GetCanonicalRpcStats(context.Context, *google_protobuf1.Empty) (*CanonicalRpcStats, error)
	// Query the server for specific stats
	GetStats(context.Context, *StatsRequest) (*StatsResponse, error)
	// Request the server to stream back snapshots of the requested stats
	WatchStats(*StatsRequest, Monitoring_WatchStatsServer) error
	// Return request traces.
	GetRequestTraces(context.Context, *TraceRequest) (*TraceResponse, error)
	// Return application-defined groups of monitoring data.
	// This is a low level facility to allow extension of the monitoring API to
	// application-specific monitoring data. Frameworks may use this to define
	// additional groups of monitoring data made available by servers.
	GetCustomMonitoringData(context.Context, *MonitoringDataGroup) (*CustomMonitoringData, error)
}

func RegisterMonitoringServer(s *grpc.Server, srv MonitoringServer) {
	s.RegisterService(&_Monitoring_serviceDesc, srv)
}

func _Monitoring_GetCanonicalRpcStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetCanonicalRpcStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.instrumentation.Monitoring/GetCanonicalRpcStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetCanonicalRpcStats(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.instrumentation.Monitoring/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_WatchStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitoringServer).WatchStats(m, &monitoringWatchStatsServer{stream})
}

type Monitoring_WatchStatsServer interface {
	Send(*StatsResponse) error
	grpc.ServerStream
}

type monitoringWatchStatsServer struct {
	grpc.ServerStream
}

func (x *monitoringWatchStatsServer) Send(m *StatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Monitoring_GetRequestTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetRequestTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.instrumentation.Monitoring/GetRequestTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetRequestTraces(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetCustomMonitoringData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitoringDataGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetCustomMonitoringData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.instrumentation.Monitoring/GetCustomMonitoringData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetCustomMonitoringData(ctx, req.(*MonitoringDataGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.instrumentation.Monitoring",
	HandlerType: (*MonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCanonicalRpcStats",
			Handler:    _Monitoring_GetCanonicalRpcStats_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Monitoring_GetStats_Handler,
		},
		{
			MethodName: "GetRequestTraces",
			Handler:    _Monitoring_GetRequestTraces_Handler,
		},
		{
			MethodName: "GetCustomMonitoringData",
			Handler:    _Monitoring_GetCustomMonitoringData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchStats",
			Handler:       _Monitoring_WatchStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/google/instrumentation-proto/service/monitoring.proto",
}

func init() {
	proto.RegisterFile("github.com/google/instrumentation-proto/service/monitoring.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 878 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x57, 0xdd, 0x6f, 0x1b, 0x45,
	0x10, 0x8f, 0x9b, 0xd0, 0x26, 0x93, 0xef, 0x25, 0x4d, 0x5d, 0xf3, 0x21, 0x74, 0xa2, 0x55, 0x4b,
	0xe9, 0x39, 0x0a, 0x6f, 0x3c, 0x20, 0x88, 0x1b, 0x42, 0x85, 0x0a, 0xe8, 0x5a, 0x81, 0x5a, 0x40,
	0xa7, 0xcb, 0x7a, 0xea, 0xae, 0xe4, 0xdb, 0x3d, 0x76, 0xf7, 0x5c, 0xe5, 0x8d, 0x7f, 0x8a, 0xbf,
	0x8b, 0x47, 0xde, 0x10, 0xda, 0x0f, 0x9f, 0xef, 0xec, 0xbb, 0x10, 0x5d, 0xcc, 0x13, 0x6f, 0xde,
	0x9d, 0x99, 0xdf, 0x6f, 0x66, 0xfc, 0x9b, 0xdd, 0x3d, 0xf8, 0x72, 0xc4, 0xf4, 0x9b, 0xfc, 0x3c,
	0xa4, 0x22, 0xed, 0x8f, 0x84, 0x18, 0x8d, 0xb1, 0xcf, 0xb8, 0xd2, 0x32, 0x4f, 0x91, 0xeb, 0x44,
	0x33, 0xc1, 0x1f, 0x67, 0x52, 0x68, 0xd1, 0x57, 0x28, 0x27, 0x8c, 0x62, 0x3f, 0x15, 0x9c, 0x69,
	0x21, 0x19, 0x1f, 0x85, 0xd6, 0x40, 0x0e, 0x5d, 0x58, 0x38, 0x17, 0xd6, 0xfb, 0xfc, 0xca, 0xc8,
	0x3a, 0xd1, 0xaa, 0x4f, 0x91, 0xab, 0x5c, 0x39, 0xcc, 0xde, 0x5d, 0x1f, 0x60, 0x57, 0xe7, 0xf9,
	0xeb, 0x7e, 0xc2, 0x2f, 0xbc, 0xe9, 0xbd, 0x79, 0x13, 0xa6, 0x99, 0xf6, 0xc6, 0xe0, 0xf7, 0x3d,
	0xd8, 0x1f, 0x24, 0x5c, 0x70, 0x46, 0x93, 0x71, 0x94, 0xd1, 0xe7, 0x06, 0x9a, 0xbc, 0x82, 0x7d,
	0x99, 0xd1, 0x98, 0x8e, 0x19, 0x72, 0x1d, 0xa3, 0x94, 0x42, 0xaa, 0x6e, 0xe7, 0xa3, 0xce, 0x83,
	0xcd, 0xe3, 0x30, 0xac, 0xcf, 0x3e, 0x5c, 0x40, 0x09, 0x7f, 0x64, 0xf8, 0x36, 0xda, 0x95, 0x19,
	0x1d, 0x58, 0x9c, 0x53, 0x0b, 0x43, 0x18, 0xdc, 0x2d, 0x61, 0x53, 0x91, 0x66, 0x63, 0xd4, 0x38,
	0x8c, 0x65, 0x46, 0x55, 0xf7, 0x46, 0x2b, 0x8e, 0xc3, 0x82, 0x63, 0x30, 0x85, 0x8b, 0x32, 0xaa,
	0x08, 0xc2, 0x9d, 0x12, 0x95, 0xd2, 0x89, 0x2c, 0x88, 0x56, 0x5b, 0x11, 0x1d, 0x14, 0x44, 0xcf,
	0x1d, 0x58, 0x0d, 0x0d, 0x8e, 0x93, 0x4c, 0xe1, 0x30, 0xd6, 0x2c, 0xc5, 0xee, 0xda, 0x35, 0x69,
	0x4e, 0x1d, 0xd8, 0x0b, 0x96, 0x22, 0x91, 0xf0, 0x61, 0xb9, 0x1a, 0x94, 0x13, 0x94, 0x55, 0xb6,
	0x77, 0x5a, 0xb1, 0xf5, 0x66, 0x45, 0x59, 0xcc, 0x32, 0xe7, 0x08, 0xba, 0x25, 0x4e, 0x89, 0xbf,
	0xe5, 0xa8, 0x74, 0x7c, 0x7e, 0xa1, 0x51, 0x75, 0x6f, 0xb6, 0x62, 0xbb, 0x5d, 0xb0, 0x45, 0x0e,
	0xed, 0xc4, 0x80, 0xcd, 0xa9, 0x42, 0xa2, 0xca, 0x04, 0x57, 0xe8, 0x99, 0x6e, 0x5d, 0x53, 0x15,
	0x91, 0x87, 0x73, 0x54, 0xf5, 0x35, 0x51, 0x91, 0x73, 0xdd, 0x5d, 0x5f, 0x4e, 0x4d, 0x03, 0x03,
	0xd6, 0x54, 0x93, 0x63, 0xda, 0x58, 0x52, 0x4d, 0x8e, 0xca, 0x0f, 0xec, 0x54, 0x14, 0x6e, 0x60,
	0xa1, 0xf5, 0xc0, 0x7a, 0x21, 0x54, 0x06, 0xd6, 0x63, 0xcf, 0x0d, 0xec, 0x66, 0xeb, 0x32, 0x1c,
	0x47, 0x75, 0x60, 0xbd, 0xc4, 0x3d, 0x55, 0x9d, 0xc4, 0xb7, 0x5a, 0x4b, 0xdc, 0xf1, 0x35, 0x4a,
	0xdc, 0x93, 0x55, 0x25, 0xbe, 0xdd, 0x5a, 0x0e, 0x8e, 0xa7, 0x4e, 0xe2, 0x05, 0x51, 0x45, 0xe2,
	0x3b, 0xd7, 0xec, 0x63, 0xad, 0xc4, 0xe7, 0x6a, 0x72, 0xc2, 0xdb, 0x5d, 0x4e, 0x4d, 0x15, 0x89,
	0xcf, 0xd7, 0xe4, 0x98, 0xf6, 0x96, 0x54, 0x93, 0xa3, 0xf2, 0xa7, 0x6c, 0x9d, 0x28, 0xf6, 0x5b,
	0x9f, 0xb2, 0x0b, 0x72, 0xe8, 0xfd, 0xd9, 0x81, 0x35, 0x63, 0x26, 0x43, 0x38, 0x4c, 0x31, 0x51,
	0xb9, 0x44, 0x83, 0x14, 0x0f, 0x51, 0x51, 0xc9, 0x32, 0x2d, 0xa4, 0xbf, 0x08, 0x1f, 0x37, 0xd1,
	0x3d, 0x9b, 0x45, 0x3d, 0x29, 0x82, 0xa2, 0xdb, 0x69, 0xdd, 0x36, 0xf9, 0x1e, 0x76, 0x27, 0x0c,
	0xdf, 0x96, 0xe1, 0xdd, 0x1d, 0x78, 0xbf, 0x09, 0xde, 0x24, 0x57, 0xc2, 0xdd, 0x99, 0x54, 0xd6,
	0xe4, 0x08, 0xd6, 0xcc, 0x8e, 0xbf, 0xe0, 0xde, 0xbf, 0x0c, 0x25, 0xb2, 0x9e, 0xc1, 0x1f, 0x1d,
	0xd8, 0xb2, 0x6d, 0xf1, 0xff, 0x2c, 0xf9, 0x00, 0xc0, 0xe6, 0xc4, 0x93, 0x14, 0xcd, 0xb5, 0xbf,
	0xfa, 0x60, 0x23, 0xda, 0x30, 0x3b, 0xdf, 0x99, 0x0d, 0xf2, 0x08, 0xf6, 0xcb, 0x8d, 0x71, 0x5e,
	0x37, 0xac, 0xd7, 0x5e, 0xc9, 0xe0, 0x9c, 0x5f, 0xc2, 0x27, 0x43, 0xc1, 0x75, 0xcc, 0x38, 0x1d,
	0xe7, 0x43, 0x2c, 0xd5, 0xa9, 0x62, 0xc6, 0xe3, 0xd7, 0x4c, 0xaa, 0xd9, 0xd1, 0x68, 0x93, 0x5e,
	0x8f, 0xee, 0x99, 0x88, 0xa7, 0x2e, 0x60, 0x56, 0x99, 0x7a, 0xca, 0xbf, 0x36, 0xde, 0x53, 0x55,
	0x04, 0xbf, 0xc0, 0xb6, 0x4f, 0xdb, 0x6d, 0x90, 0x6f, 0xc1, 0x36, 0xa3, 0x80, 0x73, 0xb9, 0x6f,
	0x1e, 0x7f, 0x7c, 0x69, 0x13, 0xbc, 0x73, 0xb4, 0x3d, 0x29, 0xad, 0x54, 0xb0, 0x03, 0x5b, 0x2f,
	0x64, 0x42, 0xd1, 0x37, 0x25, 0xd8, 0x85, 0x6d, 0xbf, 0xf6, 0xf4, 0x0f, 0xe1, 0xdd, 0x67, 0xc5,
	0xcb, 0xee, 0x49, 0xa2, 0x93, 0x33, 0x29, 0xf2, 0x8c, 0x10, 0x58, 0x33, 0x1d, 0xb1, 0x22, 0xd9,
	0x88, 0xec, 0xef, 0xe0, 0x1b, 0x38, 0x18, 0xe4, 0x4a, 0x8b, 0xb4, 0x1a, 0x40, 0x8e, 0x60, 0x9d,
	0x0a, 0xae, 0x91, 0xeb, 0xe9, 0xeb, 0xea, 0x60, 0x9a, 0xea, 0xf4, 0xb1, 0x16, 0x7e, 0xc5, 0x2f,
	0xa2, 0xc2, 0x2b, 0xf8, 0xab, 0x03, 0x5b, 0xe5, 0xac, 0xff, 0x37, 0x2a, 0x3d, 0xfe, 0x7b, 0x15,
	0x60, 0xd6, 0x3e, 0xf2, 0x33, 0x1c, 0x9c, 0xa1, 0x5e, 0x7c, 0xb9, 0x1e, 0x2e, 0x34, 0xf0, 0xd4,
	0xbc, 0x76, 0x7b, 0x0f, 0xaf, 0x7c, 0x38, 0x04, 0x2b, 0xe4, 0x25, 0xac, 0x9f, 0xa1, 0x76, 0x80,
	0x8d, 0xe2, 0x29, 0x8f, 0x4c, 0xef, 0xde, 0xbf, 0x78, 0x79, 0xcd, 0xac, 0x90, 0x5f, 0x01, 0x7e,
	0x4a, 0x34, 0x7d, 0xf3, 0x5f, 0x80, 0x1f, 0x75, 0x48, 0x02, 0x7b, 0x67, 0x38, 0x7d, 0x85, 0x58,
	0xbd, 0x5e, 0x42, 0x52, 0xd6, 0x77, 0x33, 0x49, 0x55, 0xf5, 0x2b, 0x64, 0x02, 0x77, 0x4c, 0xe7,
	0xeb, 0xf4, 0xfc, 0xa8, 0x51, 0x6c, 0x8b, 0x83, 0xd2, 0xfb, 0xb4, 0xf1, 0x1f, 0xa9, 0x81, 0x0e,
	0x56, 0x4e, 0xbe, 0x80, 0xfb, 0x54, 0xa4, 0x4d, 0x41, 0xf6, 0x93, 0xc8, 0xfd, 0xe9, 0x27, 0x9b,
	0x03, 0xfb, 0x65, 0xf4, 0x83, 0x59, 0xbc, 0xba, 0xe5, 0x3f, 0xc3, 0xce, 0x6f, 0x5a, 0xe3, 0x67,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x57, 0x99, 0x69, 0xc0, 0x0d, 0x00, 0x00,
}
