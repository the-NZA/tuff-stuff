package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

type OptionRepo struct {
	db *sqlx.DB
}

// Create save new option.
func (o OptionRepo) Create(option model.Option) (model.Option, error) {
	// Validate option
	if err := option.Validate(); err != nil {
		return option, err
	}

	// Insert new option
	result, err := o.db.Exec(
		"INSERT INTO options (name, title, value) VALUES (?, ?, ?)",
		option.Name,
		option.Title,
		option.Value,
	)
	if err != nil {
		return option, err
	}

	// Get new option ID
	id, err := parseID(result)
	if err != nil {
		return option, err
	}

	option.ID = id

	return option, nil

}

// UpdateValue patch given option's value by ID.
func (o *OptionRepo) UpdateValue(option model.Option) (model.Option, error) {
	// Validate option
	if err := option.Validate(); err != nil {
		return option, err
	}

	// Update option's value
	_, err := o.db.Exec(
		"UPDATE options SET value = ? WHERE id = ?",
		option.Value,
		option.ID,
	)

	return option, err
}

// GetAll returns all options.
func (o OptionRepo) GetAll() ([]model.Option, error) {
	var (
		err     error
		options []model.Option
	)

	err = o.db.Select(
		&options,
		"SELECT * FROM options",
	)
	if err != nil {
		return nil, err
	}

	return options, nil
}
