package restapi

import(
	//"fmt"
	"net/http"
	"time"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/mitchellh/mapstructure"

	"github.com/lambda-financial-moviment/internal/core/domain"
	"github.com/lambda-financial-moviment/internal/erro"
)

var childLogger = log.With().Str("restApi", "FinancialMovimentRestpi").Logger()

type FinancialMovimentRestApi struct {
	Url	string
}

func NewFinancialMovimentRestApi(url string) (*FinancialMovimentRestApi, error){
	childLogger.Debug().Msg("*** FinancialMoviment")
	return &FinancialMovimentRestApi {
		Url: url,
	}, nil
}

func (r *FinancialMovimentRestApi) GetPersonData(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("getPerson")

	url := r.Url + "/" + person.ID
	person_interface, err :=makeGet(url, person)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Request")
		return nil, err
	}

	var person_result domain.Person
	err = mapstructure.Decode(person_interface, &person_result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error parse interface")
		return nil, erro.ErrParceInterface
    }
    
	return &person_result, nil
}

func makeGet(url string, inter interface{}) (interface{}, error) {
	childLogger.Debug().Msg("makeGet")

	client := &http.Client{Timeout: time.Second * 29}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Request")
		return false, erro.ErrHTTPRequest
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8");
	resp, err := client.Do(req)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Do Request")
		return false, erro.ErrHTTPRequest
	}

	switch (resp.StatusCode) {
		case 401:
			return false, erro.ErrHTTPForbiden
		case 200:
		case 404:
			return false, erro.ErrNotFound
		default:
			return false, erro.ErrNotFound
	}

	result := inter
	err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error no ErrUnmarshal")
		return false, erro.ErrUnmarshal
    }

	return result, nil
}
