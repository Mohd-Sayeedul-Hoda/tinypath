package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	jsonlog "github.com/Mohd-Sayeedul-Hoda/url_shortner/internal/jsonLog"
)

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
		getenv = func(key string) string {
			return os.Getenv(key)
		}
	}
	err := run(ctx, getenv, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdin, "%s\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context, getenv func(string) string, w io.Writer) error {

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := jsonlog.New(w, jsonlog.LevelInfo)

	return nil
}
