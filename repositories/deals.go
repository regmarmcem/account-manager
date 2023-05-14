package repositories

import (
	"database/sql"
	"log"

	"github.com/regmarmcem/account-manager/models"
)

func SelectDealList(db *sql.DB) ([]models.Deal, error) {
	sqlStr := `
		select
			id, date, amount, category_id, description
		from
			deals;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Println("SelectDealList failed")
		return nil, err
	}
	defer rows.Close()

	dealArray := make([]models.Deal, 0)
	for rows.Next() {
		var deal models.Deal
		rows.Scan(&deal.ID, &deal.Date, &deal.Amount, &deal.CategoryID, &deal.Description)

		dealArray = append(dealArray, deal)
	}
	return dealArray, nil
}
