package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ProdukKategori mendefinisikan struktur data untuk kategori barang
type ProdukKategori struct {
	ID        int    `json:"id"`
	Nama      string `json:"name"`
	Deskripsi string `json:"description"`
}

// Data sementara disimpan dalam memory (slice)
var daftarKategori = []ProdukKategori{
	{ID: 1, Nama: "Elektronik", Deskripsi: "Gadget dan perangkat elektronik"},
	{ID: 2, Nama: "Fashion", Deskripsi: "Pakaian dan aksesoris"},
}

// Helper untuk mengirim respon JSON agar kode lebih bersih
func responJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Handler untuk /api/categories
func kategoriHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		responJSON(w, http.StatusOK, daftarKategori)

	case http.MethodPost:
		var item ProdukKategori
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			responJSON(w, http.StatusBadRequest, map[string]string{"error": "Format data salah"})
			return
		}
		item.ID = len(daftarKategori) + 1
		daftarKategori = append(daftarKategori, item)
		responJSON(w, http.StatusCreated, item)

	default:
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

// Handler untuk /api/categories/{id}
func kategoriDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 {
		http.Error(w, "ID dibutuhkan", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		responJSON(w, http.StatusBadRequest, map[string]string{"error": "ID harus berupa angka"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		for _, k := range daftarKategori {
			if k.ID == id {
				responJSON(w, http.StatusOK, k)
				return
			}
		}
		responJSON(w, http.StatusNotFound, map[string]string{"error": "Kategori tidak ditemukan"})

	case http.MethodPut:
		var updateData ProdukKategori
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			responJSON(w, http.StatusBadRequest, map[string]string{"error": "Input tidak valid"})
			return
		}

		for i, k := range daftarKategori {
			if k.ID == id {
				updateData.ID = id
				daftarKategori[i] = updateData
				responJSON(w, http.StatusOK, updateData)
				return
			}
		}
		responJSON(w, http.StatusNotFound, map[string]string{"error": "Gagal update, data tidak ada"})

	case http.MethodDelete:
		for i, k := range daftarKategori {
			if k.ID == id {
				daftarKategori = append(daftarKategori[:i], daftarKategori[i+1:]...)
				responJSON(w, http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
				return
			}
		}
		responJSON(w, http.StatusNotFound, map[string]string{"error": "Data tidak ditemukan"})

	default:
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Endpoint Utama
	http.HandleFunc("/api/categories", kategoriHandler)        // List & Create
	http.HandleFunc("/api/categories/", kategoriDetailHandler) // Detail, Update, Delete

	// Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		responJSON(w, http.StatusOK, map[string]string{
			"status":  "UP",
			"service": "Kasir API",
		})
	})

	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Gagal menjalankan server: %s\n", err)
	}
}
