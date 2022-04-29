package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

var (
	minUserNameLength = 8
	maxUserNameLength = 255
	minPasswordLength = 8
	maxPasswordLength = 128
	maxTitleLength    = 100

	languageRules = validation.In("ru", "en")
	titleLength   = validation.Length(1, maxTitleLength)
)
