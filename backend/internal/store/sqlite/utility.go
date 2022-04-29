package sqlite

import (
	"database/sql"
	"errors"
	"strconv"
)

const idBase = 10

var (
	ErrEmptySQLResult = errors.New("empty sql result")
)

// parseID converts last inserted ID into string representation.
func parseID(result sql.Result) (string, error) {
	// Check for nil result
	if result == nil {
		return "", ErrEmptySQLResult
	}

	// Get inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	// Convert ID to string
	return strconv.FormatInt(id, idBase), nil
}
