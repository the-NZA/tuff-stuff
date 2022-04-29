package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

type ImageRepo struct {
	db *sqlx.DB
}

// Create saves new image.
func (i ImageRepo) Create(image model.Image) (model.Image, error) {
	// Validate image struct
	if err := image.Validate(); err != nil {
		return image, err
	}

	// Save new image
	result, err := i.db.Exec(
		"INSERT INTO images(url) VALUES (?)",
		image.URL,
	)
	if err != nil {
		return image, err
	}

	// Parse new images id
	id, err := parseID(result)
	if err != nil {
		return image, err
	}

	image.ID = id

	return image, nil
}

// GetByID returns image by given ID.
func (i ImageRepo) GetByID(id string) (model.Image, error) {
	var (
		err   error
		image model.Image
	)

	err = i.db.Get(
		&image,
		"SELECT * FROM  images WHERE id = ?",
		id,
	)

	return image, err
}
