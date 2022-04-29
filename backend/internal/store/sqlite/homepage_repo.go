package sqlite

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// HomepageRepo implements storer.HomepageRepo.
type HomepageRepo struct {
	db *sqlx.DB
}

type homepage struct {
	ID      string `db:"id"`
	Lang    string `db:"lang"`
	Content string `db:"content"`
}

// Create saves new home page.
func (h *HomepageRepo) Create(page *model.Homepage) (*model.Homepage, error) {
	// Validate homepage
	if err := page.Validate(); err != nil {
		return page, err
	}

	// Marshal new homepage
	jsonPage, err := json.Marshal(page)
	if err != nil {
		return page, err
	}

	// Save new homepage
	result, err := h.db.Exec(
		"INSERT INTO homepages (lang, content) VALUES (?, ?)",
		page.Lang,
		string(jsonPage),
	)
	if err != nil {
		return page, err
	}

	// Get inserted value ID
	id, err := parseID(result)
	if err != nil {
		return page, err
	}

	page.ID = id

	return page, nil
}

// GetByLang returns homepage by given lang.
func (h *HomepageRepo) GetByLang(lang string) (*model.Homepage, error) {
	var (
		dbPage homepage
		page   = &model.Homepage{}
		err    error
	)

	// Get homepage with language == lang variable
	err = h.db.Get(
		&dbPage,
		"SELECT * FROM homepages WHERE lang = ?",
		lang,
	)
	if err != nil {
		return page, err
	}

	// Unmarshal homepage
	err = json.Unmarshal([]byte(dbPage.Content), page)
	if err != nil {
		return page, err
	}

	page.ID = dbPage.ID

	return page, nil
}

// Update homepage by ID.
func (h *HomepageRepo) Update(homepage *model.Homepage) (*model.Homepage, error) {
	// Validate homepage
	if err := homepage.Validate(); err != nil {
		return homepage, err
	}

	// Marshal new homepage
	jsonPage, err := json.Marshal(homepage)
	if err != nil {
		return homepage, err
	}

	// Update (replace) homepage by ID
	_, err = h.db.Exec(
		"UPDATE homepages SET content = ? WHERE id = ?",
		jsonPage,
		homepage.ID,
	)

	return homepage, err
}
