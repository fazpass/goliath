package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
