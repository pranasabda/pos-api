package handlers

import (
	"KASIR-API/models"
	"KASIR-API/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{s}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := h.service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL path /api/products/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)

	product, err := h.service.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product

	// Bagian krusial: Decode JSON ke Struct
	// json.NewDecoder(r.Body).Decode(&p) --> code awal sebelum handling error
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	res, err := h.service.Store(p)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// res, _ := h.service.Store(p) --> code awal sebelum handling error
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdateProduct menangani request PUT /api/products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)

	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)

	res, err := h.service.Update(id, p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Produk tidak ditemukan"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// DeleteProduct menangani request DELETE /api/products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)

	err := h.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Gagal menghapus"})
		return
	}
	w.WriteHeader(http.StatusNoContent) // Sukses hapus (204 No Content)
}
