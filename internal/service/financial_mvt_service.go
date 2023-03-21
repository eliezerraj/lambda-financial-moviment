package service

import (
	"fmt"

	"github.com/lambda-financial-moviment/internal/core/domain"

)

func (s *FinancialMovimentService) AddFinancialMoviment(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("AddFinancialMoviment")

	// Get financial moviment
	c, err := s.financialMovimentRepository.AddFinancialMoviment(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *FinancialMovimentService) GetFinancialMoviment(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("GetFinancialMoviment")

	// Get Person data
	person := domain.NewPerson(financialMoviment.PersonID,"","","")
	p, err := s.financialMovimentRestApi.GetPersonData(*person)
	if err != nil {
		return nil, err
	}

	fmt.Println("p => ", p)

	// Get financial moviment
	c, err := s.financialMovimentRepository.GetFinancialMoviment(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}