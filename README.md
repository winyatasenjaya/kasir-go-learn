# Kasir API

A lightweight RESTful API built with Go for managing point-of-sale (POS) operations. This API provides endpoints for managing products and categories with in-memory storage.

## Features

- ✅ **Product Management** - Create, read, update, and delete products
- ✅ **Category Management** - Create, read, update, and delete categories
- ✅ **RESTful Architecture** - Clean and intuitive API endpoints
- ✅ **CORS Enabled** - Cross-Origin Resource Sharing support
- ✅ **Health Check** - Monitor service status
- ✅ **In-Memory Storage** - Fast data access (Note: data resets on restart)

## Tech Stack

- **Language**: Go 1.25.4
- **HTTP Framework**: Native `net/http` package
- **Data Format**: JSON

## Getting Started

### Prerequisites

- Go 1.25.4 or higher installed on your system

### Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd kasir-api
```

2. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8888`

### Building the Application

To build a standalone executable:

```bash
go build -o kasir-api.exe main.go
```

Then run:

```bash
./kasir-api.exe
```

## API Documentation

### Health Check

**GET** `/health`

Returns the health status of the service.

**Response:**

```json
{
  "status": "OK",
  "message": "Service is healthy"
}
```

---

### Products

#### Get All Products

**GET** `/api/product`

Returns a list of all products.

**Response:**

```json
[
  {
    "id": 1,
    "nama": "Laptop",
    "harga": 15000000,
    "stok": 10
  }
]
```

#### Get Product by ID

**GET** `/api/product/{id}`

Returns a single product by ID.

**Response:**

```json
{
  "id": 1,
  "nama": "Laptop",
  "harga": 15000000,
  "stok": 10
}
```

#### Create Product

**POST** `/api/product`

Creates a new product.

**Request Body:**

```json
{
  "nama": "Headphones",
  "harga": 2000000,
  "stok": 30
}
```

**Response:**

```json
[
  {
    "id": 4,
    "nama": "Headphones",
    "harga": 2000000,
    "stok": 30
  }
]
```

#### Update Product

**PUT** `/api/product/{id}`

Updates an existing product.

**Request Body:**

```json
{
  "nama": "Gaming Laptop",
  "harga": 25000000,
  "stok": 15
}
```

**Response:**

```json
{
  "id": 1,
  "nama": "Gaming Laptop",
  "harga": 25000000,
  "stok": 15
}
```

#### Delete Product

**DELETE** `/api/product/{id}`

Deletes a product by ID.

**Response:**

```json
{
  "message": "Product deleted successfully"
}
```

---

### Categories

#### Get All Categories

**GET** `/api/category`

Returns a list of all categories.

**Response:**

```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Devices and gadgets"
  }
]
```

#### Get Category by ID

**GET** `/api/category/{id}`

Returns a single category by ID.

**Response:**

```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Devices and gadgets"
}
```

#### Create Category

**POST** `/api/category`

Creates a new category.

**Request Body:**

```json
{
  "name": "Electronics",
  "description": "Devices and gadgets"
}
```

**Response:**

```json
[
  {
    "id": 4,
    "name": "Electronics",
    "description": "Devices and gadgets"
  }
]
```

#### Update Category

**PUT** `/api/category/{id}`

Updates an existing category.

**Request Body:**

```json
{
  "name": "Gaming Electronics",
  "description": "Devices and gadgets for gaming"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Gaming Electronics",
  "description": "Devices and gadgets for gaming"
}
```

#### Delete Category

**DELETE** `/api/category/{id}`

Deletes a category by ID.

**Response:**

```json
{
  "message": "Category deleted successfully"
}
```

## Testing

The project includes a `testAPI.http` file with sample requests for testing all endpoints. You can use tools like:

- **VS Code REST Client** extension
- **Postman**
- **cURL**
- **Thunder Client**

## Project Structure

```
kasir-api/
├── main.go          # Main application file with all handlers
├── go.mod           # Go module definition
├── testAPI.http     # Sample API requests for testing
└── README.md        # Project documentation
```

## Data Models

### Product

```go
type Product struct {
    ID    int     `json:"id"`
    Nama  string  `json:"nama"`  // Product name
    Harga float64 `json:"harga"` // Price
    Stok  int     `json:"stok"`  // Stock quantity
}
```

### Category

```go
type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

## Sample Data

The API comes pre-loaded with sample data:

**Products:**

- Laptop (Rp 15,000,000)
- Smartphone (Rp 5,000,000)
- Tablet (Rp 7,000,000)

**Categories:**

- Electronics
- Home Appliances
- Books

## Notes

⚠️ **Important**: This API uses in-memory storage. All data will be reset when the server restarts. For persistent storage, consider integrating a database like PostgreSQL, MySQL, or MongoDB.

## Future Improvements

- [ ] Add database integration for persistent storage
- [ ] Implement authentication and authorization
- [ ] Add pagination for list endpoints
- [ ] Implement search and filtering
- [ ] Add input validation
- [ ] Add unit and integration tests
- [ ] Implement logging middleware
- [ ] Add API rate limiting

## License

This project is open source and available under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
