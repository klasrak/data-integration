package data_integration

import "time"

// Negativation represents Navigation domain type
type Negativation struct {
	ID               string
	CompanyDocument  string    `json:"id"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}
