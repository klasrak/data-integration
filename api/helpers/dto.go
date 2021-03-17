package helpers

import "time"

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Negativation struct {
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}
