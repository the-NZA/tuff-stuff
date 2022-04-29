package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	whyUsCardCreate    = `INSERT INTO why_us_cards (title, content, lang) VALUES (?, ?, ?)`
	whyUsCardGetByLang = `SELECT * FROM why_us_cards WHERE lang = ?`
	whyUsCardGetByID   = `SELECT * FROM why_us_cards WHERE id = ?`
	whyUsCardUpdate    = `UPDATE why_us_cards SET title = ?, content = ? WHERE id = ?`
	whyUsCardDelete    = `DELETE FROM why_us_cards WHERE id = ?`
)

type WhyUsCardRepo struct {
	db *sqlx.DB
}

func (wu WhyUsCardRepo) Create(card model.Card) (model.Card, error) {
	// Validate why us  card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Insert new why us  card
	result, err := wu.db.Exec(
		whyUsCardCreate,
		card.Title,
		card.Content,
		card.Lang,
	)
	if err != nil {
		return card, err
	}

	// Convert ID to string
	id, err := parseID(result)
	if err != nil {
		return card, err
	}

	card.ID = id

	return card, nil
}

func (wu WhyUsCardRepo) GetByLang(lang string) ([]model.Card, error) {
	var (
		err   error
		cards []model.Card
	)

	err = wu.db.Select(
		&cards,
		whyUsCardGetByLang,
		lang,
	)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (wu WhyUsCardRepo) GetByID(id string) (model.Card, error) {
	var (
		err  error
		card model.Card
	)

	err = wu.db.Get(
		&card,
		whyUsCardGetByID,
		id,
	)

	return card, err
}

func (wu WhyUsCardRepo) Update(card model.Card) (model.Card, error) {
	// Validate why us  card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Update (replace) why us  card by ID
	_, err := wu.db.Exec(
		whyUsCardUpdate,
		card.Title,
		card.Content,
		card.ID,
	)

	return card, err
}

func (wu WhyUsCardRepo) Delete(id string) error {
	_, err := wu.db.Exec(
		whyUsCardDelete,
		id,
	)
	return err
}
