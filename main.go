package main

import (
	"KASIR-API/database"
	"KASIR-API/handlers"
	"KASIR-API/repositories"
	"KASIR-API/services"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Inisialisasi Database
	database.ConnectDB()

	// Inisialisasi Repository, Service, dan Handler untuk Category
	catRepo := repositories.NewCategoryRepository(database.DB)
	catService := services.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catService)

	// Inisialisasi Repository, Service, dan Handler untuk Product
	prodRepo := repositories.NewProductRepository(database.DB)
	prodService := services.NewProductService(prodRepo)
	prodHandler := handlers.NewProductHandler(prodService)

	// --- ROUTING ---

	// Routes Category
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			catHandler.GetCategories(w, r)
		}
		if r.Method == http.MethodPost {
			catHandler.CreateCategory(w, r)
		}
	})
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			catHandler.UpdateCategory(w, r)
		}
		if r.Method == http.MethodDelete {
			catHandler.DeleteCategory(w, r)
		}
		if r.Method == http.MethodGet {
			catHandler.GetCategoryDetail(w, r)
		}
	})

	// Routes Product
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		// Handle /api/products/{id} (DETAIL, UPDATE, DELETE)
		if r.Method == http.MethodGet {
			prodHandler.GetProducts(w, r)
		}
		if r.Method == http.MethodPost {
			prodHandler.CreateProduct(w, r)
		}
		// Dengan switch
		// switch r.Method {
		// case http.MethodGet:
		//     prodHandler.GetProductDetail(w, r)
		// case http.MethodPut:
		//     prodHandler.UpdateProduct(w, r)
		// case http.MethodDelete:
		//     prodHandler.DeleteProduct(w, r)
		// }
	})
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			prodHandler.GetProductDetail(w, r)
		}
		if r.Method == http.MethodPut {
			prodHandler.UpdateProduct(w, r)
		}
		if r.Method == http.MethodDelete {
			prodHandler.DeleteProduct(w, r)
		}
	})

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server POS berjalan di http://localhost:8080")
	http.ListenAndServe(":"+port, nil)
}
