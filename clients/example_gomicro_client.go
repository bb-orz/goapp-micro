package clients

import (
	"fmt"
	"github.com/micro/go-micro"
	"goapp/protobuf/generate/examplePb"
	"goapp/starter/exampleGoMicroServer"
)

func GetExampleGoMicroClient() (examplePb.ExampleRpcService) {
	service := micro.NewService (
		micro.Name(exampleGoMicroServer.Cfg.Name),
		micro.Address(fmt.Sprintf("%s:%d", exampleGoMicroServer.Cfg.ServerHost, exampleGoMicroServer.Cfg.ServerPort)),
	)

	service.Init()

	return examplePb.NewExampleRpcService(exampleGoMicroServer.Cfg.Name, service.Client())
}
