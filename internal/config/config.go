package config

import "time"

const EnvPath = "local.env"

type AppConfig struct {
	LogLevel   string `envconfig:"LOG_LEVEL" required:"true"`
	Rest       Rest
	PostgresDB PostgresDB
	Service    Service
}

type Rest struct {
	ListenAddress string        `envconfig:"PORT" required:"true"`
	WriteTimeout  time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	ServerName    string        `envconfig:"SERVER_NAME" required:"true"`
	Token         string        `envconfig:"TOKEN" required:"true"`
}

type Service struct {
	PasswordSalt string `envconfig:"PASSWORD_SALT" required:"true"`
	Token        string `envconfig:"TOKEN" required:"true"`
}

type PostgresDB struct {
	Host                string        `envconfig:"DB_HOST" required:"true"`
	Port                int           `envconfig:"DB_PORT" required:"true"`
	Database            string        `envconfig:"DB_NAME" required:"true"`
	User                string        `envconfig:"DB_USER" required:"true"`
	Password            string        `envconfig:"DB_PASSWORD" required:"true"`
	SSLMode             string        `envconfig:"DB_SSL_MODE" required:"true"`
	PoolMaxConn         int           `envconfig:"DB_POOL_MAX_CONNS" required:"true"`
	PoolMaxConnLifeTime time.Duration `envconfig:"DB_POOL_MAX_CONN_LIFETIME" required:"true"`
	PoolMaxConnIdleTime time.Duration `envconfig:"DB_POOL_MAX_CONN_IDLE_TIME" required:"true"`
}
