# crm-service

A simple CRM (Customer Relationship Management) service written in Go.  
This project provides a RESTful API for managing customer data, including creating, retrieving, updating, and listing customers.

---

## Features

- RESTful API for customer management
- List all customers or retrieve by ID
- Update customer details
- Stream customer lists for large datasets
- Modular, testable code structure
- Uses [chi](https://github.com/go-chi/chi) for routing

---

## Getting Started

### Prerequisites

- Go 1.24 or newer

### Running the Server

To run the server locally:

```sh
go run cmd/main.go
```

The server will start on `http://localhost:8888` by default.
The Swagger UI describing the entire API can be accessed through the same URL or `http://localhost:8888/swagger-ui`

---

## API Endpoints

| Method | Endpoint                | Description                |
|--------|-------------------------|----------------------------|
| GET    | `/v1/customers`         | List all customers         |
| GET    | `/v1/customers/{id}`    | Get customer by ID         |
| PATCH  | `/v1/customers/{id}`    | Update customer by ID      |
| DELETE | `/v1/customers/{id}`    | Delete customer by ID      |
| POST   | `/v1/customers`         | Create a customer          |

### Example: Get All Customers

```sh
curl -X 'GET' \
  'http://localhost:8888/v1/customers' \
  -H 'accept: application/json'
```

### Example: Get Customer by ID

```sh
curl -X 'GET' \
  'http://localhost:8888/v1/customers/0197ab80-571d-727c-a55e-653f85a17854' \
  -H 'accept: application/json'
```

### Example: Create Customer

```sh
curl -X 'POST' \
  'http://localhost:8888/v1/customers' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "firstName": "John",
  "lastName": "Doe",
  "role": "student",
  "email": "jdoe@gmail.com"
}'
```


### Example: Update Customer

```sh
curl -X 'PATCH' \
  'http://localhost:8888/v1/customers/0197ab76-13ac-7a77-b1f5-c98a07560281' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "contacted": true
}'
```

### Example: Delete Customer

```sh
curl -X 'DELETE' 'http://localhost:8888/v1/customers/0197ab82-d19c-746c-8cf0-9eb5c8d8b110'
```

---

## Project Structure

```
cmd/             # Application entrypoint
internal/        # Application core logic
  models/        # Domain models and DTOs
  server/        # HTTP handlers and routing
  store/         # Storage interfaces and implementations
  utils/         # Internal helpers
```

---

## Development

- Instal [air](https://github.com/air-verse/air): `go install github.com/air-verse/air@latest`
- Start the server with `air`
- Code is formatted with `gofmt`.
- Lint with `golangci-lint run`.

---

## License

MIT License

---

