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
	appNameKey   = "RUNTIME_APP_NAME"
)

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
	return fmt.Sprintf("/nacos")
}

func GetNacosUrl(url string) string {
	//url := fmt.Sprintf("http://configserver-%s.xxx.cloud", env)
	if url != "" {
		return url
	}
	return fmt.Sprintf("https://nacos.default.com")
}

// ValidateNacosConf 验证nacos参数
func ValidateNacosConf(nacosConf *Nacos) error {

	appName := GetAppName(nacosConf.DataId)
	if appName == "" {
		return errors.Errorf("Env appName Has Empty")
	} else {
		nacosConf.DataId = appName
	}

	tenant := GetTenant(nacosConf.Tenant)
	if tenant == "" {
		return errors.Errorf("Env tenant Has Empty")
	} else {
		nacosConf.Tenant = tenant
	}

	group := GetGroup(nacosConf.Group)
	if group == "" {
		return errors.Errorf("Group Has Empty")
	} else {
		nacosConf.Group = group
	}

	nacosUrl := GetNacosUrl(nacosConf.Url)
	if nacosUrl == "" {
		return errors.Errorf("NacosUrl Has Empty")
	} else {
		nacosConf.Url = nacosUrl
	}

	path := GetNacosPath(nacosConf.Path)
	if path == "" {
		return errors.Errorf("Env path Has Empty")
	} else {
		nacosConf.Path = path
	}
	return nil
}
