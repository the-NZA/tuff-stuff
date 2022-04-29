package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Option represent each application option.
type Option struct {
	ID    string `json:"id"`
	Name  string `json:"name"`  // Internal slug
	Title string `json:"title"` // Public name
	Value string `json:"value"`
}

// Validate options fields.
func (o Option) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required),
		validation.Field(&o.Value, validation.Required),
	)
}
