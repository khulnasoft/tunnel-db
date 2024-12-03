package vulnsrc

import (
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/alpine"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/amazon"
	archlinux "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/arch-linux"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/bundler"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/cargo"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/composer"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/debian"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/ghsa"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/glad"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/govulndb"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/node"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/nvd"
	oracleoval "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/oracle-oval"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/photon"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/python"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/redhat"
	redhatoval "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/redhat-oval"
	susecvrf "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/suse-cvrf"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/ubuntu"
)

type VulnSrc interface {
	Name() string
	Update(dir string) (err error)
}

var (
	// All holds all data sources
	All = []VulnSrc{
		// NVD
		nvd.NewVulnSrc(),

		// OS packages
		alpine.NewVulnSrc(),
		archlinux.NewVulnSrc(),
		redhat.NewVulnSrc(),
		redhatoval.NewVulnSrc(),
		debian.NewVulnSrc(),
		ubuntu.NewVulnSrc(),
		amazon.NewVulnSrc(),
		oracleoval.NewVulnSrc(),
		susecvrf.NewVulnSrc(susecvrf.SUSEEnterpriseLinux),
		susecvrf.NewVulnSrc(susecvrf.OpenSUSE),
		photon.NewVulnSrc(),

		// Language-specific packages
		bundler.NewVulnSrc(),
		composer.NewVulnSrc(),
		node.NewVulnSrc(),
		python.NewVulnSrc(),
		cargo.NewVulnSrc(),
		ghsa.NewVulnSrc(ghsa.Composer),
		ghsa.NewVulnSrc(ghsa.Maven),
		ghsa.NewVulnSrc(ghsa.Npm),
		ghsa.NewVulnSrc(ghsa.Nuget),
		ghsa.NewVulnSrc(ghsa.Pip),
		ghsa.NewVulnSrc(ghsa.Rubygems),
		glad.NewVulnSrc(),
		govulndb.NewVulnSrc(),
	}
)
