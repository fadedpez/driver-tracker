DYNAMODBLOCAL_DIR ?= ~/dynamodblocal

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
	docker run \
		-p 8000:8000 \
		-v $(DYNAMODBLOCAL_DIR)/data:/home/dynamodblocal/data \
		amazon/dynamodb-local \
		-jar DynamoDBLocal.jar -sharedDb -dbPath /home/dynamodblocal/data/
