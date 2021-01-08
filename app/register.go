package main

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/spf13/viper"
	"goapp/starter/exampleGoMicroServer"
)

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter(viperCfg *viper.Viper) {
	goinfras.RegisterStarter(XGlobal.NewStarter())

	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// Register RPC Service Server
	goinfras.RegisterStarter(exampleGoMicroServer.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
