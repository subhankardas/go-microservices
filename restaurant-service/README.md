# Restaurant Service - Go Microservice using Gin Framework + GORM  

This is a simple Go-based microservice with *RESTful CRUD APIs* implementation using the Gin framework for handling HTTP requests and the GORM library as the ORM tool for CRUD operations on the data. This project uses controller > service > data access layer architecture.

## Prerequisites  

1. Install and setup *Air* for enabling hot reloading

    ```bash
     make install-air
    ```

2. Install and setup *Mockgen* for mocking unit tests

    ```bash
     make install-mockgen
    ```

### Commands

The makefile consists of all necessary commands given below.

1. Start docker containers for *Postgres DB*, move to the *docker* directory, and run the below command.

   ```bash
   docker compose up pgadmin postgres
   ```

2. Run service locally with hot reloading using *Air*

    ```bash
    make run-service-<profile-name>
    ```

3. Run integration tests (service should be *running*)

    ```bash
    make run-int-tests
    ```

4. Run unit tests

    ```bash
    make run-unit-tests
    ```

#### Local Development Setup

1. Connect to the PostgresDB server
    Login at [http://localhost:5050](http://localhost:5050/) using the admin credentials for pgAdmin > *Add New Server*

    ```bash
    Name: pg_server (any name will work)
    Hostname: host.docker.internal
    Port: 5432
    Username: pgadmin_user
    Password: pgadmin@123
    ```

#### Features

The basic functionalities include performing CRUD operations of a restaurant. The curl commands to test them are given below.

1. Fetch all menu - **GET**

```curl
curl -X GET \
  'localhost:8080/api/menu'
```

2. Add new menu - **POST**
  
```curl
curl -X POST \
  'localhost:8080/api/menu' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "title": "Breakfast Menu",
  "items": [
    {
      "name": "Sandwich",
      "price": 1.2,
      "description": " text!"
    }
  ]
}'
```

3. Update existing menu - **PUT**

```curl
curl -X PUT \
  'localhost:8080/api/menu/aaedb03f9c584e83922d0d269a03f784' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "title": "Dinner Menu",
  "items": [
    {
      "id": 24,
      "name": "Rice",
      "price": 5.6,
      "description": "new text!"
    }
  ]
}'
```

4. Delete menu by ID - **DELETE**  

```curl
curl -X DELETE \
  'localhost:8080/api/menu/aaedb03f9c584e83922d0d269a03f784'
```

### References

1. [Learn Go by building a REST API](https://learninggolang.com/)
2. [Build a RESTful API using Golang and Gin](https://www.twilio.com/blog/build-restful-api-using-golang-and-gin)
