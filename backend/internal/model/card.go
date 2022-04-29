package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Card represents any card in application.
type Card struct {
	ID      string `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Lang    string `json:"lang" db:"lang"`
}

// Validate faq card's fields.
func (a Card) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Content, validation.Required),
		validation.Field(&a.Lang, validation.Required, languageRules),
	)
}
