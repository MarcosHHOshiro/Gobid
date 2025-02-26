package api

import (
	"github.com/MarcosHHOshiro/Gobid/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}
