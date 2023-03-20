package repository

import(
	"os"

	"github.com/rs/zerolog/log"
	"github.com/lambda-financial-moviment/internal/erro"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var childLogger = log.With().Str("repository", "FinancialMovimentRepository").Logger()

type FinancialMovimentRepository struct {
	client 		dynamodbiface.DynamoDBAPI
	tableName   *string
}

func NewFinancialMovimentRepository(tableName string) (*FinancialMovimentRepository, error){
	childLogger.Debug().Msg("*** FinancialMoviment")

	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrOpenDatabase
	}

	return &FinancialMovimentRepository {
		client: dynamodb.New(awsSession),
		tableName: aws.String(tableName),
	}, nil
}
