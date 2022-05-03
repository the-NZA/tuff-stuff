package tuff

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
	"github.com/the-NZA/tuff-stuff/backend/internal/views"
)

const (
	siteLangCookieName = "tuffStuffLang"
)

var (
	defaultLangCookie = &http.Cookie{
		Name:   siteLangCookieName,
		MaxAge: 0,
		Value:  "ru",
	}
)

// pagesRouter returns sub router for pages.
func (app *App) pagesRouter() http.Handler {
	router := chi.NewRouter()

	router.Get("/", app.rootPageHandler())
	router.Get("/{lang:[a-z]{2}}", app.homePageHandler())

	router.NotFound(app.pageNotFoundHandler())

	return router
}

// rootPageHandler performs language detection and corresponding redirection.
func (app *App) rootPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get language cookie
		langCookie, err := r.Cookie(siteLangCookieName)
		if err != nil {
			// Set default language cookie
			http.SetCookie(w, defaultLangCookie)

			// Redirect to default language
			http.Redirect(w, r, "/ru", http.StatusSeeOther)
			return
		}

		// Redirect based on language cookie value
		switch langCookie.Value {
		case "ru":
			http.Redirect(w, r, "/ru", http.StatusSeeOther)
			return

		case "en":
			http.Redirect(w, r, "/en", http.StatusSeeOther)
			return
		}
	}
}

// homePageHandler handles homepages for different languages.
func (app *App) homePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := chi.URLParam(r, "lang")
		if len(lang) != langParamLength {
			app.logger.Logf("[ERROR] Invalid language param: %s", lang)

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get homepage text content for the language.
		homepage, err := app.store.Homepages().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during homepage searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get shopping cards for the language.
		shoppingCards, err := app.store.ShoppingCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping cards searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get why us cards for the language.
		whyUsCards, err := app.store.WhyUsCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during why us cards searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get how works cards for the language.
		howWorksCards, err := app.store.HowWorksCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during how works cards searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get commission cards for the language.
		commissionCards, err := app.store.CommissionCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during commission cards searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get options map for the language.
		options, err := func(app *App) (map[string]model.Option, error) {
			// Get all options.
			options, internalError := app.store.Options().GetAll()
			if internalError != nil {
				app.logger.Logf("[ERROR] during options searching :%s\n", internalError.Error())

				return nil, internalError
			}

			// Create options map.
			optionsMap := make(map[string]model.Option, len(options))
			for i := range options {
				optionsMap[options[i].Name] = options[i]
			}

			return optionsMap, nil
		}(app)
		if err != nil {
			app.logger.Logf("[ERROR] during options searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Get grid image urls.
		gridImageUrls, err := app.store.GridImages().GetURLs()
		if err != nil {
			app.logger.Logf("[ERROR] during grid image urls searching :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}

		// Render homepage template.
		err = views.Homepage(w, &views.HomepageViewData{
			Homepage:        homepage,
			ShoppingCards:   shoppingCards,
			WhyUsCards:      whyUsCards,
			HowWorksCards:   howWorksCards,
			CommissionCards: commissionCards,
			Options:         options,
			GridImageURLs:   gridImageUrls,
		})
		if err != nil {
			app.logger.Logf("[ERROR] during homepage rendering :%s\n", err.Error())

			app.pageNotFoundHandler()(w, r)
			return
		}
	}
}

// pageNotFoundHandler response when no page are found for given url.
func (app *App) pageNotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		res := fmt.Sprintf("<h1>Page can't be found with given url</h1><p>URL: %s</p>", r.URL)
		_, err := w.Write([]byte(res))

		if err != nil {
			app.logger.Logf("[ERROR] during not found writing %v", err)
		}
	}
}
