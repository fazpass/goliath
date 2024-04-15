package module

import "github.com/go-chi/chi/v5"

type ModuleInterface interface {
	GetHttpRouter() *chi.Mux
}
