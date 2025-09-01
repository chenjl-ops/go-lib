package postgresql_grom

// defaults postgresql config
const (
	HOST          = "localhost"
	PORT          = 5432
	USERNAME      = "postgres"
	PASSWORD      = "postgres"
	DBNAME        = "postgres"
	CHARSET       = "utf8"
	COLLATION     = "utf8_general_ci"
	MAX_IDLE_CONN = 10
	MAX_OPEN_CONN = 100
	SSL_MODE      = "disable"
)

// NewDB ...
func NewDB(opts ...PostgresConfigOption) (*DB, error) {
	result := &DB{
		UserName:    USERNAME,
		Password:    PASSWORD,
		Host:        HOST,
		Port:        PORT,
		DBName:      DBNAME,
		Charset:     CHARSET,
		Collation:   COLLATION,
		MaxIdleConn: MAX_IDLE_CONN,
		MaxOpenConn: MAX_OPEN_CONN,
		SSLMode:     SSL_MODE,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result, nil
}

type PostgresConfigOption func(*DB)

// WithHost ...
func WithHost(host string) PostgresConfigOption {
	return func(db *DB) {
		db.Host = host
	}
}

// WithPort ...
func WithPort(port int) PostgresConfigOption {
	return func(db *DB) {
		db.Port = port
	}
}

// WithUserName ...
func WithUserName(username string) PostgresConfigOption {
	return func(db *DB) {
		db.UserName = username
	}
}

// WithPassword ...
func WithPassword(password string) PostgresConfigOption {
	return func(db *DB) {
		db.Password = password
	}
}

// WithDBName ...
func WithDBName(dbName string) PostgresConfigOption {
	return func(db *DB) {
		db.DBName = dbName
	}
}

// WithCharset ...
func WithCharset(charset string) PostgresConfigOption {
	return func(db *DB) {
		db.Charset = charset
	}
}

// WithCollation ...
func WithCollation(collation string) PostgresConfigOption {
	return func(db *DB) {
		db.Collation = collation
	}
}

// WithMaxIdleConn ...
func WithMaxIdleConn(maxIdleConn int) PostgresConfigOption {
	return func(db *DB) {
		db.MaxIdleConn = maxIdleConn
	}
}

// WithMaxOpenConn ...
func WithMaxOpenConn(maxOpenConn int) PostgresConfigOption {
	return func(db *DB) {
		db.MaxOpenConn = maxOpenConn
	}
}

// WithSSLMode ...
func WithSSLMode(sslMode string) PostgresConfigOption {
	return func(db *DB) {
		db.SSLMode = sslMode
	}
}
