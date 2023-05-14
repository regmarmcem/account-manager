package models

import "time"

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DealType string `json:"deal_type"`
}

type MonthlyDealSummary struct {
	CategoryID int       `json:"category_id"`
	Month      time.Time `json:"month"`
	SumAmount  int       `json:"sum_amount"`
}

type Deal struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"deal_date"`
	Amount      int       `json:"amount"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
}
