package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	runEnvKey    = "RUNTIME_ENV"
	runGroupKey  = "RUNTIME_GROUP"
	runTenantKey = "RUNTIME_TENANT"
	// runGROUPIdKey = "RUNTIME_GROUPID"
	appNameKey = "RUNTIME_APP_NAME"
)

//var Config *Specification

// GetEnv 获取当前运行环境
func GetEnv(name string) string {
	var env string
	if name != "" {
		env = name
	} else {
		env = os.Getenv(runEnvKey)
	}
	return env
}

// GetAppName 获取当前运行应用名称
func GetAppName(name string) string {
	var appName string
	if name != "" {
		appName = name
	} else {
		appName = os.Getenv(appNameKey)
	}
	return appName
}

// GetGroup 获取当前group
func GetGroup(name string) string {
	var groupName string
	if name != "" {
		groupName = name
	} else {
		groupName = os.Getenv(runGroupKey)
		if groupName == "" {
			groupName = "DEFAULT_GROUP"
		}
	}
	return groupName
}

// GetTenant 获取当前namespace
func GetTenant(name string) string {
	var tenant string
	if name != "" {
		tenant = name
	} else {
		tenant = os.Getenv(runTenantKey)
		if tenant == "" {
			tenant = "public"
		}
	}
	return tenant
}

// GetNacosPath 获取nacos路径
func GetNacosPath(path string) string {
	if path != "" {
		return fmt.Sprintf(path)
	}
	return fmt.Sprintf("/data/nacos")
}

func GetNacosUrl(url string) string {
	//url := fmt.Sprintf("http://configserver-%s.xxx.cloud", env)
	//url := fmt.Sprintf("http://kyc-prod-1-b349198798d93d66.elb.ap-southeast-1.amazonaws.com:8848/data/nacos")
	if url != "" {
		return url
	}
	return fmt.Sprintf("http://kyc-prod-1-b349198798d93d66.elb.ap-southeast-1.amazonaws.com")
}

func NewNacos() (nacos *Nacos, err error) {
	appName := GetAppName("")
	if appName == "" {
		return nil, errors.Errorf("Env appName Has Empty")
	}
	env := GetEnv("")
	if env == "" {
		return nil, errors.Errorf("Env env Has Empty")
	}
	tenant := GetTenant("")
	if tenant == "" {
		return nil, errors.Errorf("Env tenant Has Empty")
	}
	group := GetGroup("")
	if group == "" {
		return nil, errors.Errorf("Env group Has Empty")
	}
	nacosUrl := GetNacosUrl("")
	path := GetNacosPath("")
	if path == "" {
		return nil, errors.Errorf("Env path Has Empty")
	}
	nacos = &Nacos{
		Tenant:   tenant,
		Group:    group,
		DataId:   appName,
		Url:      nacosUrl,
		Path:     path,
		LogDir:   "/tmp/nacos/log",
		CacheDir: "/tmp/nacos/cache",
		LogLevel: "debug",
	}
	return nacos, nil
}

func (nacos *Nacos) GetNacosConfigs() (nacosClient *constant.ClientConfig, nacosServer *constant.ServerConfig, err error) {
	nacosServerConfigs := *constant.NewServerConfig(nacos.Url, 8848, constant.WithContextPath(nacos.Path))

	nacosClientConfigs := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(nacos.LogDir),
		constant.WithCacheDir(nacos.CacheDir),
		constant.WithLogLevel(nacos.LogLevel),
		//constant.WithUpdateCacheWhenEmpty(true),
	)

	return &nacosClientConfigs, &nacosServerConfigs, nil
}

func (nacos *Nacos) NewNacosBySdk() (iClient config_client.IConfigClient, err error) {
	nacosClientConfig, nacosServerConfig, err := nacos.GetNacosConfigs()
	if err != nil {
		return nil, err
	}
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  nacosClientConfig,
			ServerConfigs: []constant.ServerConfig{*nacosServerConfig},
		},
	)
}

func ReadRemoteConfig(data any) error {
	return ReadRemoteConfigCustom(nil, data)
}

func ReadRemoteConfigCustom(input *Nacos, v any) error {
	config, err := NewNacos()
	if err != nil {
		return err
	}

	if input != nil {
		if input.LogDir != "" {
			config.LogDir = input.LogDir
		}
		if input.CacheDir != "" {
			config.CacheDir = input.CacheDir
		}
		if input.LogLevel != "" {
			config.LogLevel = input.LogLevel
		}
		if input.Group != "" {
			config.Group = input.Group
		}
		if input.DataId != "" {
			config.DataId = input.DataId
		}
		if input.Tenant != "" {
			config.Tenant = input.Tenant
		}
	}

	// 判断logDir cacheDir logLevel是否为空
	if strings.TrimSpace(config.LogDir) == "" {
		config.LogDir = "/tmp/nacos/log"
	}
	if strings.TrimSpace(config.CacheDir) == "" {
		config.CacheDir = "/tmp/nacos/cache"
	}
	if strings.TrimSpace(config.LogLevel) == "" {
		config.LogLevel = "debug"
	}

	client, err := config.NewNacosBySdk()
	if err != nil {
		return err
	}

	//fmt.Println("config: ", config)
	//fmt.Println("Get Config Start: ===============")
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	})
	if err != nil {
		return err
	}
	//fmt.Println("nacos Data: ", content)

	err = json.Unmarshal([]byte(content), v)
	//fmt.Println("unmarshal done data: ", Config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//func ReadRemoteConfigCustom(input *Nacos) error {
//	config, err := NewNacos()
//	if err != nil {
//		return err
//	}
//
//	if input != nil {
//		if input.NacosServerUrl != "" {
//			config.NacosServerUrl = input.NacosServerUrl
//		}
//		if input.Group != "" {
//			config.Group = input.Group
//		}
//		if input.DataId != "" {
//			config.DataId = input.DataId
//		}
//		if input.Tenant != "" {
//			config.Tenant = input.Tenant
//		}
//	}
//	v := viper.New()
//	v.SetConfigType("prop")
//	// TODO List
//	// 1、实现nacos方法。 a.通过三方库 b. 自己注册nacos provider的实现方式
//	// 参考： https://github.com/yoyofxteam/nacos-viper-remote/tree/main
//	// 参考： https://github.com/nacos-group/nacos-sdk-go
//	err = v.AddRemoteProvider("nacos", config.NacosServerUrl, "/v1/cs/configs")
//	if err != nil {
//		return err
//	}
//	err = v.ReadRemoteConfig()
//	if err != nil {
//		return err
//	}
//	err = v.Unmarshal(&Config)
//	if err != nil {
//		return nil
//	}
//	fmt.Printf("nacos data: %+v", Config)
//	return nil
//}
