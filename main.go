package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Product struct represents a product
type Product struct {
	ID    int     `json:"id"`
	Nama  string  `json:"nama"`
	Harga float64 `json:"harga"`
	Stok  int     `json:"stok"`
}

// Category struct represents a product category
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// sample products data
var product = []Product{
	{ID: 1, Nama: "Laptop", Harga: 15000000, Stok: 10},
	{ID: 2, Nama: "Smartphone", Harga: 5000000, Stok: 25},
	{ID: 3, Nama: "Tablet", Harga: 7000000, Stok: 15},
}

// sample categories data
var categories = []Category{
	{ID: 1, Name: "Electronics", Description: "Devices and gadgets"},
	{ID: 2, Name: "Home Appliances", Description: "Appliances for home use"},
	{ID: 3, Name: "Books", Description: "Various genres of books"},
}

// get product by id
func getProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// find product by id
	for _, p := range product {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// update product by id
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// decode request body
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// find product by id
	for i := range product {
		if product[i].ID == id {
			updatedProduct.ID = id
			product[i] = updatedProduct
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// delete product by id
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// find product by id
	for i, p := range product {
		if p.ID == id {
			// delete product from slice
			product = append(product[:i], product[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(
				map[string]string{"message": "Product deleted successfully"},
			)
			return
		}
	}

	// delete product from slice
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Endpoint to get and add categories by id
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// find category by id
	for _, c := range categories {
		if c.ID == id {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// Endpoint to update categories by id
func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// decode request body
	var updatedCategory Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// find category by id
	for i := range categories {
		if categories[i].ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories[i])
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// Endpoint to delete categories by id
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")

	// convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// find category by id
	for i, c := range categories {
		if c.ID == id {
			// delete category from slice
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(
				map[string]string{"message": "Category deleted successfully"},
			)
			return
		}
	}

	// delete category from slice
	http.Error(w, "Category not found", http.StatusNotFound)
}

func main() {
	// GET localhost:8888/api/product{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	// Endpoint to get and add products
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		} else if r.Method == "POST" {
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			newProduct.ID = len(product) + 1
			product = append(product, newProduct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		}
	})

	// Endpoint to get and add categories
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
			return
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(categories)
		}
	})

	// Endpoint to get and add categories by id
	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Service is healthy",
		})
	})

	fmt.Println("Server running at localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Failed Server Running")
	}
}
