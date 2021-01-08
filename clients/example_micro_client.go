package clients

import (
	"fmt"
	"github.com/micro/go-micro"
	"goapp/protobuf/generate/examplePb"
	"goapp/starter/exampleMicro"
)

func GetExampleMicroClient() (examplePb.ExampleMicroService) {
	service := micro.NewService (
		micro.Name(exampleMicro.Cfg.Name),
		micro.Address(fmt.Sprintf("%s:%d", exampleMicro.Cfg.ServerHost, exampleMicro.Cfg.ServerPort)),
	)

	service.Init()

	return examplePb.NewExampleMicroService(exampleMicro.Cfg.Name, service.Client())
}
