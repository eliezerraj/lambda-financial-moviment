package restapi

import(
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"context"
	"crypto/sha256"

	"github.com/rs/zerolog/log"
	"github.com/mitchellh/mapstructure"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/lambda-financial-moviment/internal/core/domain"
	"github.com/lambda-financial-moviment/internal/erro"
)

var childLogger = log.With().Str("restApi", "FinancialMovimentRestpi").Logger()

type FinancialMovimentRestApi struct {
	Url			string
	PersonPath	string
}

func NewFinancialMovimentRestApi(url string, personpath string) (*FinancialMovimentRestApi, error){
	childLogger.Debug().Msg("*** NewFinancialMovimentRestApi")
	return &FinancialMovimentRestApi {
		Url: url,
		PersonPath: personpath,
	}, nil
}

func (r *FinancialMovimentRestApi) GetPersonData(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("GetPersonData")

	url := r.Url + r.PersonPath +"/" + person.ID
	//person_interface, err :=makeGet(url, person)
	person_interface, err :=makeGetAuthIAMRole(url, person)
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
	req.Header.Add("x-api-key", "NvbtUByZfK3PO0trigViL2OSQPlx8KTMa8pszPgt");

	resp, err := client.Do(req)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Do Request")
		return false, erro.ErrHTTPRequest
	}

	childLogger.Debug().Int("StatusCode :", resp.StatusCode).Msg("----")
	switch (resp.StatusCode) {
		case 401:
			return false, erro.ErrHTTPForbiden
		case 403:
			return false, erro.ErrHTTPForbiden
		case 200:
		case 400:
			return false, erro.ErrNotFound
		case 404:
			return false, erro.ErrNotFound
		default:
			return false, erro.ErrHTTPForbiden
	}

	result := inter
	err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error no ErrUnmarshal")
		return false, erro.ErrUnmarshal
    }

	return result, nil
}

func makeGetAuthIAMRole(url string, inter interface{}) (interface{}, error) {
	childLogger.Debug().Msg("makeGetAuthIAMRole")

	// ---------------------------
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		childLogger.Error().Err(err).Msg("error get Context")
		return false, erro.ErrHTTPRequest
	}
	credentials, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		childLogger.Error().Err(err).Msg("error get Credentials")
		return false, erro.ErrHTTPRequest
	}
	fmt.Println("***** ",credentials," === ",cfg.Region)
	//--------------------------

	client := &http.Client{Timeout: time.Second * 29}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Request")
		return false, erro.ErrHTTPRequest
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8");
	req.Header.Add("x-api-key", "NvbtUByZfK3PO0trigViL2OSQPlx8KTMa8pszPgt");
	//-------------
	hash := sha256.Sum256([]byte(""))
	hexHash := fmt.Sprintf("%x", hash)

	signer := v4.NewSigner()
	err = signer.SignHTTP(context.TODO(), credentials, req, hexHash, "execute-api", cfg.Region, time.Now())
	if err != nil {
		childLogger.Error().Err(err).Msg("error signer with credentials")
		return false, erro.ErrHTTPRequest
	}
	//-------------
	resp, err := client.Do(req)
	if err != nil {
		childLogger.Error().Err(err).Msg("error Do Request")
		return false, erro.ErrHTTPRequest
	}

	childLogger.Debug().Int("StatusCode :", resp.StatusCode).Msg("----")
	switch (resp.StatusCode) {
		case 401:
			return false, erro.ErrHTTPForbiden
		case 403:
			return false, erro.ErrHTTPForbiden
		case 200:
		case 400:
			return false, erro.ErrNotFound
		case 404:
			return false, erro.ErrNotFound
		default:
			return false, erro.ErrHTTPForbiden
	}

	result := inter
	err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error no ErrUnmarshal")
		return false, erro.ErrUnmarshal
    }

	return result, nil
}
