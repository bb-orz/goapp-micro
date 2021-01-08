### 应用示例
一个使用bb-orz/goinfras 搭建的GRpc应用脚手架


#### 目录结构

```
goapp       项目
|-- app             构建程序目录：应用主包，资源组件注册，makefile、dockerfile、sh等皆放置于此；
|-- client          GRPC Service Clients
|-- common          常量定义，错误处理定义、响应定义等；
|-- config          配置文件；
|-- dtos            数据传输对象dto定义；            
|-- protobuf        protobuf 协议定义及生成文件；
|-- services        protobuf服务协议的实现，在此调用core核心领域逻辑完成业务交互
|-- core            核心业务逻辑：服务层业务逻辑，以模块化方式封装，每个模块封装各自的领域设计、数据模型、数据访问、缓存等实现；
    |
    |-- moduleName1       核心业务模块
    |    |
    |    |-- _domain.go    模块领域层,领域驱动开发，实现具体业务逻辑
    |    |-- _model.go     模块数据模型层，定义领域内部的数据模型
    |    |-- _dao.go       模块数据访问层，定义具体的持久化数据访问操作
    |    |-- _cache.go     模块数据缓冲层，定义具体的缓存数据访问操作
    |    |-- ...          其它根据需要定义
    |    
    |-- moduleName2  
         |
         |-- ...  
    
```

#### app
 应用构建程序主目录：
 - main.go : 应用入口
 - register.go : 注册项目需要的 goinfras starter
 - const.go : 应用级常量定义；
 - ... : 其他应用级定义及设置； 

#### configs
 - 此处放置infras基础设施层中的资源组件配置文件，如需使用里面的组件只需编写响应的配置文件即可；
 - 配置文件名与组件同名，可以使用yaml/ini/json等数据格式。

#### dtos
 - 数据传输对象定义

#### protobuf
 - grpc 服务协议定义 
 - 可根据需要选择不同的rpc框架生成
 - 建议实现的service server做成启动器的方式运行，client做成连接的调用实例
 
#### services
 - 文件名称使用为可以描述其业务含义的单词
 - rpc_service为protobuf协议的实现

#### core 
 - core层为业务逻辑的具体实现；
 - 以模块划分,如user/order/verify/oauth/....每个模块集成一个目录；
 - 每个模块实现各自的领域层、数据模型层、数据访问层、缓存层等；
 - 每个模块中的领域层、数据访问层、缓存层各层级的错误传输需统一定义包装，最终传递给最外层统一处理；
 - 关于模块中各层之间的数据交换使用DTO数据传输对象；
 - 关于数据访问层DAO，可根据需要使用orm、sql builder、sql原生驱动实现；
 - 文件名称需满足可以描述其业务的英文；