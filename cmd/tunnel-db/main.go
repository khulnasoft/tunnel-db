package main

import (
	"context"
	"log"
	"os"

	"go.khulnasoft.com/tunnel-db/pkg/github"

	"go.khulnasoft.com/tunnel-db/pkg"
)

var (
	version = "0.0.1"
)

func main() {
	ctx := context.Background()
	ac := pkg.AppConfig{
		Client: github.NewClient(ctx),
	}

	app := ac.NewApp(version)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
