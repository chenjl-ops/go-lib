package postgresql_grom

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBServer creates a new DbServer instance with the provided configuration.
func NewDBServer(config *DB, engine *gorm.DB) (*DbServer, error) {
	dbServer := &DbServer{
		Config: config,
		Engine: engine,
	}

	return dbServer, nil
}

// NewConnect establishes a new connection to the PostgreSQL database using the provided configuration.
func (db *DB) NewConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.UserName, db.Password, db.DBName, db.SSLMode)
	gormEngine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := gormEngine.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if db.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(db.MaxOpenConn)
	}
	if db.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(db.MaxIdleConn)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return gormEngine, nil
}

// AutoMigrate ...
func (db *DB) AutoMigrate(model interface{}) error {
	engine, err := db.NewConnect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	err = engine.AutoMigrate(model)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}
	fmt.Println("Database migrated successfully")
	return nil
}
