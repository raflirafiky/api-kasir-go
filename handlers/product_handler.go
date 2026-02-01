package handlers

import (
	"encoding/json"
	"kasir-api-go/models"
	"kasir-api-go/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	products, err := h.service.GetAll()
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data produk", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil mengambil data produk", products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Product ID", err.Error())
		return
	}

	product, err := h.service.GetByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Produk tidak ditemukan", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil mengambil data produk", product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.service.Create(&product); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal membuat produk", err.Error())
		return
	}

	// Fetch created data with category to return
	createdProduct, _ := h.service.GetByID(product.ID)
	SuccessResponse(w, http.StatusCreated, "Berhasil membuat produk", createdProduct)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Product ID", err.Error())
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.service.Update(uint(id), &product); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal update produk", err.Error())
		return
	}

	// Fetch updated data to return
	updatedProduct, _ := h.service.GetByID(uint(id))
	SuccessResponse(w, http.StatusOK, "Berhasil update produk", updatedProduct)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Product ID", err.Error())
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal menghapus produk", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil menghapus produk", nil)
}
