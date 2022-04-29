package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	shoppingCardCreate    = `INSERT INTO shopping_cards (title, content, lang) VALUES (?, ?, ?)`
	shoppingCardGetByLang = `SELECT * FROM shopping_cards WHERE lang = ?`
	shoppingCardGetByID   = `SELECT * FROM shopping_cards WHERE id = ?`
	shoppingCardUpdate    = `UPDATE shopping_cards SET title = ?, content = ? WHERE id = ?`
	shoppingCardDelete    = `DELETE FROM shopping_cards WHERE id = ?`
)

type ShoppingCardRepo struct {
	db *sqlx.DB
}

// Create saves new shopping card.
func (cr *ShoppingCardRepo) Create(card model.Card) (model.Card, error) {
	// Validate shopping card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Insert new shopping card
	result, err := cr.db.Exec(
		shoppingCardCreate,
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

// GetByLang finds all shopping cards by given lang.
func (cr *ShoppingCardRepo) GetByLang(lang string) ([]model.Card, error) {
	var (
		err   error
		cards []model.Card
	)

	err = cr.db.Select(
		&cards,
		shoppingCardGetByLang,
		lang,
	)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// GetByID find shopping card with given ID.
func (cr *ShoppingCardRepo) GetByID(id string) (model.Card, error) {
	var (
		err  error
		card model.Card
	)

	err = cr.db.Get(
		&card,
		shoppingCardGetByID,
		id,
	)

	return card, err
}

// Update updates shopping card's title and content with given ID.
func (cr *ShoppingCardRepo) Update(card model.Card) (model.Card, error) {
	// Validate shopping card
	if err := card.Validate(); err != nil {
		return card, err
	}

	// Update (replace) shopping card by ID
	_, err := cr.db.Exec(
		shoppingCardUpdate,
		card.Title,
		card.Content,
		card.ID,
	)

	return card, err
}

// Delete remove shopping card with given ID.
func (cr *ShoppingCardRepo) Delete(id string) error {
	_, err := cr.db.Exec(
		shoppingCardDelete,
		id,
	)
	return err
}
