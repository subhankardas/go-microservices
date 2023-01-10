# go-microservices

The aim of this repository is to explore, learn and develop production grade web-services that are scalable and reliable and implement various microservices patterns.

## Projects

1. **products-service** - Simple RESTful service implementation with the following features.
   1. Uses native library and *Gorilla Mux* for serving HTTP requests.
   2. Middleware for input validation with custom field validation.
   3. Swagger auto-generation with *swaggo* and API client code generation with *open-api*.
   4. CORS and file(multi-part) handling, with zipped response using *gzip*.
   5. Custom errors with formatted string with *xerrors*.
2. **restaurant-service** - Simple RESTful service with CRUD implementation and the following features.
   1. Serving HTTP requests using the Gin framework.
   2. Hot reloading using *air*.
   3. Custom structured logging using *zap*.
   4. Uses *gorm* as the ORM tool for data access.
   5. Uses clean architecture i.e. controllers, services, and data access layers.
   6. Custom recovery and timeout handling middleware.
   7. Uses *viper* for dynamically loading YAML configs based on profiles i.e. dev, QA, and prod.
   8. Unit and integration testing of different layers with *testify* and mock generation.

## Roadmap

* [X] Simple microservice with RESTful APIs and file handling capability.
* [X] Microservice with CRUD functionality using a framework and ORM tool.
* [ ] Implement patterns like service discovery, API gateway, asynchronous messaging, and circuit breaker.
* [ ] Inter-service communication using REST, GraphQL, and GRPC.

## References

1. [https://microservices.io](https://microservices.io/)
