package mysql_gorm

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
	result := &DB{
		UserName: MYSQL_USERNAME,
		Password: MYSQL_PASSWORD,
		Host:     MYSQL_HOST,
		Port:     MYSQL_PORT,
		DBName:   MYSQL_DBNAME,
		Charset:  MYSQL_CHARSET,
	}

	//fmt.Println("init mysql config: ", result)

	for _, opt := range opts {
		opt(result)
	}

	//fmt.Println("end get mysql config: ", result)

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
