package api

import (
	"context"
	"fmt"
	"io"
	"os"
)

func main() {
	ctx := context.Background()
	err := run(ctx, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdin, "%s\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context, w io.Writer) error {

	return nil
}
