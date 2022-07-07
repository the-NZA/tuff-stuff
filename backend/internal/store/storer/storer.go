package storer

import "github.com/the-NZA/tuff-stuff/backend/internal/model"

type Storer interface {
	Users() UserRepo
	ShoppingCards() ShoppingCardRepo
	CommissionCards() CommissionCardRepo
	HowWorksCards() HowWorksCardRepo
	WhyUsCards() WhyUsCardRepo
	Options() OptionRepo
	Homepages() HomepageRepo
	Images() ImageRepo
	GridImages() GridImageRepo
	Close() error
}

type UserRepo interface {
	Register(login, password string) (model.User, error)
	Login(login, password string) (model.User, error)
}

type ShoppingCardRepo interface {
	Create(card model.Card) (model.Card, error)
	GetByLang(lang string) ([]model.Card, error)
	GetByID(id string) (model.Card, error)
	Update(card model.Card) (model.Card, error)
	Delete(id string) error
}

type CommissionCardRepo interface {
	Create(card model.Card) (model.Card, error)
	GetByLang(lang string) ([]model.Card, error)
	GetByID(id string) (model.Card, error)
	Update(card model.Card) (model.Card, error)
	Delete(id string) error
}

type HowWorksCardRepo interface {
	Create(card model.Card) (model.Card, error)
	GetByLang(lang string) ([]model.Card, error)
	GetByID(id string) (model.Card, error)
	Update(card model.Card) (model.Card, error)
	Delete(id string) error
}

type WhyUsCardRepo interface {
	Create(card model.Card) (model.Card, error)
	GetByLang(lang string) ([]model.Card, error)
	GetByID(id string) (model.Card, error)
	Update(card model.Card) (model.Card, error)
	Delete(id string) error
}

type OptionRepo interface {
	Create(option model.Option) (model.Option, error)
	UpdateValue(option model.Option) (model.Option, error)
	UpdateMultiple(options []model.Option) ([]model.Option, error)
	GetAll() ([]model.Option, error)
}

type HomepageRepo interface {
	Create(page *model.Homepage) (*model.Homepage, error)
	GetByLang(lang string) (*model.Homepage, error)
	Update(page *model.Homepage) (*model.Homepage, error)
}

type ImageRepo interface {
	Create(image model.Image) (model.Image, error)
	GetByID(id string) (model.Image, error)
}

type GridImageRepo interface {
	GetAll() ([]model.GridImage, error)
	GetAllWithURLs() ([]model.GridImageWithURL, error)
	GetURLs() ([]string, error)
	Update(image model.GridImage) (model.GridImage, error)
}
