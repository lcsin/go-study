package main

import (
	"fmt"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/nacos/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
	"user-service/internal/conf"
)

func main() {

}

func TestConfigFile(t *testing.T) {
	c1 := kconfig.New(
		kconfig.WithSource(file.NewSource("../configs/bootstrap.yaml")),
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
	fmt.Println("nacos addr: ", bc1.Nacos.Addr)
	fmt.Println("nacos port: ", bc1.Nacos.Port)

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(bc1.Nacos.Addr, uint64(bc1.Nacos.Port)),
	}

	cc := constant.ClientConfig{
		NamespaceId:         bc1.Nacos.NamespaceId, //namespace id
		TimeoutMs:           uint64(bc1.Nacos.TimeoutMS),
		NotLoadCacheAtStart: bc1.Nacos.NotLoadCacheAtStart,
		LogDir:              bc1.Nacos.LogDir,
		CacheDir:            bc1.Nacos.CacheDir,
		RotateTime:          bc1.Nacos.RotateTime,
		MaxAge:              bc1.Nacos.MaxAge,
		LogLevel:            bc1.Nacos.LogLevel,
	}

	// 连接nacos配置中心
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatalln("连接配置中心失败: ", err)
	} else {
		log.Println("连接配置中心成功")
	}

	// 读取配置中心的配置
	c2 := kconfig.New(
		kconfig.WithSource(
			config.NewConfigSource(client, config.Group("COMMON_GROUP"), config.DataID("mysql.datasource")),
			config.NewConfigSource(client, config.Group("USER_SERVICE_GROUP"), config.DataID("user.service.config1")),
		),
		kconfig.WithDecoder(func(kv *kconfig.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c2.Load(); err != nil {
		panic(err)
	}

	var bc2 conf.Bootstrap
	if err := c2.Scan(&bc2); err != nil {
		panic(err)
	}
	fmt.Printf("bc2: %#v\n", bc2.Server)
	fmt.Printf("bc2: %#v\n", bc2.Data)
	s, err := c2.Value("server.http.addr").String()
	fmt.Println("server.http.addr: ", s)
}
