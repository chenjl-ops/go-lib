package nacos

// NewNacosConfig 支持动态传参
func NewNacosConfig(opts ...ClientOption) (nacosConf *Nacos, err error) {
	nacosConf = &Nacos{
		Tenant:   "",
		Group:    "",
		DataId:   "",
		Url:      "",
		Port:     8848,
		Path:     "/nacos",
		LogDir:   "/tmp/nacos/log",
		CacheDir: "/tmp/nacos/cache",
		LogLevel: "debug",
	}

	for _, opt := range opts {
		opt(nacosConf)
	}

	err = ValidateNacosConf(nacosConf)
	if err != nil {
		return nil, err
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
