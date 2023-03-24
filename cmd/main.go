package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lambda-financial-moviment/internal/adapter/handler"
	"github.com/lambda-financial-moviment/internal/repository"
	"github.com/lambda-financial-moviment/internal/service"
	"github.com/lambda-financial-moviment/internal/adapter/restapi"

)

var (
	logLevel		=	zerolog.DebugLevel // InfoLevel DebugLevel
	version			=	"lambda-aggregation_person_card (github) version 1.5"
	tableName 		= "financial_moviment"
	Url 			= "https://kfyn94nf42.execute-api.us-east-2.amazonaws.com"
	PersonPath 		= "/live/person"	

	rest_api					*restapi.FinancialMovimentRestApi
	response					*events.APIGatewayProxyResponse
	financialHandler			*handler.FinancialMovimentHandler
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
	log.Debug().Msg("main lambda-financial-moviment (go) v 1.0")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	movimentFinancialRepository, err := repository.NewFinancialMovimentRepository(tableName)
	if err != nil{
		return
	}

	rest_api, _ 				:= restapi.NewFinancialMovimentRestApi(Url, PersonPath)
	financialMovimentService 	= service.NewFinancialMovimentService(*movimentFinancialRepository, *rest_api)
	financialHandler 			= handler.NewFinancialMovimentHandler(*financialMovimentService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("handler")
	log.Debug().Msg("**************************")
	log.Debug().Str("req.Body", req.Body).
				Msg("APIGateway Request.Body")
	log.Debug().Msg("-*******************")

	switch req.HTTPMethod {
		case "GET":
			if (req.Resource == "/financialmovimentbyperson/{id}"){
				response, _ = financialHandler.GetFinancialMovimentByPerson(req)
			}else if (req.Resource == "/version"){
				response, _ = financialHandler.GetVersion(version)
			}else {
				response, _ = financialHandler.UnhandledMethod()
			}
		case "POST":
			if (req.Resource == "/financialmovimentbyperson"){
				response, _ = financialHandler.AddFinancialMovimentByPerson(req)
			}else {
				response, _ = financialHandler.UnhandledMethod()
			}
		case "DELETE":
			response, _ = financialHandler.UnhandledMethod()
		case "PUT":
			response, _ = financialHandler.UnhandledMethod()
		default:
			response, _ = financialHandler.UnhandledMethod()
	}

	return response, nil
}