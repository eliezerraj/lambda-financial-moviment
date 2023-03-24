package handler

import(
	"github.com/rs/zerolog/log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/lambda-financial-moviment/internal/service"
	"github.com/lambda-financial-moviment/internal/erro"
	"github.com/lambda-financial-moviment/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"

)

var childLogger = log.With().Str("handler", "FinancialMovimentHandler").Logger()

var transactionSuccess	= "Transação com sucesso"
var time_create_at = time.Now()
var time_update_at = time.Time{}

type FinancialMovimentHandler struct {
	financialMoviment service.FinancialMovimentService
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type MessageBody struct {
	Msg *string `json:"message,omitempty"`
}

func (h *FinancialMovimentHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func NewFinancialMovimentHandler(financialMoviment service.FinancialMovimentService) *FinancialMovimentHandler{
	childLogger.Debug().Msg("NewFinancialMovimentHandler")

	return &FinancialMovimentHandler{
		financialMoviment: financialMoviment,
	}
}

func (h *FinancialMovimentHandler) AddFinancialMovimentByPerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("AddFinancialMovimentByPerson")

    var financialMoviment domain.FinancialMoviment
    if err := json.Unmarshal([]byte(req.Body), &financialMoviment); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.financialMoviment.AddFinancialMovimentByPerson(financialMoviment)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *FinancialMovimentHandler) GetFinancialMovimentByPerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetFinancialMovimentByPerson")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

	financialMoviment := domain.NewFinancialMoviment("",
													"",
													"",
													id,
													"",
													0,
													"",
													time_create_at,
													&time_update_at,
													"")

	response, err := h.financialMoviment.GetFinancialMovimentByPerson(*financialMoviment)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *FinancialMovimentHandler) GetVersion(version string) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetVersion")

	response := MessageBody { Msg: &version }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	
	return handlerResponse, nil
}
