# Products Service - Go Microservice using Gorilla Mux

This is a simple Go-based microservice implemented following tutorials by *@nicholasjackson*. Here we have implemented a basic RESTful API using the *Gorilla MUX* library.

## Commands

The makefile consists of all important commands given below.

1. Run service locally

```bash
make run-service
```

2. Run unit tests

```bash
make run-tests
```

3. Install and auto-generate swagger spec

```bash
make install-swagger
make swagger-generate
```

4. Install swagger and auto generate API client code

```bash
make install-go-swagger
make client-generate
```

### Features

The basic functionalities include performing crud operations on a few coffee products. The curl commands to test them are given below.

1. **GET** - Fetch products list

```curl
curl -X GET \
  'http://localhost:8080/api/products'
```

2. **POST** - Add a new product

```curl
curl -X POST \
  'http://localhost:8080/api/products' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "Cappuccino",
  "description": "Semi-frothy milky coffee.",
  "price": 6.45,
  "sku": "COFF678"
}'
```

3. **PUT** - Update an existing product

```curl
curl -X PUT \
  'http://localhost:8080/api/products/3' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "test12",
  "description": "Semi-frothy milky coffee.",
  "price": 10,
  "sku": "COFF678"
}'
```

We also added some file-handling functionality to upload and download (using *GZIP*) product images, CURL snippets are given below.

4. **POST** - Upload product images

```curl
curl -X POST \
  'http://localhost:8080/images' \
  --form 'id="1"' \
  --form 'file=@c:\Users\Subhankar.Das\Pictures\Saved Pictures\AVTR.jpg'
```

5. **GET** - Download product images

```curl
curl -X GET \
  'http://localhost:8080/images/1/AVTR.jpg' \
  --header 'Accept-Encoding: gzip'
```

### References

1. [Building Microservices with Go -
Nic Jackson](https://youtu.be/VzBGi_n65iU)
