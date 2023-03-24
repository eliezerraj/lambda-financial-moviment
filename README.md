GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

zip -jrm ../build/main.zip ../build/main

aws lambda update-function-code \
--function-name lambda-financial-moviment \
--zip-file fileb:///mnt/c/Eliezer/workspace/github.com/lambda-financial-moviment/build/main.zip \
--publish 