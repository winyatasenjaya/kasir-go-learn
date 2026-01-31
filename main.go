package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"strings"
	"encoding/json"

	"kasir-api/repositories"
	"kasir-api/config"
	"kasir-api/handlers"
	"kasir-api/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
	DBConnectionString string `mapstructure:"db_connection_string"`
}

func main() {
	//load configuration
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	configEnv := Config{
		Port:  viper.GetString("port"),
		DBConnectionString: viper.GetString("db_connection_string"),
	}

	// Initialize database
	db, err := config.InitDB(configEnv.DBConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	// productRepo := memory.NewProductRepository()
	// categoryRepo := memory.NewCategoryRepository()

	// Initialize use cases
	// productUseCase := usecase.NewProductUseCase(productRepo)
	// categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	// Initialize handler
	// handler := httpDelivery.NewHandler(productUseCase, categoryUseCase)

	// Setup routes
	// router := httpDelivery.SetupRoutes(handler)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup routes
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Start server
	fmt.Println("Server running at localhost:" + configEnv.Port)

	err = http.ListenAndServe(":" + configEnv.Port, nil)
	if err != nil {
		fmt.Println("Failed Server Running:", err)
	}
}
