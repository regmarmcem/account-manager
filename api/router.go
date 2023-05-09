package api

import (
	"database/sql"

	"github.com/go-chi/chi"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	return r
}
