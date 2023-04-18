GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

zip -jrm ../build/main.zip ../build/main

aws lambda update-function-code \
--function-name lambda-financial-moviment \
--zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-financial-moviment/build/main.zip \
--publish 

//------------------------

Endpoints

GET /version
GET financialmovimentbyperson/{901}
POST /financialmovimentbyperson

  {
        "account": "ACC-001",
        "person_id": "901",
        "currency": "BRL",
        "amount": 2222,
        "balance_type": "CREDIT"
  }

//------

APIGW ==> Lambda ==> CALL sign v4 ==> Lambda (Person) ==>
                 ==> DynamoDB (financial_moviment + person)

//-----