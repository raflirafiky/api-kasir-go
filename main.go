package main

import (
	"encoding/json"
	"fmt"
	"kasir-api-go/models"
	"net/http"
	"strconv"
	"strings"
)

var product = []models.Product{
	{ID: 1, Name: "Indomie Tekor (Telor Kornet)", Price: 12000, Stock: 10},
	{ID: 2, Name: "Indomie PakLeng (Paket Lengkap)", Price: 20000, Stock: 10},
	{ID: 3, Name: "Es Jeruk Seger (Es Jeger)", Price: 5000, Stock: 40},
}

var categories = []models.Category{
	{ID: 1, Name: "Makanan", Description: "Kategori untuk produk makanan"},
	{ID: 2, Name: "Minuman", Description: "Kategori untuk produk minuman"},
	{ID: 3, Name: "Snack", Description: "Kategori untuk produk snack"},
}

// GET localhost:8080/api/product/{id}
func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/product/{id}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id, ganti sesuai data dari request
	for i := range product {
		if product[i].ID == id {
			updateProduct.ID = id
			product[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range product {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			product = append(product[:i], product[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})

			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// GET localhost:8080/api/categories/{id}
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// PUT localhost:8080/api/categories/{id}
func updateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop kategori, cari id, ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCategory)
			return
		}
	}
	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// DELETE localhost:8080/api/categories/{id}
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	// loop kategori cari ID, dapet index yang mau dihapus
	for i, c := range categories {
		if c.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete kategori",
			})

			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

func main() {
	// GET localhost:8080/api/product/{id}
	// PUT localhost:8080/api/product/{id}
	// DELETE localhost:8080/api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProductByID(w, r)
		} else if r.Method == http.MethodPut {
			updateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			deleteProduct(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "method tidak didukung",
			})
		}
	})

	// GET localhost:8080/api/product
	// POST localhost:8080/api/product
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		} else if r.Method == "POST" {
			// baca data dari request
			var productNew models.Product
			err := json.NewDecoder(r.Body).Decode(&productNew)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// masukkin data ke dalam variable produk
			productNew.ID = len(product) + 1
			product = append(product, productNew)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(productNew)
		}

	})

	// GET localhost:8080/api/categories/{id}
	// PUT localhost:8080/api/categories/{id}
	// DELETE localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategoryByID(w, r)
		} else if r.Method == http.MethodPut {
			updateCategory(w, r)
		} else if r.Method == http.MethodDelete {
			deleteCategory(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "method tidak didukung",
			})
		}
	})

	// GET localhost:8080/api/categories
	// POST localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			// baca data dari request
			var categoryNew models.Category
			err := json.NewDecoder(r.Body).Decode(&categoryNew)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// masukkin data ke dalam variable categories
			categoryNew.ID = len(categories) + 1
			categories = append(categories, categoryNew)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(categoryNew)
		}

	})

	// localhost:8080/api/status
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
