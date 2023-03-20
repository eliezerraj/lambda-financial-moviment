package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lambda-financial-moviment/internal/repository"
	"github.com/lambda-financial-moviment/internal/service"

)

var (
	logLevel		=	zerolog.DebugLevel // InfoLevel DebugLevel
	tableName		=	"agregation_card_person"
	version			=	"lambda-aggregation_person_card (github) version 1.5"

	movimentFinancialRepository	*repository.FinancialMovimentRepository
	financialMovimentService	*service.FinancialMovimentService
)

func getEnv(){
	if os.Getenv("TABLE_NAME") !=  "" {
		tableName = os.Getenv("TABLE_NAME")
	}
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.DebugLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init(){
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main() {
	log.Debug().Msg("main lambda-aggregation_person_card (go) v 1.5")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	movimentFinancialRepository, err := repository.NewFinancialMovimentRepository(tableName)
	if err != nil{
		return
	}

	financialMovimentService = service.NewFinancialMovimentService(*movimentFinancialRepository)
	
}