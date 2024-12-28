# E-Commerce API

This project is a RESTful API built for managing an e-commerce platform. It allows users to register, log in, manage orders, and for admins to manage products. The API is built using Go with the Gin framework and PostgreSQL for the database.

---

## Features

### User Endpoints:
- **Register**: Create a new user account.
- **Login**: Authenticate and receive a JWT token.

### Order Endpoints:
- **Place Order**: Users can place orders with a list of products.
- **Cancel Order**: Users can cancel an order if it is in a pending state.
- **Get User Orders**: Retrieve all orders placed by the authenticated user.

### Admin Endpoints:
- **Create Product**: Admins can add new products.
- **Update Product**: Admins can update product details.
- **Delete Product**: Admins can remove products.
- **Get All Products**: Retrieve all products.

---

## Technologies Used

- **Go (Gin Framework)**: For building the RESTful API.
- **GORM**: For ORM with PostgreSQL.
- **PostgreSQL**: As the database.
- **Swagger**: For API documentation.
- **Clever Cloud**: For hosting the PostgreSQL database.

---

## Setup and Installation

### Prerequisites
- Go installed on your system.
- PostgreSQL database.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/echewisi/ecommerce_api.git
   cd ecommerce_api
   ```

2. Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<database>?sslmode=require(i used from clever cloud here)
   JWT_SECRET=your_jwt_secret
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

5. Access the API at `http://localhost:8080/`.

6. Access the Swagger UI at `http://localhost:8080/swagger/index.html`.

---

## API Documentation

API documentation is available via Swagger. Visit `http://localhost:8080/swagger/index.html` to explore and test the API.

---

## Directory Structure

```
ecommerce_api/
├── cmd/
│   └── main.go           # Application entry point
├── config/
│   └── config.go         # Configuration handling
├── controllers/
│   ├── auth_controller.go
│   ├── product_controller.go
│   └── order_controller.go
├── database/
│   └── db.go             # Database connection
├── docs/
│   └── swagger.json      # Swagger documentation
├── middleware/
│   └── auth_middleware.go
├── models/
│   ├── user.go
│   ├── product.go
│   └── order.go
├── repositories/
│   ├── user_repository.go
│   ├── product_repository.go
│   └── order_repository.go
├── routes/
│   └── routes.go         # Route definitions
├── services/
│   ├── auth_service.go
│   ├── product_service.go
│   └── order_service.go
├── go.mod                # Go module file
└── go.sum
```

---

## Key Endpoints

### Public Routes
- **POST** `/register`: Register a new user.
- **POST** `/login`: Log in to receive a JWT.

### User Routes (Authentication Required)
- **POST** `/orders`: Place an order.
- **PUT** `/orders/:id/cancel`: Cancel an order.
- **GET** `/orders`: Get all user orders.

### Admin Routes (Admin Authentication Required)
- **POST** `/admin/products`: Create a new product.
- **PUT** `/admin/products/:id`: Update a product.
- **DELETE** `/admin/products/:id`: Delete a product.
- **GET** `/admin/products`: Retrieve all products.
- **PUT** `/admin/orders/:id`: Update an order status.

---

## Notes
- Ensure the PostgreSQL database is running and accessible.
- Replace placeholder values in `.env` with your actual Clever Cloud credentials.
- Use `gin.SetMode(gin.ReleaseMode)` for production environments.

---

## License
This project is licensed under the MIT License.
