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
	Url = "https://kfyn94nf42.execute-api.us-east-2.amazonaws.com/live/person"		
)

/*func TestAddFinancialMoviment(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	repository, err := repository.NewFinancialMovimentRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetFinancialMoviment Create Repository DynanoDB")
	}

	rest_api, _ := restapi.NewFinancialMovimentRestApi(Url)
	service	:= NewFinancialMovimentService(*repository, *rest_api)

	var time_create_at = time.Now()
	var time_update_at = time.Time{}
	data01 := domain.NewFinancialMoviment(	"001",
											"001",
											"Eliezer",
											"BRL",
											453,
											time_create_at,
											&time_update_at,
											"")

	fmt.Println(" data01  ",data01)
	result, err := service.AddFinancialMoviment(*data01)
	if err != nil {
		t.Errorf("Error -TestAddFinancialMoviment Access DynanoDB %v ", tableName)
	}

	if result == nil {
		t.Logf("Success on TestAddFinancialMoviment!!!")
	} else {
		if (data01.SK == result.SK){
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

	rest_api, _ := restapi.NewFinancialMovimentRestApi(Url)
	service	:= NewFinancialMovimentService(*repository, *rest_api)

	var time_create_at = time.Now()
	var time_update_at = time.Time{}
	data01 := domain.NewFinancialMoviment(	"",
											"",
											"902",
											"",
											0,
											time_create_at,
											&time_update_at,
											"")

	fmt.Println(" data01  ",data01)
	result, err := service.GetFinancialMoviment(*data01)
	if err != nil {
		t.Errorf("Error -TestGetFinancialMoviment Access DynanoDB %v ", tableName)
	}

	if result == nil {
		t.Logf("Success on TestGetFinancialMoviment!!!")
	} else {
		if (data01.SK == result.SK){
			t.Logf("Success on TestGetFinancialMoviment!!! result : %v ", result)
		}else {
			t.Errorf("Error TestGetFinancialMoviment input : %v || result : %v "  , *data01, result)
		}
	}
}
