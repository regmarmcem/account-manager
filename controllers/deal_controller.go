package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/regmarmcem/account-manager/repositories"
)

type DealController struct {
	db *sql.DB
}

func NewDealController(db *sql.DB) *DealController {
	return &DealController{db: db}
}

func (c *DealController) ListDeals(w http.ResponseWriter, r *http.Request) {
	deals, err := repositories.SelectDealList(c.db)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(deals)
}
