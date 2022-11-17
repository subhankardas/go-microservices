# ----- RUN SERVICE LOCALLY ----- #
run-service:
	go run main.go

# ----- RUN UNIT TESTS ----- #
run-tests:
	go test ./...

# ----- SWAGGER DOCS GENERATION ----- #
install-swagger:
	go get github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag

swagger-generate:
	swag init

# ----- API CLIENT GENERATION FROM SWAGGER ----- #
install-go-swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger

client-generate:
	swagger generate client -f ./docs/swagger.yaml -A products-service
	go mod tidy