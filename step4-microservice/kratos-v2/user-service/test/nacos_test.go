package main

import (
	"fmt"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/nacos/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
	"user-service/internal/conf"
)

func TestNacosConfig(t *testing.T) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.5.130", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "a16dd0d7-08ef-4bcb-a64e-d6e10fab2f7e", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		//LogDir:              "../../tmp/nacos/log",
		//CacheDir:            "../../tmp/nacos/cache",
		RotateTime: "1h",
		MaxAge:     3,
		LogLevel:   "debug",
	}

	// 连接etcd配置中心
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatalln("连接配置中心失败: ", err)
	}

	// 读取配置中心的配置
	c := kconfig.New(
		kconfig.WithSource(
			config.NewConfigSource(client, config.Group("COMMON_GROUP"), config.DataID("mysql.datasource")),
			config.NewConfigSource(client, config.Group("USER_SERVICE_GROUP"), config.DataID("user.service.config1")),
		),
		kconfig.WithDecoder(func(kv *kconfig.KeyValue, v map[string]interface{}) error {
			fmt.Println("print from with decoder: ", v)
			return yaml.Unmarshal(kv.Value, v)
		}),
	)

	// 加载配置信息
	if err := c.Load(); err != nil {
		panic(err)
	}

	// 将配置信息绑定到conf实体类
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	fmt.Printf("bc: %#v\n", bc.Server.Http.Addr)
	fmt.Printf("bc: %#v\n", bc.Server.Http.Network)
	s, err := c.Value("server.http.addr").String()
	fmt.Println("server.http.addr: ", s)
}
