package services

import (
	"context"
	"goapp/protobuf/generate/examplePb"
)

// TODO 请实现 examplePb.ExampleRpcServiceServer 接口
type ExampleRpcService struct{}

func (*ExampleRpcService) Ping(context.Context, *examplePb.EmptyReqMsg, *examplePb.CommonRespMsg) error {
	panic("implement me")
}

func (*ExampleRpcService) Foo(context.Context, *examplePb.FooReqMsg, *examplePb.FooRespMsg) error {
	panic("implement me")
}

func (*ExampleRpcService) Bar(context.Context, *examplePb.BarReqMsg, *examplePb.BarRespMsg) error {
	panic("implement me")
}




