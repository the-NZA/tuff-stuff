package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Homepage represents home page's content for different languages.
type Homepage struct {
	ID      string          `json:"id"`
	Lang    string          `json:"lang"`
	Content HomepageContent `json:"content"`
}

// HomepageContent includes all needed fields for homepage.
type HomepageContent struct {
	HeroTitle       string   `json:"hero_title"`
	HeroSubtitle    string   `json:"hero_subtitle"`
	AboutTitle      string   `json:"about_title"`
	ShoppingTitle   string   `json:"shopping_title"`
	HowWorksTitle   string   `json:"how_works_title"`
	CommissionTitle string   `json:"commission_title"`
	WhyUsTitle      string   `json:"why_us_title"`
	AboutText       []string `json:"about_text"`
}

// Validate whole homepage struct.
func (h *Homepage) Validate() error {
	// Validate lang
	if err := validation.ValidateStruct(h, validation.Field(
		&h.Lang,
		validation.Required,
		languageRules),
	); err != nil {
		return err
	}

	// Validate content
	return h.Content.Validate()
}

// Validate homepage content.
func (h *HomepageContent) Validate() error {
	// Validate homepage content fields
	return validation.ValidateStruct(h,
		validation.Field(
			&h.HeroTitle,
			validation.Required,
		),
		validation.Field(
			&h.HeroSubtitle,
			validation.Required,
		),
		validation.Field(
			&h.AboutTitle,
			validation.Required,
		),
		validation.Field(
			&h.AboutText,
			validation.Required,
		),
		validation.Field(
			&h.ShoppingTitle,
			validation.Required,
			titleLength,
		),
		validation.Field(
			&h.HowWorksTitle,
			validation.Required,
			titleLength,
		),
		validation.Field(
			&h.CommissionTitle,
			validation.Required,
			titleLength,
		),
		validation.Field(
			&h.WhyUsTitle,
			validation.Required,
			titleLength,
		),
	)
}
