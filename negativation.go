package data_integration

import "time"

// Negativation represents Navigation domain type
type Negativation struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyDocument  string    `json:"companyDocument,omitempty" bson:"companyDocument,omitempty"`
	CompanyName      string    `json:"companyName,omitempty" bson:"companyName,omitempty"`
	CustomerDocument string    `json:"customerDocument,omitempty" bson:"customerDocument,omitempty"`
	Value            float64   `json:"value,omitempty" bson:"value,omitempty"`
	Contract         string    `json:"contract" bson:"contract"`
	DebtDate         time.Time `json:"debtDate,omitempty" bson:"debtDate,omitempty"`
	InclusionDate    time.Time `json:"inclusionDate,omitempty" bson:"inclusionDate,omitempty"`
}
