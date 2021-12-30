DYNAMODBLOCAL_DIR ?= ~/Downloads/dynamodb_local_latest

build:
	go build

server:
	go run main.go server

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

dynamo:
	java -Djava.library.path=$(DYNAMODBLOCAL_DIR)/DynamoDBLocal_lib -jar $(DYNAMODBLOCAL_DIR)/DynamoDBLocal.jar -sharedDb