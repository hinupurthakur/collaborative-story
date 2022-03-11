include .env
clean:
	@go mod tidy
run:
	@go run main.go
checkrace:
	@go run -race main.go
build:
	@go build .
	
# GO111MODULE=on go get mvdan.cc/gofumpt
lint:
	@gofumpt -l -w .
test:
	@go test -v -count=1 ./...

migratedown:
	migrate -path db/migrations -database "${POSTGRES_DSN}" -verbose down ${VERSION}

migratefix:
	migrate -path db/migrations -database ${POSTGRES_DSN} force ${VERSION}

.PHONY: build

