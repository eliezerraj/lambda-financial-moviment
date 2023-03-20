package service

import (
	"github.com/rs/zerolog/log"

	"github.com/lambda-financial-moviment/internal/repository"

)

var childLogger = log.With().Str("service", "FinancialMovimentService").Logger()

type FinancialMovimentService struct {
	financialMovimentRepository repository.FinancialMovimentRepository
}

func NewFinancialMovimentService(financialMovimentRepository repository.FinancialMovimentRepository) *FinancialMovimentService{
	childLogger.Debug().Msg("NewFinancilMovimentService")

	return &FinancialMovimentService{
		financialMovimentRepository: financialMovimentRepository,
	}
}