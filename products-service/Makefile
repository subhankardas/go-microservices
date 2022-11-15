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
