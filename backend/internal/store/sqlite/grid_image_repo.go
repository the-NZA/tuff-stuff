package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	gridImageGetAll  = `SELECT * FROM grid_images`
	gridImageUpdate  = `UPDATE grid_images SET image_id = ? WHERE id = ?`
	gridImageGetURLs = `SELECT images.url FROM image_grid JOIN images ON image_grid.image_id = images.id;`
)

type GridImageRepo struct {
	db *sqlx.DB
}

// GetAll returns all grid images.
func (r *GridImageRepo) GetAll() ([]model.GridImage, error) {
	var images []model.GridImage

	err := r.db.Select(
		&images,
		gridImageGetAll,
	)

	return images, err
}

// Update updates a grid image.
func (r *GridImageRepo) Update(image model.GridImage) (model.GridImage, error) {
	// Validate the image.
	if err := image.Validate(); err != nil {
		return image, err
	}

	// Update the image.
	_, err := r.db.Exec(
		gridImageUpdate,
		image.ImageID,
		image.ID,
	)

	return image, err
}

// GetURLs returns the URLs for the grid images.
func (r *GridImageRepo) GetURLs() ([]string, error) {
	var urls []string

	err := r.db.Select(
		&urls,
		gridImageGetURLs,
	)

	return urls, err
}
