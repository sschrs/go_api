package models

import "time"

type Customer struct {
	CustomerID     string    `json:"customer_id" gorm:"type:varchar(50)" validate:"required"`
	Date           time.Time `json: "date" gorm:"type:date" validate:"required"`
	Segment        string    `json:"segment" gorm:"type:varchar(50)" validate:"required"`
	TransactionSum float32   `json:"transaction_sum" validate:"required"`
	CreatedDate    time.Time `json:"created_date" validate:"required"`
}
