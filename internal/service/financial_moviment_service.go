package service

import (
	"github.com/lambda-financial-moviment/internal/core/domain"

)

func (s *FinancialMovimentService) GetFinancialMoviment(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("GetFinancialMoviment")

	// Get financial moviment
	c, err := s.financialMovimentRepository.GetFinancialMoviment(financialMoviment)
	if err != nil {
		return nil, err
	}

	return c, nil
}