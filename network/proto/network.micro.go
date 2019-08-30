// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/router/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Network service

type NetworkService interface {
	ListRoutes(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.ListResponse, error)
	ListNodes(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	ListNeighbours(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type networkService struct {
	c    client.Client
	name string
}

func NewNetworkService(name string, c client.Client) NetworkService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.network"
	}
	return &networkService{
		c:    c,
		name: name,
	}
}

func (c *networkService) ListRoutes(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.ListResponse, error) {
	req := c.c.NewRequest(c.name, "Network.ListRoutes", in)
	out := new(proto1.ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) ListNodes(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Network.ListNodes", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) ListNeighbours(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Network.ListNeighbours", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Network service

type NetworkHandler interface {
	ListRoutes(context.Context, *proto1.Request, *proto1.ListResponse) error
	ListNodes(context.Context, *ListRequest, *ListResponse) error
	ListNeighbours(context.Context, *ListRequest, *ListResponse) error
}

func RegisterNetworkHandler(s server.Server, hdlr NetworkHandler, opts ...server.HandlerOption) error {
	type network interface {
		ListRoutes(ctx context.Context, in *proto1.Request, out *proto1.ListResponse) error
		ListNodes(ctx context.Context, in *ListRequest, out *ListResponse) error
		ListNeighbours(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Network struct {
		network
	}
	h := &networkHandler{hdlr}
	return s.Handle(s.NewHandler(&Network{h}, opts...))
}

type networkHandler struct {
	NetworkHandler
}

func (h *networkHandler) ListRoutes(ctx context.Context, in *proto1.Request, out *proto1.ListResponse) error {
	return h.NetworkHandler.ListRoutes(ctx, in, out)
}

func (h *networkHandler) ListNodes(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.NetworkHandler.ListNodes(ctx, in, out)
}

func (h *networkHandler) ListNeighbours(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.NetworkHandler.ListNeighbours(ctx, in, out)
}
