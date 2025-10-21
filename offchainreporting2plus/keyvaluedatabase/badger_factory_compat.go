//go:build !no_ocr2
// +build !no_ocr2

package keyvaluedatabase

import (
	"fmt"
	"path/filepath"
)

// NewBadgerKeyValueDatabaseFactory provides backward compatibility for older Chainlink code
// that expected a Badger DB factory. This wrapper uses Pebble under the hood instead.
func NewBadgerKeyValueDatabaseFactory(basePath string) func(configFilename string) (KeyValueDatabase, error) {
	return func(configFilename string) (KeyValueDatabase, error) {
		fullPath := filepath.Join(basePath, configFilename)
		db, err := NewPebbleKeyValueDatabase(fullPath, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create Pebble database at %s: %w", fullPath, err)
		}
		return db, nil
	}
}

