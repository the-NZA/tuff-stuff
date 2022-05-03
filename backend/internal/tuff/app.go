package tuff

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/lgr"
	"github.com/the-NZA/tuff-stuff/backend/internal/config"
	"github.com/the-NZA/tuff-stuff/backend/internal/store/storer"
)

var (
	defaultTimeout = time.Second * 15
)

// Apper declares application interface.
type Apper interface {
	Start() error
	Stop(context.Context) error
}

// App contains all needed application's fields
// and implements Apper interface.
type App struct {
	server *http.Server
	router *chi.Mux
	store  storer.Storer
	logger lgr.L
	config config.Config
}

// NewApp creates new application with given parameters.
func NewApp(c config.Config, store storer.Storer) Apper {
	return &App{
		store:  store,
		config: c,
		router: chi.NewRouter(),
		logger: newLogger(c.IsDebugLogs),
	}
}

// newLogger returns new logger with debug or production config.
func newLogger(isDebug bool) lgr.L {
	if isDebug {
		return lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.LevelBraces)
	}

	return lgr.New(lgr.Msec, lgr.LevelBraces)
}

// configureRouter sets up application router.
func (app *App) configureRouter() {
	app.router.Mount("/", app.pagesRouter())
	app.router.Mount("/static", app.staticRouter())
	app.router.Mount("/api/v1", app.apiRouter())
	app.router.Mount("/auth/v1", app.authRouter())
}

// configureServer sets up application server.
func (app *App) configureServer() {
	app.server = &http.Server{
		Addr:         app.config.GetBindAddr(),
		Handler:      app.router,
		ReadTimeout:  defaultTimeout,
		WriteTimeout: defaultTimeout,
		IdleTimeout:  defaultTimeout,
	}
}

// Start performs configuration and starts application.
func (app *App) Start() error {
	app.configureRouter()
	app.configureServer()

	app.logger.Logf("[INFO] Application is starting at %s...\n", app.config.GetBindAddr())

	return app.server.ListenAndServe()
}

// Stop shutdowns application.
func (app *App) Stop(ctx context.Context) error {
	// Close db connection
	if err := app.store.Close(); err != nil {
		return err
	}

	return app.server.Shutdown(ctx)
}
