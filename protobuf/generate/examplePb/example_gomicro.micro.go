// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: example_gomicro.proto

/*
Package examplePb is a generated protocol buffer package.

It is generated from these files:
	example_gomicro.proto

It has these top-level messages:
	EmptyReqMsg
	CommonRespMsg
	FooReqMsg
	FooRespMsg
	BarReqMsg
	BarRespMsg
*/
package examplePb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ExampleRpcService service

type ExampleRpcService interface {
	Ping(ctx context.Context, in *EmptyReqMsg, opts ...client.CallOption) (*CommonRespMsg, error)
	Foo(ctx context.Context, in *FooReqMsg, opts ...client.CallOption) (*FooRespMsg, error)
	Bar(ctx context.Context, in *BarReqMsg, opts ...client.CallOption) (*BarRespMsg, error)
}

type exampleRpcService struct {
	c    client.Client
	name string
}

func NewExampleRpcService(name string, c client.Client) ExampleRpcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "examplerpcservice"
	}
	return &exampleRpcService{
		c:    c,
		name: name,
	}
}

func (c *exampleRpcService) Ping(ctx context.Context, in *EmptyReqMsg, opts ...client.CallOption) (*CommonRespMsg, error) {
	req := c.c.NewRequest(c.name, "ExampleRpcService.Ping", in)
	out := new(CommonRespMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleRpcService) Foo(ctx context.Context, in *FooReqMsg, opts ...client.CallOption) (*FooRespMsg, error) {
	req := c.c.NewRequest(c.name, "ExampleRpcService.Foo", in)
	out := new(FooRespMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleRpcService) Bar(ctx context.Context, in *BarReqMsg, opts ...client.CallOption) (*BarRespMsg, error) {
	req := c.c.NewRequest(c.name, "ExampleRpcService.Bar", in)
	out := new(BarRespMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExampleRpcService service

type ExampleRpcServiceHandler interface {
	Ping(context.Context, *EmptyReqMsg, *CommonRespMsg) error
	Foo(context.Context, *FooReqMsg, *FooRespMsg) error
	Bar(context.Context, *BarReqMsg, *BarRespMsg) error
}

func RegisterExampleRpcServiceHandler(s server.Server, hdlr ExampleRpcServiceHandler, opts ...server.HandlerOption) error {
	type exampleRpcService interface {
		Ping(ctx context.Context, in *EmptyReqMsg, out *CommonRespMsg) error
		Foo(ctx context.Context, in *FooReqMsg, out *FooRespMsg) error
		Bar(ctx context.Context, in *BarReqMsg, out *BarRespMsg) error
	}
	type ExampleRpcService struct {
		exampleRpcService
	}
	h := &exampleRpcServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExampleRpcService{h}, opts...))
}

type exampleRpcServiceHandler struct {
	ExampleRpcServiceHandler
}

func (h *exampleRpcServiceHandler) Ping(ctx context.Context, in *EmptyReqMsg, out *CommonRespMsg) error {
	return h.ExampleRpcServiceHandler.Ping(ctx, in, out)
}

func (h *exampleRpcServiceHandler) Foo(ctx context.Context, in *FooReqMsg, out *FooRespMsg) error {
	return h.ExampleRpcServiceHandler.Foo(ctx, in, out)
}

func (h *exampleRpcServiceHandler) Bar(ctx context.Context, in *BarReqMsg, out *BarRespMsg) error {
	return h.ExampleRpcServiceHandler.Bar(ctx, in, out)
}