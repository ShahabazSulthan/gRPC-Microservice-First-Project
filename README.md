
# First gRPC Microservices Project

This project demonstrates a simple microservices architecture using gRPC in Go. The system comprises four main services: **API Gateway**, **Auth Service**, **Order Service**, and **Product Service**.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Services](#services)
  - [API Gateway](#api-gateway)
  - [Auth Service](#auth-service)
  - [Order Service](#order-service)
  - [Product Service](#product-service)
- [Environment Variables](#environment-variables)
- [Getting Started](#getting-started)
- [Build and Run](#build-and-run)
- [gRPC Protobufs](#grpc-protobufs)
- [Contributing](#contributing)
- [License](#license)

## Overview

This project is designed to help you understand the basics of building microservices using gRPC with Go. Each service is responsible for a specific domain and communicates with others via gRPC. 

## Project Structure

```bash
.
+ ├── Api-Gateway
+ ├── Auth-Service
+ ├── Order-Service
+ ├── Product-Service
+ └── .gitignore
```

### Common Structure for Each Service

- **cmd/**: Entry point for the service.
- **pkg/**: Contains the main logic, configurations, gRPC client and server implementations.
  - **pb/**: Contains the protobuf-generated files.
  - **routes/**: Contains the route handlers for the API Gateway.
  - **config/**: Manages environment variables and configurations.
  - **models/**: Contains the data models.
  - **services/**: Contains the business logic.
  - **db/**: Manages database connections.
- **Makefile**: Automates the build process.
- **go.mod/go.sum**: Go modules for dependency management.

## Services

### API Gateway

The API Gateway routes incoming requests to the appropriate microservices. It acts as a reverse proxy and is the entry point for the application.

**Key Files:**
- `main.go`: Initializes and runs the API Gateway.
- `routes/`: Contains the routing logic for user login, signup, and product/order actions.

### Auth Service

Handles authentication and authorization, including user login and signup.

**Key Files:**
- `main.go`: Initializes and runs the Auth service.
- `service.go`: Contains the authentication logic.
- `models/model.go`: Defines the user data model.
- `pb/`: Protobuf files for gRPC communication.

### Order Service

Manages orders within the system, including creating and retrieving orders.

**Key Files:**
- `main.go`: Initializes and runs the Order service.
- `order.go`: Contains the order management logic.
- `models/model.go`: Defines the order data model.
- `pb/`: Protobuf files for gRPC communication.

### Product Service

Manages product information, including creating and finding products.

**Key Files:**
- `main.go`: Initializes and runs the Product service.
- `server.go`: Contains the product management logic.
- `models/product.go`: Defines the product data model.
- `pb/`: Protobuf files for gRPC communication.

## Environment Variables

Each service has its own `.env` file located in `pkg/config/envs/` to manage environment-specific configurations.

### Example `.env` file for Development (`dev.env`):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
JWT_SECRET=your_jwt_secret
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (v1.16+)
- [Protocol Buffers](https://grpc.io/docs/protoc-installation/) (v3.0.0+)
- [Make](https://www.gnu.org/software/make/)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name
   ```

2. Install dependencies:
   ```bash
   make install
   ```

## Build and Run

### Using Makefile

To build and run each service, navigate to the service directory and use the Makefile:

```bash
# Example for API Gateway
cd Api-Gateway
make run
```

### Manually

Alternatively, you can manually build and run each service:

```bash
# Build the service
go build -o bin/service_name cmd/main.go

# Run the service
./bin/service_name
```

## gRPC Protobufs

The project uses Protocol Buffers (Protobuf) for defining the gRPC services and messages. The `.proto` files are located in the `pb/` directory of each service.

### Generate gRPC Code

To regenerate the gRPC code from `.proto` files:

```bash
protoc --go_out=. --go-grpc_out=. pb/*.proto
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss the changes you want to make.

