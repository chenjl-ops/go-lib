package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
)

// NewNacos ...
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
		Port:     8848,
		Path:     path,
		LogDir:   "/tmp/nacos/log",
		CacheDir: "/tmp/nacos/cache",
		LogLevel: "debug",
	}
	return nacos, nil
}

func (nacos *Nacos) GetNacosConfigs() (nacosClient *constant.ClientConfig, nacosServer *constant.ServerConfig, err error) {
	nacosServerConfigs := *constant.NewServerConfig(nacos.Url, nacos.Port, constant.WithContextPath(nacos.Path))

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

func (nacos *Nacos) ReadRemoteConfig(data any) error {
	return readRemoteConfigCustom(nacos, data)
}

func readRemoteConfigCustom(input *Nacos, value any) error {
	client, err := input.NewNacosBySdk()
	if err != nil {
		return err
	}

	//fmt.Println("config: ", config)
	//fmt.Println("Get Config Start: ===============")
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: input.DataId,
		Group:  input.Group,
	})
	if err != nil {
		return err
	}
	//fmt.Println("nacos Data: ", content)

	err = json.Unmarshal([]byte(content), value)
	//fmt.Println("unmarshal done data: ", value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
