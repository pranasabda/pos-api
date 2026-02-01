package handlers

import (
	"KASIR-API/models"
	"KASIR-API/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(s services.CategoryService) *CategoryHandler {
	return &CategoryHandler{s}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, _ := h.service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)
	res, _ := h.service.Store(cat)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Detail Categories
func (h *CategoryHandler) GetCategoryDetail(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL path /api/categories/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, _ := strconv.Atoi(idStr)

	category, err := h.service.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// UpdateCategory menangani request PUT /api/categories/{id}
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, _ := strconv.Atoi(idStr)

	var p models.Category
	json.NewDecoder(r.Body).Decode(&p)

	res, err := h.service.Update(id, p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Kategori tidak ditemukan"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// DeleteCategories menangani request DELETE /api/categories/{id}
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, _ := strconv.Atoi(idStr)

	err := h.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Gagal menghapus"})
		return
	}
	w.WriteHeader(http.StatusNoContent) // Sukses hapus (204 No Content)
}
