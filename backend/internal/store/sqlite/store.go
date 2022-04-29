package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/the-NZA/tuff-stuff/backend/internal/config"
	"github.com/the-NZA/tuff-stuff/backend/internal/store/storer"
)

// Store SQLiteStore implements store.Storer interface.
type Store struct {
	db              *sqlx.DB
	users           *UserRepo
	shoppingCards   *ShoppingCardRepo
	commissionCards *CommissionCardRepo
	howWorksCards   *HowWorksCardRepo
	whyUsCards      *WhyUsCardRepo
	options         *OptionRepo
	homepages       *HomepageRepo
	images          *ImageRepo
	gridImages      *GridImageRepo
}

// NewStore returns Storer with given config.
func NewStore(c config.Config) (storer.Storer, error) {
	dbURL := fmt.Sprintf("file:%s?_foreign_keys=on", c.DbPath)
	db, err := sqlx.Connect("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

// Close store connection.
func (store *Store) Close() error {
	return store.db.Close()
}

// Users return corresponding repository.
func (store *Store) Users() storer.UserRepo {
	if store.users == nil {
		store.users = &UserRepo{
			db: store.db,
		}
	}

	return store.users
}

func (store *Store) ShoppingCards() storer.ShoppingCardRepo {
	if store.shoppingCards == nil {
		store.shoppingCards = &ShoppingCardRepo{
			db: store.db,
		}
	}

	return store.shoppingCards
}

func (store *Store) CommissionCards() storer.CommissionCardRepo {
	if store.commissionCards == nil {
		store.commissionCards = &CommissionCardRepo{
			db: store.db,
		}
	}

	return store.commissionCards
}

func (store *Store) HowWorksCards() storer.HowWorksCardRepo {
	if store.howWorksCards == nil {
		store.howWorksCards = &HowWorksCardRepo{
			db: store.db,
		}
	}

	return store.howWorksCards
}

func (store *Store) WhyUsCards() storer.WhyUsCardRepo {
	if store.whyUsCards == nil {
		store.whyUsCards = &WhyUsCardRepo{
			db: store.db,
		}
	}

	return store.whyUsCards
}

func (store *Store) Options() storer.OptionRepo {
	if store.options == nil {
		store.options = &OptionRepo{
			db: store.db,
		}
	}

	return store.options
}

func (store *Store) Homepages() storer.HomepageRepo {
	if store.homepages == nil {
		store.homepages = &HomepageRepo{
			db: store.db,
		}
	}

	return store.homepages
}

func (store *Store) Images() storer.ImageRepo {
	if store.images == nil {
		store.images = &ImageRepo{
			db: store.db,
		}
	}

	return store.images
}

func (store *Store) GridImages() storer.GridImageRepo {
	if store.gridImages == nil {
		store.gridImages = &GridImageRepo{
			db: store.db,
		}
	}

	return store.gridImages
}
