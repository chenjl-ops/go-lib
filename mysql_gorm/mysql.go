package mysql_gorm

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-lib/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Engine *gorm.DB

func NewDB() (*DB, error) {
	var config Config
	err := nacos.ReadRemoteConfig(&config)
	if nil != err {
		log.Fatal(err)
	}

	//nacos.Config = Config
	fmt.Sprintln("start mysql new config: ======")
	fmt.Sprintln("nacos config: ", config)
	result := &DB{
		UserName: config.MysqlUserName,
		Password: config.MysqlPassword,
		Port:     config.MysqlPort,
		DBName:   config.MysqlDBName,
		Host:     config.MysqlHost,
		Charset:  config.MysqlCharset,
	}

	fmt.Sprintln("end get mysql config: ", result)
	return result, nil
}

func (db *DB) InitMysql() error {
	fmt.Sprintln("start init mysql: =========")
	newEngine, err := db.NewConnect()
	if err != nil {
		return err
	}
	Engine = newEngine
	return nil
}

func (db *DB) NewConnect() (*gorm.DB, error) {
	fmt.Sprintln("db config: ", db)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db.UserName, db.Password, db.Host, db.Port, db.DBName, db.Charset)
	fmt.Sprintln("mysql connect str: ", dsn)
	gormEngine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	sqlDB, err := gormEngine.DB()
	if err != nil {
		return nil, err
	}
	if db.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(db.MaxOpenConn)
	}
	if db.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(db.MaxIdleConn)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return gormEngine, nil
}
