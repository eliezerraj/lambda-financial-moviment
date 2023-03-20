package domain

import (
	"time"

)

type FinancialMoviment struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	Person			string  `json:"person,omitempty"`
	CreateAt	*time.Time 	`json:"create_at,omitempty"`
	UpdateAt	*time.Time 	`json:"update_at,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewFinancialMoviment(id string, 
			sk 			string, 
			person 		string,
			createAt	*time.Time,
			updateAt	*time.Time,
			tenant	string) *FinancialMoviment{
	return &FinancialMoviment{
		ID:	id,
		SK:	sk,
		Person: person,
		CreateAt: createAt, 
		UpdateAt: updateAt, 
		Tenant: tenant,
	}
}