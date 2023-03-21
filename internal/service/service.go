package service

import (
	"github.com/rs/zerolog/log"

	"github.com/lambda-financial-moviment/internal/repository"
	"github.com/lambda-financial-moviment/internal/adapter/restapi"

)

var childLogger = log.With().Str("service", "FinancialMovimentService").Logger()

type FinancialMovimentService struct {
	financialMovimentRepository repository.FinancialMovimentRepository
	financialMovimentRestApi 	restapi.FinancialMovimentRestApi
}

func NewFinancialMovimentService(financialMovimentRepository repository.FinancialMovimentRepository,
								financialMovimentRestApi restapi.FinancialMovimentRestApi) *FinancialMovimentService{
	childLogger.Debug().Msg("NewFinancilMovimentService")

	return &FinancialMovimentService{
		financialMovimentRepository: financialMovimentRepository,
		financialMovimentRestApi: financialMovimentRestApi,
	}
}