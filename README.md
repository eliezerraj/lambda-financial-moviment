# Lambda-financial-moviment

POC Lambda for technical purposes

Create/Get transaction account

Lambda retrieve data from PERSON (call another lambda) and agregated the PERSON data with FINANCIAL MOVIMENT (account statement)

The PERSON CALL use a sig v4 signature

Diagrama Flow

Manually compile the function

    APIGW ==> Lambda ==> APIGW CALL Lambda (Person) using sign v4 ==> DynamoDB (financial_moviment + person)

## Compile

    GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

    zip -jrm ../build/main.zip ../build/main

    aws lambda update-function-code \
    --function-name lambda-financial-moviment \
    --zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-financial-moviment/build/main.zip \
    --publish 

## Endpoints

GET /version

GET financialmovimentbyperson/{901}

    [
        {
            "id": "PERSON-010",
            "sk": "PERSON:PERSON-010#ACCOUNT:ACC-010#1681779976772",
            "account": "ACC-010",
            "person_id": "PERSON-010",
            "currency": "BRL",
            "amount": 150,
            "balance_type": "CREDIT",
            "create_at": "2023-04-18T01:06:15.410481435Z",
            "update_at": "0001-01-01T00:00:00Z"
        },
        {
            "id": "PERSON-010",
            "sk": "PERSON:PERSON-010#ACCOUNT:ACC-010#1681780020317",
            "account": "ACC-010",
            "person_id": "PERSON-010",
            "currency": "BRL",
            "amount": 80,
            "balance_type": "CREDIT",
            "create_at": "2023-04-18T01:06:15.410481435Z",
            "update_at": "0001-01-01T00:00:00Z"
        }
    ]

POST /financialmovimentbyperson

    {
          "account": "ACC-001",
          "person_id": "901",
          "currency": "BRL",
          "amount": 2222,
          "balance_type": "CREDIT"
    }

## Pipeline

## DynamoDB

    PERSON-100
    PERSON:PERSON-100#ACCOUNT:ACC-010#1689189531552
    ACC-010
    15
    CREDIT
    2023-07-12T19:16:47.142975592Z
    BRL
    PERSON-100
    0001-01-01T00:00:00Z

