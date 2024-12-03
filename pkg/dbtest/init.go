package dbtest

import (
	"testing"

	fixtures "github.com/khulnasoft-lab/bolt-fixtures"
	"github.com/stretchr/testify/require"
	"go.khulnasoft.com/tunnel-db/pkg/db"
)

func InitDB(t *testing.T, fixtureFiles []string) string {
	t.Helper()

	// Create a temp dir
	dbDir := t.TempDir()
	dbPath := db.Path(dbDir)

	// Load testdata into BoltDB
	loader, err := fixtures.New(dbPath, fixtureFiles)
	require.NoError(t, err)
	require.NoError(t, loader.Load())
	require.NoError(t, loader.Close())

	// Initialize DB
	require.NoError(t, db.Init(dbDir))

	return dbDir
}
