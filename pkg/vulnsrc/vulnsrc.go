package vulnsrc

import (
	"go.khulnasoft.com/tunnel-db/pkg/types"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/alma"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/alpine"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/amazon"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/azure"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/bitnami"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/bundler"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/chainguard"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/composer"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/debian"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/ghsa"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/glad"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/govulndb"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/k8svulndb"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/node"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/nvd"
	oracleoval "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/oracle-oval"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/photon"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/redhat"
	redhatoval "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/redhat-oval"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/rocky"
	susecvrf "go.khulnasoft.com/tunnel-db/pkg/vulnsrc/suse-cvrf"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/ubuntu"
	"go.khulnasoft.com/tunnel-db/pkg/vulnsrc/wolfi"
)

type VulnSrc interface {
	Name() types.SourceID
	Update(dir string) (err error)
}

var (
	// All holds all data sources
	All = []VulnSrc{
		// NVD
		nvd.NewVulnSrc(),

		// OS packages
		alma.NewVulnSrc(),
		alpine.NewVulnSrc(),
		redhat.NewVulnSrc(),
		redhatoval.NewVulnSrc(),
		debian.NewVulnSrc(),
		ubuntu.NewVulnSrc(),
		amazon.NewVulnSrc(),
		oracleoval.NewVulnSrc(),
		rocky.NewVulnSrc(),
		susecvrf.NewVulnSrc(susecvrf.SUSEEnterpriseLinux),
		susecvrf.NewVulnSrc(susecvrf.OpenSUSE),
		photon.NewVulnSrc(),
		azure.NewVulnSrc(azure.Azure),
		azure.NewVulnSrc(azure.Mariner),
		wolfi.NewVulnSrc(),
		chainguard.NewVulnSrc(),
		bitnami.NewVulnSrc(),

		k8svulndb.NewVulnSrc(),

		// Language-specific packages
		bundler.NewVulnSrc(),
		composer.NewVulnSrc(),
		node.NewVulnSrc(),
		ghsa.NewVulnSrc(),
		glad.NewVulnSrc(),
		govulndb.NewVulnSrc(), // For Go stdlib packages
	}
)
