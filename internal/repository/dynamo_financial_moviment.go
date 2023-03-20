package repository

import(

	"github.com/lambda-financial-moviment/internal/core/domain"
	"github.com/lambda-financial-moviment/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func (r *FinancialMovimentRepository) Ping() (bool, error){
	return true, nil
}

/*func (r *FinancialMovimentRepository) AddAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("AddAgregation")

	item, err := dynamodbattribute.MarshalMap(agregation)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: r.tableName,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrInsert
	}

	_, err = r.client.TransactWriteItems(transaction)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrInsert
	}

	return &agregation ,nil
}*/

func (r *FinancialMovimentRepository) GetFinancialMoviment(financialMoviment domain.FinancialMoviment) (*domain.FinancialMoviment, error){
	childLogger.Debug().Msg("GetFinancialMoviment")

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(financialMoviment.ID)),
		expression.Key("sk").BeginsWith(financialMoviment.SK),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.QueryInput{
			TableName:                 r.tableName,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.Query(key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrQuery
	}

	final_result := []domain.FinancialMoviment{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &final_result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(final_result) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &final_result[0], nil
	}
}