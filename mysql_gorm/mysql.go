package mysql_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDbServer ...
func NewDbServer(conf *DB, engine *gorm.DB) (*DbServer, error) {
	//fmt.Println("start new mysql server: ======")

	result := &DbServer{
		Engine: engine,
		Config: conf,
	}
	return result, nil
}

func (db *DB) NewConnect() (*gorm.DB, error) {
	//fmt.Println("db config: ", db)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db.UserName, db.Password, db.Host, db.Port, db.DBName, db.Charset)
	//fmt.Println("mysql connect str: ", dsn)
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
