### Restaurant Service - Go Microservice using Gin Framework + GORM  

This is simple Go based microservice with *RESTful CRUD APIs* implementation using the Gin framework for handling HTTP request and GORM library as the ORM tool for CRUD operations on the data.

#### Prerequisites  
1. Install and setup *Air* for enabling hot reloading
    ```
    go install github.com/cosmtrek/air@latest
    air init
    ```

#### Commands
The makefile consists all important commands given below.

1. Run service locally with hot reloading using Air
    ```
    make run-service
    ```
2. Run unit tests
    ```
    make run-tests
    ```

#### Local Development Setup
1. Connect to PostgresDB server
    Login at [http://localhost:5050](http://localhost:5050/) using the admin credentials for pgAdmin > *Add New Server*
    ```
    Name: pg_server (any name will work)
    Hostname: host.docker.internal
    Port: 5432
    Username: pgadmin_user
    Password: pgadmin@123
    ```

#### Features
The basic functionalities include performing CRUD operations of a restaurant. The curl commands to test them are given below.

**1. GET - Fetch all menu**
```
curl -X GET \
  'localhost:8080/api/menu'
```

### References
1. [Learn Go by building a REST API](https://learninggolang.com/)
2. [Build a RESTful API using Golang and Gin](https://www.twilio.com/blog/build-restful-api-using-golang-and-gin)