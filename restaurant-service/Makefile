# ----- INSTALL DEPENDENCIES ----- #
install-air:
	go install github.com/cosmtrek/air@latest
	air init

install-mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0
	go get github.com/golang/mock

# ----- RUN SERVICE LOCALLY ----- #
run-service-dev:
	air dev.config

run-service-qa:
	air qa.config

run-service-prod:
	air prod.config

# ----- GENERATE MOCKS ----- #
mock-core:
	mockgen -package mocks github.com/subhankardas/go-microservices/restaurant-service/core Database,Logger > mocks/core/mock.go

mock-data:
	mockgen -package mocks github.com/subhankardas/go-microservices/restaurant-service/data MenuData > mocks/data/mock.go

# ----- RUN INTEGRATION/UNIT TESTS ----- #
run-int-tests:
	go test -cover -count=1 -timeout 30s github.com/subhankardas/go-microservices/restaurant-service/controllers

run-unit-tests:
	go test -cover -timeout 30s \
	github.com/subhankardas/go-microservices/restaurant-service/core \
	github.com/subhankardas/go-microservices/restaurant-service/data \
	github.com/subhankardas/go-microservices/restaurant-service/middleware \
	github.com/subhankardas/go-microservices/restaurant-service/services \
	github.com/subhankardas/go-microservices/restaurant-service/utils
	