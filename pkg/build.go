package pkg

import (
	"github.com/urfave/cli"
	"go.khulnasoft.com/tunnel-db/pkg/db"
	"go.khulnasoft.com/tunnel-db/pkg/vulndb"
	"golang.org/x/xerrors"
)

func build(c *cli.Context) error {
	outputDir := c.String("output-dir")
	if err := db.Init(outputDir); err != nil {
		return xerrors.Errorf("db initialize error: %w", err)
	}

	cacheDir := c.String("cache-dir")
	targets := c.StringSlice("only-update")
	updateInterval := c.Duration("update-interval")

	vdb := vulndb.New(cacheDir, outputDir, updateInterval)
	if err := vdb.Build(targets); err != nil {
		return xerrors.Errorf("build error: %w", err)
	}

	return nil
}
