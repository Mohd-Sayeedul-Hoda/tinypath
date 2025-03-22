package config

import (
	"flag"
)

type Config struct {
	Port int
	Host string
	Env  string
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Cache struct {
		DSN string
	}
}

func InitializeConfig(getenv func(string) string) *Config {

	var cfg Config

	flag.IntVar(&cfg.Port, "port", 6600, "API server port")
	flag.StringVar(&cfg.Host, "host", "localhost", "API server host")
	flag.StringVar(&cfg.Env, "env", "development", "Enviroment(production|staging|development)")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", getenv("URLSHORTNER_DB_DSN"), "PostgreSQL dsn string")
	flag.StringVar(&cfg.Cache.DSN, "cache-dsn", getenv("REDIS_DSN"), "Redis dsn string")

	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "10m", "PostgerSQL max connection time")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 15, "PostgreSQL max open conncetions")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 15, "PostgreSQL max idle connections")

	flag.Parse()

	return &cfg
}
