package views

import "github.com/the-NZA/tuff-stuff/backend/internal/model"

type HomepageViewData struct {
	Homepage        *model.Homepage
	Options         map[string]model.Option
	ShoppingCards   []model.Card
	WhyUsCards      []model.Card
	HowWorksCards   []model.Card
	CommissionCards []model.Card
	GridImageURLs   []string
}
