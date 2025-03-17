package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/db"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository/postgres"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type application struct {
	cfg     *config.Config
	logger  *jsonlog.Logger
	urlRepo repository.UrlShortener
	wg      sync.Mutex
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

	cfg := config.InitializeConfig(getenv)

	logger := jsonlog.New(w, jsonlog.LevelInfo)

	conn, err := db.OpenDB(ctx, cfg)
	if err != nil {
		return err
	}
	defer conn.Close()

	logger.PrintInfo("database connection pool establisted", nil)

	urlRepo := postgres.NewURLShortenerRepo(conn)

	app := &application{
		cfg:     cfg,
		logger:  logger,
		urlRepo: urlRepo,
	}

	err = app.serve()
	if err != nil {
		return err
	}

	return nil
}
