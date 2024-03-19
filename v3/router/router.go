package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Debug bool
}

func InitChi(config Config) *chi.Mux {
	var router = chi.NewRouter()

	if config.Debug {
		router.Use(middleware.Logger)
	}

	return router
}
