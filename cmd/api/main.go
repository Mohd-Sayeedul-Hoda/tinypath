package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Mohd-Sayeedul-Hoda/url_shortner/internal/db"
	"io"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	ctx := context.Background()

	isTest := os.Getenv("GO_TEST") != ""
	var getenv func(string) string

	if isTest {
		// for testing we can set enviroment variable
		getenv = func(key string) string {
			switch key {
			case "MYAPP_FORMAT":
				return "markdown"
			case "MYAPP_TIMEOUT":
				return "5s"
			default:
				return ""
			}
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		getenv = func(key string) string {
			return os.Getenv(key)
		}
	}

	err := run(ctx, getenv, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context, getenv func(string) string, w io.Writer) error {

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var cfg config

	flag.IntVar(&cfg.port, "port", 6600, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment(production|staging|development)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", getenv("URLSHORTNER_DB_DSN"), "PostgreSQL dsn string")

	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "10m", "PostgerSQL max connection time")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 15, "PostgreSQL max open conncetions")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 15, "PostgreSQL max idle connections")

	flag.Parse()

	dbConfig := db.Config{
		DSN:          cfg.db.dsn,
		MaxOpenConns: cfg.db.maxOpenConns,
		MaxIdleConns: cfg.db.maxIdleConns,
		MaxIdleTime:  cfg.db.maxIdleTime,
	}

	conn, err := db.OpenDB(ctx, dbConfig)
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
