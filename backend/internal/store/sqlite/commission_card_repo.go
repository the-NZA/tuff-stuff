package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	commissionCardCreate    = `INSERT INTO commission_cards(title, content, lang) VALUES(?, ?, ?)`
	commissionCardGetByLang = `SELECT * FROM commission_cards WHERE lang = ?`
	commissionCardGetByID   = `SELECT * FROM commission_cards WHERE id = ?`
	commissionCardUpdate    = `UPDATE commission_cards SET title = ?, content = ? WHERE id = ?`
	commissionCardDelete    = `DELETE FROM commission_cards WHERE id = ?`
)

type CommissionCardRepo struct {
	db *sqlx.DB
}

// Create saves new commission card.
func (cr *CommissionCardRepo) Create(card model.Card) (model.Card, error) {
	// Validate commission card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Insert new commission card
	result, err := cr.db.Exec(
		commissionCardCreate,
		card.Title,
		card.Content,
		card.Lang,
	)
	if err != nil {
		return card, err
	}

	// Get inserted ID
	id, err := parseID(result)
	if err != nil {
		return card, err
	}

	card.ID = id

	return card, nil
}

// GetByLang finds all commission cards by given lang.
func (cr *CommissionCardRepo) GetByLang(lang string) ([]model.Card, error) {
	var (
		err   error
		cards []model.Card
	)

	err = cr.db.Select(
		&cards,
		commissionCardGetByLang,
		lang,
	)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// GetByID find commission card with given ID.
func (cr *CommissionCardRepo) GetByID(id string) (model.Card, error) {
	var (
		err  error
		card model.Card
	)

	err = cr.db.Get(
		&card,
		commissionCardGetByID,
		id,
	)

	return card, err
}

func (cr *CommissionCardRepo) Update(card model.Card) (model.Card, error) {
	// Validate commission card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Update (replace) commission card by ID
	_, err := cr.db.Exec(
		commissionCardUpdate,
		card.Title,
		card.Content,
		card.ID,
	)

	return card, err
}

// Delete remove commission card with given ID.
func (cr *CommissionCardRepo) Delete(id string) error {
	_, err := cr.db.Exec(
		commissionCardDelete,
		id,
	)
	return err
}
