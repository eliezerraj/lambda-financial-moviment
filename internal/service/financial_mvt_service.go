package service

import (
	"fmt"
	"time"
	"strconv"

	"github.com/lambda-financial-moviment/internal/core/domain"

)

var time_create_at = time.Now()
var time_update_at = time.Time{}

func (s *FinancialMovimentService) AddFinancialMovimentByPerson(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("AddFinancialMovimentByPerson")

	// Get Person data
	person := domain.NewPerson(financialMoviment.PersonID,"","","")
	p, err := s.financialMovimentRestApi.GetPersonData(*person)
	if err != nil {
		return nil, err
	}

	// Add financial moviment
	financialMoviment.ID =  p.ID
	uuid := time.Now().UnixNano() / int64(time.Millisecond)
	financialMoviment.SK = "PERSON:" + p.ID + "#ACCOUNT:" + financialMoviment.Account + "#" + strconv.FormatInt(uuid, 10)
	financialMoviment.PersonID = p.ID
	financialMoviment.CreateAt = time_create_at
	financialMoviment.UpdateAt = &time_update_at

	c, err := s.financialMovimentRepository.AddFinancialMovimentByPerson(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *FinancialMovimentService) GetFinancialMovimentByPerson(financialMoviment domain.FinancialMoviment) (*[]domain.FinancialMoviment, error){
	childLogger.Debug().Msg("GetFinancialMoviment")

	// Get Person data
	person := domain.NewPerson(financialMoviment.PersonID,"","","")
	p, err := s.financialMovimentRestApi.GetPersonData(*person)
	if err != nil {
		return nil, err
	}

	fmt.Println("p => ", p)

	// Get financial moviment
	financialMoviment.ID =  p.ID

	financialMoviment.SK = "PERSON:" + p.ID + "#ACCOUNT:" + financialMoviment.Account 
	c, err := s.financialMovimentRepository.GetFinancialMovimentByPerson(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}