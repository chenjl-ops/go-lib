package mysql_gorm

type DB struct {
	Host        string `json:"host"`                        // db连接地址
	Port        int    `json:"port"`                        // 端口
	UserName    string `json:"username"`                    // 用户名
	Password    string `json:"password"`                    // 密码
	DBName      string `json:"dbname"`                      // 库名
	Charset     string `json:"charset" default:"utf8mb4"`   // 字符集
	MaxIdleConn int    `json:"max_idle_conn" default:"10"`  // 最大空闲连接
	MaxOpenConn int    `json:"max_open_conn" default:"100"` // 最大连接
}

type Config struct {
	// ServerRunPort        string `envconfig:"SERVER_RUN_PORT" mapstructure:"server_run_port"`
	MysqlUserName string `envconfig:"MYSQL_USERNAME" mapstructure:"mysql_db_user" json:"MYSQL_USERNAME"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD" mapstructure:"mysql_db_passwd" json:"MYSQL_PASSWORD"`
	MysqlHost     string `envconfig:"MYSQL_HOST" mapstructure:"mysql_db_host" json:"MYSQL_HOST"`
	MysqlPort     int    `envconfig:"MYSQL_PORT" mapstructure:"mysql_db_port" json:"MYSQL_PORT"`
	MysqlDBName   string `envconfig:"MYSQL_DBNAME" mapstructure:"mysql_db_name" json:"MYSQL_DBNAME"`
	MysqlCharset  string `envconfig:"MYSQL_CHARSET" mapstructure:"mysql_charset" json:"MYSQL_CHARSET"`
}
