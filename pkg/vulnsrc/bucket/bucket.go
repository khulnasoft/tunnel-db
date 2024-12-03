package bucket

import (
	"fmt"

	"go.khulnasoft.com/tunnel-db/pkg/types"
)

const separator = "::"

func Name(ecosystem types.Ecosystem, dataSource string) string {
	return fmt.Sprintf("%s%s%s", ecosystem, separator, dataSource)
}
