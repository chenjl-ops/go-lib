package nacos

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
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

func NewNacosConfig(opts ...ClientOption) (nacosConf *Nacos, err error) {
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
	nacosConf = &Nacos{
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

	for _, opt := range opts {
		opt(nacosConf)
	}
	return nacosConf, nil
}

// ClientOption ...
type ClientOption func(*Nacos)

// WithTenant ...
func WithTenant(tenant string) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.Tenant = tenant
	}
}

// WithGroup ...
func WithGroup(group string) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.Group = group
	}
}

// WithDataId ...
func WithDataId(dataId string) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.DataId = dataId
	}
}

// WithUrl ...
func WithUrl(url string) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.Url = url
	}
}

// WithPort ...
func WithPort(port uint64) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.Port = port
	}
}

// WithPath ...
func WithPath(path string) ClientOption {
	return func(nacosConf *Nacos) {
		nacosConf.Path = path
	}
}
