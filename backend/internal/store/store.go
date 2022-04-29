package store

import (
	"github.com/the-NZA/tuff-stuff/backend/internal/config"
	"github.com/the-NZA/tuff-stuff/backend/internal/store/sqlite"
	"github.com/the-NZA/tuff-stuff/backend/internal/store/storer"
)

// NewStore returns new Storer with given config.
func NewStore(c config.Config) (storer.Storer, error) {
	return sqlite.NewStore(c)
}
