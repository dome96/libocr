package keyvaluedatabase

import (
	"path/filepath"

	"github.com/cockroachdb/pebble"
)

// Legacy-compatible factory for Chainlink <= 2.28.0
// Redirects Badger factory calls to Pebble implementation.
func NewBadgerKeyValueDatabaseFactory(basePath string) Factory {
	return func(persistenceID string) (Database, error) {
		dbPath := filepath.Join(basePath, persistenceID)
		opts := &pebble.Options{}
		db, err := pebble.Open(dbPath, opts)
		if err != nil {
			return nil, err
		}
		return NewPebbleKeyValueDatabase(db), nil
	}
}

