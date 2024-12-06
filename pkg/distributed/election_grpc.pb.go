// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: api/distributed.proto

package election

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Election_SubmitVote_FullMethodName     = "/election.Election/SubmitVote"
	Election_AggregateVotes_FullMethodName = "/election.Election/AggregateVotes"
)

// ElectionClient is the client API for Election service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionClient interface {
	SubmitVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	AggregateVotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
}

type electionClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionClient(cc grpc.ClientConnInterface) ElectionClient {
	return &electionClient{cc}
}

func (c *electionClient) SubmitVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, Election_SubmitVote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionClient) AggregateVotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Result)
	err := c.cc.Invoke(ctx, Election_AggregateVotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServer is the server API for Election service.
// All implementations must embed UnimplementedElectionServer
// for forward compatibility.
type ElectionServer interface {
	SubmitVote(context.Context, *VoteRequest) (*VoteResponse, error)
	AggregateVotes(context.Context, *Empty) (*Result, error)
	mustEmbedUnimplementedElectionServer()
}

// UnimplementedElectionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedElectionServer struct{}

func (UnimplementedElectionServer) SubmitVote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVote not implemented")
}
func (UnimplementedElectionServer) AggregateVotes(context.Context, *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregateVotes not implemented")
}
func (UnimplementedElectionServer) mustEmbedUnimplementedElectionServer() {}
func (UnimplementedElectionServer) testEmbeddedByValue()                  {}

// UnsafeElectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServer will
// result in compilation errors.
type UnsafeElectionServer interface {
	mustEmbedUnimplementedElectionServer()
}

func RegisterElectionServer(s grpc.ServiceRegistrar, srv ElectionServer) {
	// If the following call pancis, it indicates UnimplementedElectionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Election_ServiceDesc, srv)
}

func _Election_SubmitVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).SubmitVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Election_SubmitVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).SubmitVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Election_AggregateVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).AggregateVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Election_AggregateVotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).AggregateVotes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Election_ServiceDesc is the grpc.ServiceDesc for Election service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Election_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "election.Election",
	HandlerType: (*ElectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitVote",
			Handler:    _Election_SubmitVote_Handler,
		},
		{
			MethodName: "AggregateVotes",
			Handler:    _Election_AggregateVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/distributed.proto",
}
