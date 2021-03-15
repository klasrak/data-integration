package data_integration

import "time"

// Negativation represents Navigation domain type
type Negativation struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyDocument  string    `json:"companyDocument" bson:"companyDocument"`
	CompanyName      string    `json:"companyName" bson:"companyName"`
	CustomerDocument string    `json:"customerDocument" bson:"customerDocument"`
	Value            float64   `json:"value" bson:"value"`
	Contract         string    `json:"contract" bson:"contract"`
	DebtDate         time.Time `json:"debtDate" bson:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate" bson:"inclusionDate"`
}
