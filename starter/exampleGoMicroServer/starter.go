package exampleGoMicroServer

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/micro/go-micro/server"
	"goapp/protobuf/generate/examplePb"
	"goapp/services"
	"github.com/micro/go-micro"
)


var Cfg *Config
type Config struct {
	Name       string
	ServerHost string
	ServerPort int64
}

func DefaultConfig() *Config {
	return &Config{
		Name:       "ExampleGoMicro",
		ServerHost: "127.0.0.1",
		ServerPort: 8889,
	}
}
func NewStarter() *starter {
	s := new(starter)
	s.cfg = &Config{}
	return s
}

type starter struct {
	goinfras.BaseStarter
	cfg        *Config
	microServer server.Server
}

func (s *starter) Name() string {
	return "ExampleGoMicro"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("GoMicro", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	s.cfg = &define
	Cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))

}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	service := micro.NewService(
		micro.Name(s.cfg.Name),
		micro.Address(fmt.Sprintf("%s:%d", s.cfg.ServerHost, s.cfg.ServerPort)),
		)

	service.Init()
	s.microServer = service.Server()

	// 注册服务端处理方法
	if err = examplePb.RegisterExampleMicroServiceHandler(s.microServer, &services.ExampleMicroService{});err != nil {
		sctx.Logger().Error(s.Name(),goinfras.StepSetup,err)
	}

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "GoMicro Service Server Setuped!")
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	if s.microServer != nil {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "GoMicro Service Server Steup Successful!")
		return true
	}
	return false
}

func (s *starter) Start(sctx *goinfras.StarterContext) {
	// 服务运行
	sctx.Logger().OK(s.Name(), goinfras.StepStart, fmt.Sprintf("GoMicro Service Server Startup ... Listening: %s:%d", s.cfg.ServerHost, s.cfg.ServerPort))
	if err := s.microServer.Start(); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	}
}

func (s *starter) StartBlocking() bool {
	return true
}

func (s *starter) Stop() {
	_ = s.microServer.Stop()

}

func (s *starter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *starter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}
