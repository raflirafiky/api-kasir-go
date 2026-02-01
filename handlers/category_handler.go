package handlers

import (
	"encoding/json"
	"kasir-api-go/models"
	"kasir-api-go/services"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	categories, err := h.service.GetAll()
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data kategori", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil mengambil data kategori", categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Category ID", err.Error())
		return
	}

	category, err := h.service.GetByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Kategori tidak ditemukan", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil mengambil data kategori", category)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.service.Create(&category); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal membuat kategori", err.Error())
		return
	}

	SuccessResponse(w, http.StatusCreated, "Berhasil membuat kategori", category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Category ID", err.Error())
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.service.Update(uint(id), &category); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal update kategori", err.Error())
		return
	}

	// Fetch updated data to return
	updatedCategory, _ := h.service.GetByID(uint(id))
	SuccessResponse(w, http.StatusOK, "Berhasil update kategori", updatedCategory)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid Category ID", err.Error())
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gagal menghapus kategori", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Berhasil menghapus kategori", nil)
}
