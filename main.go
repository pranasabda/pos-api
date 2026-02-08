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

	// Dependency Injection - Transaction & Report
	transRepo := repositories.NewTransactionRepository(database.DB)
	transService := services.NewTransactionService(transRepo, prodRepo)
	transHandler := handlers.NewTransactionHandler(transService)

	reportService := services.NewReportService(transRepo)
	reportHandler := handlers.NewReportHandler(reportService)

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

	// Transactions
	http.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			transHandler.CreateTransaction(w, r)
		}
	})

	// --- TRANSACTIONS & CHECKOUT ---
	// endpoint /api/checkout sesuai catatan, fungsi mirip dengan api/transactions
	http.HandleFunc("/api/checkout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			transHandler.CreateTransaction(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Query dengan endpoint product
	// http.HandleFunc("/api/products/product", prodHandler.SearchProduct)

	// Search
	http.HandleFunc("/api/products/search", prodHandler.SearchProduct)

	// Report Summary Hari Ini (Task 2b)
	http.HandleFunc("/api/report/hari-ini", reportHandler.GetSalesReport)
	// Report Summary Custom Range (Task 2c - Optional Challenge)
	http.HandleFunc("/api/report", reportHandler.GetSalesReport)

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// fmt.Println("Server POS berjalan di http://localhost:8080")
	fmt.Printf("Server POS berjalan di http://localhost:%s\n", port) //update agar port dinamis
	http.ListenAndServe(":"+port, nil)
}
