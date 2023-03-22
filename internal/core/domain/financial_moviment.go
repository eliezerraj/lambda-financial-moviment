package domain

import (
	"time"

)

type FinancialMoviment struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	Account			string	`json:"account,omitempty"`
	PersonID		string  `json:"person_id,omitempty"`
	Currency		string  `json:"currency,omitempty"`
	Amount			float64 `json:"amount,omitempty"`
	BalanceType		string `json:"balance_type,omitempty"`
	CreateAt	time.Time 	`json:"create_at,omitempty"`
	UpdateAt	*time.Time 	`json:"update_at,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewFinancialMoviment(	id 			string, 
							sk 			string, 
							account		string,
							personId 	string,
							currency	string,
							amount		float64,
							balance_type	string,
							createAt	time.Time,
							updateAt	*time.Time,
							tenant	string) *FinancialMoviment{
	return &FinancialMoviment{
		ID:	id,
		SK:	sk,
		Account: account,
		PersonID: personId,
		Currency: currency,
		Amount: amount,
		BalanceType: balance_type,
		CreateAt: createAt, 
		UpdateAt: updateAt, 
		Tenant: tenant,
	}
}