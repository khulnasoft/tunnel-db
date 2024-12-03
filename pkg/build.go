package pkg

import (
	"github.com/urfave/cli"
	"golang.org/x/xerrors"

	"go.khulnasoft.com/tunnel-db/pkg/db"
	"go.khulnasoft.com/tunnel-db/pkg/vulndb"
)

func build(c *cli.Context) error {
	cacheDir := c.String("cache-dir")
	if err := db.Init(cacheDir); err != nil {
		return err
	}

	targets := c.StringSlice("only-update")
	light := c.Bool("light")
	updateInterval := c.Duration("update-interval")

	dbType := db.TypeFull
	if light {
		dbType = db.TypeLight
	}

	vdb := vulndb.New(dbType, cacheDir, updateInterval)
	if err := vdb.Build(targets); err != nil {
		return xerrors.Errorf("build error: %w", err)
	}

	return nil

}
