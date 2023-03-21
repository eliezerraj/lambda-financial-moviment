package service

import (
	"fmt"

	"github.com/lambda-financial-moviment/internal/core/domain"

)

func (s *FinancialMovimentService) AddFinancialMovimentByPerson(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("AddFinancialMovimentByPerson")

	// Get Person data
	person := domain.NewPerson(financialMoviment.PersonID,"","","")
	p, err := s.financialMovimentRestApi.GetPersonData(*person)
	if err != nil {
		return nil, err
	}

	// Add financial moviment
	financialMoviment.PersonID = p.ID
	c, err := s.financialMovimentRepository.AddFinancialMovimentByPerson(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *FinancialMovimentService) GetFinancialMovimentByPerson(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("GetFinancialMoviment")

	// Get Person data
	person := domain.NewPerson(financialMoviment.PersonID,"","","")
	p, err := s.financialMovimentRestApi.GetPersonData(*person)
	if err != nil {
		return nil, err
	}

	fmt.Println("p => ", p)

	// Get financial moviment
	financialMoviment.PersonID = p.ID
	c, err := s.financialMovimentRepository.GetFinancialMovimentByPerson(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}