package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/account-manager/controllers"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	c := controllers.NewDealController(db)
	r.MethodFunc(http.MethodGet, "/deal/list", c.ListDeals)
	return r
}
