package mysql_gorm

import (
	"fmt"
)

// default mysql config
const (
	MYSQL_USERNAME = "root"
	MYSQL_PASSWORD = "root"
	MYSQL_HOST     = "localhost"
	MYSQL_PORT     = 3306
	MYSQL_DBNAME   = "test"
	MYSQL_CHARSET  = "utf8mb4"
)

func NewDB(opts ...MysqlConfigOption) (*DB, error) {
	//var config Config
	//err := nacos.ReadRemoteConfig(&config)
	//if nil != err {
	//	log.Fatal(err)
	//}

	fmt.Sprintln("start mysql new config: ======")

	//nacos.Config = Config
	//fmt.Sprintln("nacos config: ", config)
	//result := &DB{
	//	UserName: config.MysqlUserName,
	//	Password: config.MysqlPassword,
	//	Port:     config.MysqlPort,
	//	DBName:   config.MysqlDBName,
	//	Host:     config.MysqlHost,
	//	Charset:  config.MysqlCharset,
	//}

	result := &DB{
		UserName: MYSQL_USERNAME,
		Password: MYSQL_PASSWORD,
		Host:     MYSQL_HOST,
		Port:     MYSQL_PORT,
		DBName:   MYSQL_DBNAME,
		Charset:  MYSQL_CHARSET,
	}

	fmt.Sprintln("init mysql config: ", result)

	for _, opt := range opts {
		opt(result)
	}

	fmt.Sprintln("end get mysql config: ", result)

	return result, nil
}

// MysqlConfigOption ...
type MysqlConfigOption func(*DB)

// WithUserName ...
func WithUserName(username string) MysqlConfigOption {
	return func(db *DB) {
		db.UserName = username
	}
}

// WithPassword ...
func WithPassword(password string) MysqlConfigOption {
	return func(db *DB) {
		db.Password = password
	}
}

// WithPort ...
func WithPort(port int) MysqlConfigOption {
	return func(db *DB) {
		db.Port = port
	}
}

// WithDBName ...
func WithDBName(dbName string) MysqlConfigOption {
	return func(db *DB) {
		db.DBName = dbName
	}
}

// WithHost ...
func WithHost(host string) MysqlConfigOption {
	return func(db *DB) {
		db.Host = host
	}
}

// WithCharset ...
func WithCharset(charset string) MysqlConfigOption {
	return func(db *DB) {
		db.Charset = charset
	}
}
