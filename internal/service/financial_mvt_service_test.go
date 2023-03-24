package service

import(
	"testing"
	"time"
	"github.com/rs/zerolog"
	"fmt"

	"github.com/lambda-financial-moviment/internal/core/domain"
	"github.com/lambda-financial-moviment/internal/repository"
	"github.com/lambda-financial-moviment/internal/adapter/restapi"

)

var (
	tableName = "financial_moviment"
	Url = "https://kfyn94nf42.execute-api.us-east-2.amazonaws.com"
	PersonPath = "/live/person"	
)

/*func TestAddFinancialMoviment(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	repository, err := repository.NewFinancialMovimentRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetFinancialMoviment Create Repository DynanoDB")
	}

	rest_api, _ := restapi.NewFinancialMovimentRestApi(Url,PersonPath)
	service	:= NewFinancialMovimentService(*repository, *rest_api)

	var time_create_at = time.Now()
	var time_update_at = time.Time{}
	data01 := domain.NewFinancialMoviment(	"",
											"",
											"ACC-001",
											"901",
											"BRL",
											453,
											"CREDIT",
											time_create_at,
											&time_update_at,
											"")

	fmt.Println(" data01  ",data01)
	result, err := service.AddFinancialMovimentByPerson(*data01)
	if err != nil {
		t.Errorf("Error -TestAddFinancialMoviment Access DynanoDB %v ", tableName)
	}

	data01.ID = "PERSON-901"
	if result == nil {
		t.Logf("Success on TestAddFinancialMoviment!!!")
	} else {
		if (data01.ID == result.ID){
			t.Logf("Success on TestAddFinancialMoviment!!! result : %v ", result)
		}else {
			t.Errorf("Error TestAddFinancialMoviment input : %v || result : %v "  , *data01, result)
		}
	}
}*/

func TestGetFinancialMoviment(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	repository, err := repository.NewFinancialMovimentRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetFinancialMoviment Create Repository DynanoDB")
	}

	rest_api, _ := restapi.NewFinancialMovimentRestApi(Url, PersonPath)
	service	:= NewFinancialMovimentService(*repository, *rest_api)

	var time_create_at = time.Now()
	var time_update_at = time.Time{}
	data01 := domain.NewFinancialMoviment(	"",
											"",
											"",
											"001",
											"",
											0,
											"",
											time_create_at,
											&time_update_at,
											"")

	fmt.Println(" data01  ",data01)
	result, err := service.GetFinancialMovimentByPerson(*data01)
	if err != nil {
		t.Errorf("Error -TestGetFinancialMoviment Access DynanoDB %v ", tableName)
	}

	if result == nil {
		t.Errorf("Error on TestGetFinancialMoviment!!!")
	} else {
		t.Logf("Success on TestGetFinancialMoviment!!! result : %v ", result)
	}
}
