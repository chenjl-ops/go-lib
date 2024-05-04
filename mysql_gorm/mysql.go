package mysql_gorm

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDbServer ...
func NewDbServer(conf *DB, engine *gorm.DB) (*DbServer, error) {
	result := &DbServer{
		Engine: engine,
		Config: conf,
	}
	return result, nil
}

// NewConnect ...
func (db *DB) NewConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db.UserName, db.Password, db.Host, db.Port, db.DBName, db.Charset)
	gormEngine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to database: ", err)
		return nil, err
	}
	sqlDB, err := gormEngine.DB()
	if err != nil {
		log.Error("Failed to init database engine: ", err)
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
		log.Error("Failed to ping database: ", err)
		return nil, err
	}

	return gormEngine, nil
}
