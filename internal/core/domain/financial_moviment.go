package domain

import (
	"time"

)

type FinancialMoviment struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	PersonID		string  `json:"person_id,omitempty"`
	Currency		string  `json:"currency,omitempty"`
	Amount			float64 `json:"amount,omitempty"`
	CreateAt	time.Time 	`json:"create_at,omitempty"`
	UpdateAt	*time.Time 	`json:"update_at,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewFinancialMoviment(	id string, 
							sk 			string, 
							personId 	string,
							currency	string,
							amount		float64,
							createAt	time.Time,
							updateAt	*time.Time,
							tenant	string) *FinancialMoviment{
	return &FinancialMoviment{
		ID:	id,
		SK:	sk,
		PersonID: personId,
		Currency: currency,
		Amount: amount,
		CreateAt: createAt, 
		UpdateAt: updateAt, 
		Tenant: tenant,
	}
}