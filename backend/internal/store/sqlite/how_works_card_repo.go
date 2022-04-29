package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	howWorksCardCreate    = `INSERT INTO how_works_cards (title, content, lang) VALUES (?, ?, ?)`
	howWorksCardGetByLang = `SELECT * FROM how_works_cards WHERE lang = ?`
	howWorksCardGetByID   = `SELECT * FROM how_works_cards WHERE id = ?`
	howWorksCardUpdate    = `UPDATE how_works_cards SET title = ?, content = ? WHERE id = ?`
	howWorksCardDelete    = `DELETE FROM how_works_cards WHERE id = ?`
)

type HowWorksCardRepo struct {
	db *sqlx.DB
}

// Create saves new how works  card.
func (hw *HowWorksCardRepo) Create(card model.Card) (model.Card, error) {
	// Validate how works  card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Insert new how works  card
	result, err := hw.db.Exec(
		howWorksCardCreate,
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

// GetByLang finds all how works  cards by given lang.
func (hw *HowWorksCardRepo) GetByLang(lang string) ([]model.Card, error) {
	var (
		err   error
		cards []model.Card
	)

	err = hw.db.Select(
		&cards,
		howWorksCardGetByLang,
		lang,
	)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// GetByID find how works  card with given ID.
func (hw *HowWorksCardRepo) GetByID(id string) (model.Card, error) {
	var (
		err  error
		card model.Card
	)

	err = hw.db.Get(
		&card,
		howWorksCardGetByID,
		id,
	)

	return card, err
}

func (hw *HowWorksCardRepo) Update(card model.Card) (model.Card, error) {
	// Validate how works  card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Update (replace) how works  card by ID
	_, err := hw.db.Exec(
		howWorksCardUpdate,
		card.Title,
		card.Content,
		card.ID,
	)

	return card, err
}

// Delete remove how works  card with given ID.
func (hw *HowWorksCardRepo) Delete(id string) error {
	_, err := hw.db.Exec(
		howWorksCardDelete,
		id,
	)
	return err
}
