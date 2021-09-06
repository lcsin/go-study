package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
	"user-service/internal/conf"

	"github.com/go-kratos/kratos/v2"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/nacos/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string

	id string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout))

	// 获取配置文件
	c1 := kconfig.New(
		kconfig.WithSource(file.NewSource(flagConf)),
	)
	// 加载配置信息
	if err := c1.Load(); err != nil {
		panic(err)
	}
	// 将配置信息绑定到实体
	var bc1 conf.Bootstrap
	if err := c1.Scan(&bc1); err != nil {
		panic(err)
	}

	// nacos 配置信息
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(bc1.Nacos.Addr, bc1.Nacos.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         bc1.Nacos.NamespaceId, //namespace id
		TimeoutMs:           bc1.Nacos.TimeoutMS,
		NotLoadCacheAtStart: bc1.Nacos.NotLoadCacheAtStart,
		LogDir:              bc1.Nacos.LogDir,
		CacheDir:            bc1.Nacos.CacheDir,
		RotateTime:          bc1.Nacos.RotateTime,
		MaxAge:              bc1.Nacos.MaxAge,
		LogLevel:            bc1.Nacos.LogLevel,
	}

	// 连接 nacos 配置中心
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		logger.Log(log.LevelError, "配置中心连接失败: ", err)
	} else {
		logger.Log(log.LevelInfo, "连接配置中心成功")
	}

	// 获取 nacos 配置中心配置
	c2 := kconfig.New(
		kconfig.WithSource(
			config.NewConfigSource(client, config.Group("COMMON_GROUP"), config.DataID("mysql.datasource")),
			config.NewConfigSource(client, config.Group("USER_SERVICE_GROUP"), config.DataID("user.service.config1")),
		),
		kconfig.WithDecoder(func(kv *kconfig.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	// 加载配置信息
	if err := c2.Load(); err != nil {
		panic(err)
	}
	// 绑定配置信息到实体类
	var bc2 conf.Bootstrap
	if err := c2.Scan(&bc2); err != nil {
		panic(err)
	}

	Name, Version = bc2.Server.Name, bc2.Server.Version
	logger.Log(log.LevelInfo, "ServiceInfo", "ServerName: "+Name+" ServerVersion: "+Version)

	// 初始化应用
	app, cleanup, err := initApp(bc2.Server, bc2.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
