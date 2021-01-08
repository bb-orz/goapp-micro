package services

import (
	"context"
	"goapp/protobuf/generate/examplePb"
)

// TODO 请实现 examplePb.ExampleRpcServiceServer 接口
type ExampleMicroService struct{}

func (*ExampleMicroService) Ping(context.Context, *examplePb.EmptyReqMsg, *examplePb.CommonRespMsg) error {
	panic("implement me")
}

func (*ExampleMicroService) Foo(context.Context, *examplePb.FooReqMsg, *examplePb.FooRespMsg) error {
	panic("implement me")
}

func (*ExampleMicroService) Bar(context.Context, *examplePb.BarReqMsg, *examplePb.BarRespMsg) error {
	panic("implement me")
}




