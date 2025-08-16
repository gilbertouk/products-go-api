## Database Initialization

Before running the application, you need to create the required table in your PostgreSQL database. Use the SQL script located at `script/db.sql`:

```bash
psql -h <host> -U <user> -d <dbname> -f script/db.sql
```

Replace `<host>`, `<user>`, and `<dbname>` with your PostgreSQL connection details as configured in `db/connection.go`.

This script will create the `product` table and insert an example record.

# Products Go API

This project was developed as a learning exercise for the Go programming language and the Gin web framework. It is a simple REST API that provides CRUD operations for products.

## Features
- RESTful API built with Go and Gin
- CRUD operations for products (Create, Read, Update, Delete)
- Example of project structure for Go web applications

## Project Structure
```
cmd/            # Application entry point
controller/     # HTTP handlers/controllers
model/          # Data models
repository/     # Data access layer
usecase/        # Business logic
script/         # Database scripts
endpoint.http   # Example HTTP requests
Dockerfile      # Docker configuration
```


## Database

This project uses **PostgreSQL** as its database. You must configure your database connection settings in the file `db/connection.go` before running the application. Update the following variables as needed:

```
const (
   host     = "go_db" // Change this to your database host (e.g., localhost or go_db if using Docker)
   port     = "5432"
   user     = "postgres"
   password = "postgres"
   dbname   = "postgres"
)
```

Make sure your PostgreSQL instance is running and accessible with the credentials above, or modify them to match your environment.

## Getting Started

### Prerequisites
- Go 1.24.5 or higher
- Docker (optional, for containerization)

### Running Locally
1. Clone the repository:
   ```bash
   git clone https://github.com/gilbertouk/products-go-api.git
   cd products-go-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Using Docker
1. Build and run with Docker Compose:
   ```bash
   docker compose up --build
   ```

## API Endpoints
- `GET /products` - List all products
- `GET /products/:productId` - Get a product by ID
- `POST /products` - Create a new product
- `PUT /products/:productId` - Update a product
- `DELETE /products/:productId` - Delete a product

## License
This project is for educational purposes only.
