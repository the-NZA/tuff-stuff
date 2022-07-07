package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Image represent each image in application.
type Image struct {
	ID  string `json:"id" db:"id"`
	URL string `json:"url" db:"url"`
}

// Validate image fields.
func (i Image) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.URL, validation.Required),
	)
}

// GridImage represent image in grid.
type GridImage struct {
	ID      string `json:"id" db:"id"`
	ImageID string `json:"image_id" db:"image_id"`
}

// Validate grid image fields.
func (i GridImage) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.ImageID, validation.Required),
	)
}

// GridImageWithURL represent grid image with image URL.
type GridImageWithURL struct {
	GridImage
	URL string `json:"url" db:"url"`
}

func (i GridImageWithURL) Validate() error {
	// Validate grid image fields.
	if err := i.GridImage.Validate(); err != nil {
		return err
	}

	// Validate image URL.
	return validation.ValidateStruct(&i,
		validation.Field(&i.URL, validation.Required),
	)
}
