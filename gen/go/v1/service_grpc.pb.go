// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/service.proto

package v1

import (
	context "context"
	types "github.com/garethgeorge/backrest/gen/go/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Backrest_GetFeatures_FullMethodName        = "/v1.Backrest/GetFeatures"
	Backrest_GetConfig_FullMethodName          = "/v1.Backrest/GetConfig"
	Backrest_SetConfig_FullMethodName          = "/v1.Backrest/SetConfig"
	Backrest_AddRepo_FullMethodName            = "/v1.Backrest/AddRepo"
	Backrest_GetOperationEvents_FullMethodName = "/v1.Backrest/GetOperationEvents"
	Backrest_GetOperations_FullMethodName      = "/v1.Backrest/GetOperations"
	Backrest_ListSnapshots_FullMethodName      = "/v1.Backrest/ListSnapshots"
	Backrest_ListSnapshotFiles_FullMethodName  = "/v1.Backrest/ListSnapshotFiles"
	Backrest_Backup_FullMethodName             = "/v1.Backrest/Backup"
	Backrest_DoRepoTask_FullMethodName         = "/v1.Backrest/DoRepoTask"
	Backrest_Forget_FullMethodName             = "/v1.Backrest/Forget"
	Backrest_Restore_FullMethodName            = "/v1.Backrest/Restore"
	Backrest_Cancel_FullMethodName             = "/v1.Backrest/Cancel"
	Backrest_GetLogs_FullMethodName            = "/v1.Backrest/GetLogs"
	Backrest_RunCommand_FullMethodName         = "/v1.Backrest/RunCommand"
	Backrest_GetDownloadURL_FullMethodName     = "/v1.Backrest/GetDownloadURL"
	Backrest_ClearHistory_FullMethodName       = "/v1.Backrest/ClearHistory"
	Backrest_PathAutocomplete_FullMethodName   = "/v1.Backrest/PathAutocomplete"
)

