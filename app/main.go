package main

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XLogger"
	"github.com/bb-orz/goinfras/XValidate"
	"github.com/spf13/viper"
	"goapp/starter/exampleMicro"
)


func main() {
	// 初始化Viper配置加载器，导入配置，启动参数由命令行flag输入
	fmt.Println("Viper Config Loading  ......")
	viperCfg := goinfras.ViperLoader()

	// 注册应用组件启动器
	fmt.Println("Register Starters  ......")
	RegisterStarter(viperCfg)

	// 创建应用程序启动管理器
	app := goinfras.NewApplication(viperCfg)

	// 运行应用,启动已注册的资源组件
	fmt.Println("Application Starting ......")
	app.Up()
}

// 注册应用组件启动器，把基础设施各资源组件化
func RegisterStarter(viperCfg *viper.Viper) {
	goinfras.RegisterStarter(XGlobal.NewStarter())

	goinfras.RegisterStarter(XLogger.NewStarter())

	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// Register RPC Service Server
	goinfras.RegisterStarter(exampleMicro.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()

}
