package main

import (
	"fmt"
	"kasir-api-go/config"
	"kasir-api-go/database"
	"kasir-api-go/handlers"
	"kasir-api-go/repositories"
	"kasir-api-go/services"
	"net/http"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Setup Database
	database.Connect(cfg)
	database.Migrate()

	// 3. Initialize Repositories
	categoryRepo := repositories.NewCategoryRepository(database.DB)
	productRepo := repositories.NewProductRepository(database.DB)

	// 4. Initialize Services
	categoryService := services.NewCategoryService(categoryRepo)
	productService := services.NewProductService(productRepo)

	// 5. Initialize Handlers
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	productHandler := handlers.NewProductHandler(productService)

	// 6. Setup Routes
	// Helper function to handle method routing
	handleFunc := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		http.HandleFunc(pattern, handler)
	}

	// Product Routes
	handleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAll(w, r)
		case http.MethodPost:
			productHandler.Create(w, r)
		default:
			handlers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		}
	})

	handleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetByID(w, r)
		case http.MethodPut:
			productHandler.Update(w, r)
		case http.MethodDelete:
			productHandler.Delete(w, r)
		default:
			handlers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		}
	})

	// Category Routes
	handleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetAll(w, r)
		case http.MethodPost:
			categoryHandler.Create(w, r)
		default:
			handlers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		}
	})

	handleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetByID(w, r)
		case http.MethodPut:
			categoryHandler.Update(w, r)
		case http.MethodDelete:
			categoryHandler.Delete(w, r)
		default:
			handlers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		}
	})

	// Status Route
	handleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.SuccessResponse(w, http.StatusOK, "API Running", map[string]string{
			"status": "OK",
		})
	})

	// 7. Start Server
	fmt.Printf("Server running on port %s in %s mode\n", cfg.ServerPort, cfg.AppEnv)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}
