package nacos

type Nacos struct {
	LogDir   string
	CacheDir string
	LogLevel string
	Url      string
	Path     string
	Tenant   string
	DataId   string
	Group    string
}

//type Specification struct {
//	// ServerRunPort        string `envconfig:"SERVER_RUN_PORT" mapstructure:"server_run_port"`
//	MysqlUserName string `envconfig:"MYSQL_USERNAME" mapstructure:"mysql_db_user" json:"MYSQL_USERNAME"`
//	MysqlPassword string `envconfig:"MYSQL_PASSWORD" mapstructure:"mysql_db_passwd" json:"MYSQL_PASSWORD"`
//	MysqlHost     string `envconfig:"MYSQL_HOST" mapstructure:"mysql_db_host" json:"MYSQL_HOST"`
//	MysqlPort     int    `envconfig:"MYSQL_PORT" mapstructure:"mysql_db_port" json:"MYSQL_PORT"`
//	MysqlDBName   string `envconfig:"MYSQL_DBNAME" mapstructure:"mysql_db_name" json:"MYSQL_DBNAME"`
//	MysqlCharset  string `envconfig:"MYSQL_CHARSET" mapstructure:"mysql_charset" json:"MYSQL_CHARSET"`
//}
