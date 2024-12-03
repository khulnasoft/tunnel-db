package pkg

import (
	"time"

	"go.khulnasoft.com/tunnel-db/pkg/github"

	"go.khulnasoft.com/tunnel-db/pkg/utils"

	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc"
	"github.com/urfave/cli"
)

type AppConfig struct {
	Client github.VCSClientInterface
}

func (ac *AppConfig) NewApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "tunnel-db"
	app.Version = version
	app.Usage = "Tunnel DB builder"

	app.Commands = []cli.Command{
		{
			Name:   "build",
			Usage:  "build a database file",
			Action: build,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "light",
					Usage: "insert only CVE-ID and severity",
				},
				cli.StringSliceFlag{
					Name:  "only-update",
					Usage: "update db only specified distribution",
					Value: func() *cli.StringSlice {
						var targets cli.StringSlice
						for _, v := range vulnsrc.All {
							targets = append(targets, v.Name())
						}
						return &targets
					}(),
				},
				cli.StringFlag{
					Name:  "cache-dir",
					Usage: "cache directory path",
					Value: utils.CacheDir(),
				},
				cli.DurationFlag{
					Name:   "update-interval",
					Usage:  "update interval",
					Value:  24 * time.Hour,
					EnvVar: "UPDATE_INTERVAL",
				},
			},
		},
		{
			Name:   "upload",
			Usage:  "upload database files to GitHub Release",
			Action: ac.upload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir",
					Usage: "dir",
				},
			},
		},
	}

	return app
}