// BackrestClient is the client API for Backrest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackrestClient interface {
	// GetFeatures returns a list of features supported by the server.
	GetFeatures(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FeatureList, error)
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Config, error)
	SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error)
	AddRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Config, error)
	GetOperationEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Backrest_GetOperationEventsClient, error)
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*OperationList, error)
	ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ResticSnapshotList, error)
	ListSnapshotFiles(ctx context.Context, in *ListSnapshotFilesRequest, opts ...grpc.CallOption) (*ListSnapshotFilesResponse, error)
	// Backup schedules a backup operation. It accepts a plan id and returns empty if the task is enqueued.
	Backup(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DoRepoTask schedules a repo task. It accepts a repo id and a task type and returns empty if the task is enqueued.
	DoRepoTask(ctx context.Context, in *DoRepoTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Forget schedules a forget operation. It accepts a plan id and returns empty if the task is enqueued.
	Forget(ctx context.Context, in *ForgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Restore schedules a restore operation.
	Restore(ctx context.Context, in *RestoreSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Cancel attempts to cancel a task with the given operation ID. Not guaranteed to succeed.
	Cancel(ctx context.Context, in *types.Int64Value, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetLogs returns the keyed large data for the given operation.
	GetLogs(ctx context.Context, in *LogDataRequest, opts ...grpc.CallOption) (Backrest_GetLogsClient, error)
	// RunCommand executes a generic restic command on the repository.
	RunCommand(ctx context.Context, in *RunCommandRequest, opts ...grpc.CallOption) (Backrest_RunCommandClient, error)
	// GetDownloadURL returns a signed download URL given a forget operation ID.
	GetDownloadURL(ctx context.Context, in *types.Int64Value, opts ...grpc.CallOption) (*types.StringValue, error)
	// Clears the history of operations
	ClearHistory(ctx context.Context, in *ClearHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// PathAutocomplete provides path autocompletion options for a given filesystem path.
	PathAutocomplete(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringList, error)
}

type backrestClient struct {
	cc grpc.ClientConnInterface
}

func NewBackrestClient(cc grpc.ClientConnInterface) BackrestClient {
	return &backrestClient{cc}
}

func (c *backrestClient) GetFeatures(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FeatureList, error) {
	out := new(FeatureList)
	err := c.cc.Invoke(ctx, Backrest_GetFeatures_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, Backrest_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) SetConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, Backrest_SetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) AddRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, Backrest_AddRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) GetOperationEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Backrest_GetOperationEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Backrest_ServiceDesc.Streams[0], Backrest_GetOperationEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backrestGetOperationEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backrest_GetOperationEventsClient interface {
	Recv() (*OperationEvent, error)
	grpc.ClientStream
}

type backrestGetOperationEventsClient struct {
	grpc.ClientStream
}

func (x *backrestGetOperationEventsClient) Recv() (*OperationEvent, error) {
	m := new(OperationEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backrestClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*OperationList, error) {
	out := new(OperationList)
	err := c.cc.Invoke(ctx, Backrest_GetOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ResticSnapshotList, error) {
	out := new(ResticSnapshotList)
	err := c.cc.Invoke(ctx, Backrest_ListSnapshots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) ListSnapshotFiles(ctx context.Context, in *ListSnapshotFilesRequest, opts ...grpc.CallOption) (*ListSnapshotFilesResponse, error) {
	out := new(ListSnapshotFilesResponse)
	err := c.cc.Invoke(ctx, Backrest_ListSnapshotFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) Backup(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_Backup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) DoRepoTask(ctx context.Context, in *DoRepoTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_DoRepoTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) Forget(ctx context.Context, in *ForgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_Forget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) Restore(ctx context.Context, in *RestoreSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) Cancel(ctx context.Context, in *types.Int64Value, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_Cancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) GetLogs(ctx context.Context, in *LogDataRequest, opts ...grpc.CallOption) (Backrest_GetLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Backrest_ServiceDesc.Streams[1], Backrest_GetLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backrestGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backrest_GetLogsClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type backrestGetLogsClient struct {
	grpc.ClientStream
}

func (x *backrestGetLogsClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backrestClient) RunCommand(ctx context.Context, in *RunCommandRequest, opts ...grpc.CallOption) (Backrest_RunCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &Backrest_ServiceDesc.Streams[2], Backrest_RunCommand_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backrestRunCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backrest_RunCommandClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type backrestRunCommandClient struct {
	grpc.ClientStream
}

func (x *backrestRunCommandClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backrestClient) GetDownloadURL(ctx context.Context, in *types.Int64Value, opts ...grpc.CallOption) (*types.StringValue, error) {
	out := new(types.StringValue)
	err := c.cc.Invoke(ctx, Backrest_GetDownloadURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) ClearHistory(ctx context.Context, in *ClearHistoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Backrest_ClearHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backrestClient) PathAutocomplete(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringList, error) {
	out := new(types.StringList)
	err := c.cc.Invoke(ctx, Backrest_PathAutocomplete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackrestServer is the server API for Backrest service.
// All implementations must embed UnimplementedBackrestServer
// for forward compatibility
type BackrestServer interface {
	// GetFeatures returns a list of features supported by the server.
	GetFeatures(context.Context, *emptypb.Empty) (*FeatureList, error)
	GetConfig(context.Context, *emptypb.Empty) (*Config, error)
	SetConfig(context.Context, *Config) (*Config, error)
	AddRepo(context.Context, *Repo) (*Config, error)
	GetOperationEvents(*emptypb.Empty, Backrest_GetOperationEventsServer) error
	GetOperations(context.Context, *GetOperationsRequest) (*OperationList, error)
	ListSnapshots(context.Context, *ListSnapshotsRequest) (*ResticSnapshotList, error)
	ListSnapshotFiles(context.Context, *ListSnapshotFilesRequest) (*ListSnapshotFilesResponse, error)
	// Backup schedules a backup operation. It accepts a plan id and returns empty if the task is enqueued.
	Backup(context.Context, *types.StringValue) (*emptypb.Empty, error)
	// DoRepoTask schedules a repo task. It accepts a repo id and a task type and returns empty if the task is enqueued.
	DoRepoTask(context.Context, *DoRepoTaskRequest) (*emptypb.Empty, error)
	// Forget schedules a forget operation. It accepts a plan id and returns empty if the task is enqueued.
	Forget(context.Context, *ForgetRequest) (*emptypb.Empty, error)
	// Restore schedules a restore operation.
	Restore(context.Context, *RestoreSnapshotRequest) (*emptypb.Empty, error)
	// Cancel attempts to cancel a task with the given operation ID. Not guaranteed to succeed.
	Cancel(context.Context, *types.Int64Value) (*emptypb.Empty, error)
	// GetLogs returns the keyed large data for the given operation.
	GetLogs(*LogDataRequest, Backrest_GetLogsServer) error
	// RunCommand executes a generic restic command on the repository.
	RunCommand(*RunCommandRequest, Backrest_RunCommandServer) error
	// GetDownloadURL returns a signed download URL given a forget operation ID.
	GetDownloadURL(context.Context, *types.Int64Value) (*types.StringValue, error)
	// Clears the history of operations
	ClearHistory(context.Context, *ClearHistoryRequest) (*emptypb.Empty, error)
	// PathAutocomplete provides path autocompletion options for a given filesystem path.
	PathAutocomplete(context.Context, *types.StringValue) (*types.StringList, error)
	mustEmbedUnimplementedBackrestServer()
}

// UnimplementedBackrestServer must be embedded to have forward compatible implementations.
type UnimplementedBackrestServer struct {
}

func (UnimplementedBackrestServer) GetFeatures(context.Context, *emptypb.Empty) (*FeatureList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatures not implemented")
}
func (UnimplementedBackrestServer) GetConfig(context.Context, *emptypb.Empty) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedBackrestServer) SetConfig(context.Context, *Config) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedBackrestServer) AddRepo(context.Context, *Repo) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRepo not implemented")
}
func (UnimplementedBackrestServer) GetOperationEvents(*emptypb.Empty, Backrest_GetOperationEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOperationEvents not implemented")
}
func (UnimplementedBackrestServer) GetOperations(context.Context, *GetOperationsRequest) (*OperationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedBackrestServer) ListSnapshots(context.Context, *ListSnapshotsRequest) (*ResticSnapshotList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (UnimplementedBackrestServer) ListSnapshotFiles(context.Context, *ListSnapshotFilesRequest) (*ListSnapshotFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshotFiles not implemented")
}
func (UnimplementedBackrestServer) Backup(context.Context, *types.StringValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedBackrestServer) DoRepoTask(context.Context, *DoRepoTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRepoTask not implemented")
}
func (UnimplementedBackrestServer) Forget(context.Context, *ForgetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forget not implemented")
}
func (UnimplementedBackrestServer) Restore(context.Context, *RestoreSnapshotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedBackrestServer) Cancel(context.Context, *types.Int64Value) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedBackrestServer) GetLogs(*LogDataRequest, Backrest_GetLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedBackrestServer) RunCommand(*RunCommandRequest, Backrest_RunCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (UnimplementedBackrestServer) GetDownloadURL(context.Context, *types.Int64Value) (*types.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadURL not implemented")
}
func (UnimplementedBackrestServer) ClearHistory(context.Context, *ClearHistoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearHistory not implemented")
}
func (UnimplementedBackrestServer) PathAutocomplete(context.Context, *types.StringValue) (*types.StringList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PathAutocomplete not implemented")
}
func (UnimplementedBackrestServer) mustEmbedUnimplementedBackrestServer() {}

// UnsafeBackrestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackrestServer will
// result in compilation errors.
type UnsafeBackrestServer interface {
	mustEmbedUnimplementedBackrestServer()
}

func RegisterBackrestServer(s grpc.ServiceRegistrar, srv BackrestServer) {
	s.RegisterService(&Backrest_ServiceDesc, srv)
}

func _Backrest_GetFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).GetFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_GetFeatures_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).GetFeatures(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).SetConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_AddRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).AddRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_AddRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).AddRepo(ctx, req.(*Repo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_GetOperationEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackrestServer).GetOperationEvents(m, &backrestGetOperationEventsServer{stream})
}

type Backrest_GetOperationEventsServer interface {
	Send(*OperationEvent) error
	grpc.ServerStream
}

type backrestGetOperationEventsServer struct {
	grpc.ServerStream
}

func (x *backrestGetOperationEventsServer) Send(m *OperationEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Backrest_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_GetOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_ListSnapshots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).ListSnapshots(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_ListSnapshotFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).ListSnapshotFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_ListSnapshotFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).ListSnapshotFiles(ctx, req.(*ListSnapshotFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_Backup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).Backup(ctx, req.(*types.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_DoRepoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRepoTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).DoRepoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_DoRepoTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).DoRepoTask(ctx, req.(*DoRepoTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_Forget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).Forget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_Forget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).Forget(ctx, req.(*ForgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).Restore(ctx, req.(*RestoreSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).Cancel(ctx, req.(*types.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackrestServer).GetLogs(m, &backrestGetLogsServer{stream})
}

type Backrest_GetLogsServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type backrestGetLogsServer struct {
	grpc.ServerStream
}

func (x *backrestGetLogsServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Backrest_RunCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackrestServer).RunCommand(m, &backrestRunCommandServer{stream})
}

type Backrest_RunCommandServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type backrestRunCommandServer struct {
	grpc.ServerStream
}

func (x *backrestRunCommandServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Backrest_GetDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).GetDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_GetDownloadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).GetDownloadURL(ctx, req.(*types.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_ClearHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).ClearHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_ClearHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).ClearHistory(ctx, req.(*ClearHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backrest_PathAutocomplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackrestServer).PathAutocomplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backrest_PathAutocomplete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackrestServer).PathAutocomplete(ctx, req.(*types.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// Backrest_ServiceDesc is the grpc.ServiceDesc for Backrest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backrest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Backrest",
	HandlerType: (*BackrestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeatures",
			Handler:    _Backrest_GetFeatures_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Backrest_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Backrest_SetConfig_Handler,
		},
		{
			MethodName: "AddRepo",
			Handler:    _Backrest_AddRepo_Handler,
		},
		{
			MethodName: "GetOperations",
			Handler:    _Backrest_GetOperations_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _Backrest_ListSnapshots_Handler,
		},
		{
			MethodName: "ListSnapshotFiles",
			Handler:    _Backrest_ListSnapshotFiles_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Backrest_Backup_Handler,
		},
		{
			MethodName: "DoRepoTask",
			Handler:    _Backrest_DoRepoTask_Handler,
		},
		{
			MethodName: "Forget",
			Handler:    _Backrest_Forget_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Backrest_Restore_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Backrest_Cancel_Handler,
		},
		{
			MethodName: "GetDownloadURL",
			Handler:    _Backrest_GetDownloadURL_Handler,
		},
		{
			MethodName: "ClearHistory",
			Handler:    _Backrest_ClearHistory_Handler,
		},
		{
			MethodName: "PathAutocomplete",
			Handler:    _Backrest_PathAutocomplete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOperationEvents",
			Handler:       _Backrest_GetOperationEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetLogs",
			Handler:       _Backrest_GetLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RunCommand",
			Handler:       _Backrest_RunCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/service.proto",
}
