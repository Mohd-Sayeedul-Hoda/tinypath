package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/api"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/cache"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/cache/redis"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/config"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/db"
	jsonlog "github.com/Mohd-Sayeedul-Hoda/tinypath/internal/jsonLog"
	"github.com/Mohd-Sayeedul-Hoda/tinypath/internal/repository/postgres"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

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

	log := jsonlog.New(w, jsonlog.LevelInfo)

	conn, err := db.OpenDB(ctx, cfg)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.PrintInfo("database connection pool establisted", nil)

	urlRepo := postgres.NewURLShortenerRepo(conn)

	var cacheRepository cache.CacheRepo
	redisRepo, err := redis.NewCacheRepo(cfg)
	if err != nil {
		log.PrintError(err, map[string]string{"message": "Redis unavailable, using no-cache implementation"})
		cacheRepository = redis.NewNoCacheRepo(cfg)
	} else {
		cacheRepository = redisRepo
		log.PrintInfo("connected to redis server", nil)
	}

	srv := api.NewServer(cfg, log, urlRepo, cacheRepository)

	httpServer := http.Server{
		Addr:    net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler: srv,
	}

	go func() {
		log.PrintInfo("http server running", map[string]string{
			"addr": httpServer.Addr,
			"env":  cfg.Env,
		})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.PrintError(err, map[string]string{
				"message": "error listening and serving",
			})
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.PrintInfo("shutting down server", nil)

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.PrintError(err, map[string]string{
				"message": "error shutting down http server",
			})
		}
	}()

	wg.Wait()

	log.PrintInfo("stopped server", map[string]string{
		"addr": httpServer.Addr,
	})

	return nil
}
