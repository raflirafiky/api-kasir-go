package handlers

import (
	"encoding/json"
	"kasir-api-go/models"
	"kasir-api-go/services"
	"net/http"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

type CreateTransactionRequest struct {
	Details []TransactionDetailRequest `json:"details"`
}

type TransactionDetailRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validasi request
	if len(req.Details) == 0 {
		ErrorResponse(w, http.StatusBadRequest, "Validation error", "Details tidak boleh kosong")
		return
	}

	for _, detail := range req.Details {
		if detail.ProductID == 0 {
			ErrorResponse(w, http.StatusBadRequest, "Validation error", "Product ID wajib diisi")
			return
		}
		if detail.Quantity <= 0 {
			ErrorResponse(w, http.StatusBadRequest, "Validation error", "Quantity harus lebih dari 0")
			return
		}
	}

	// Convert request ke model
	details := make([]models.TransactionDetail, len(req.Details))
	for i, d := range req.Details {
		details[i] = models.TransactionDetail{
			ProductID: d.ProductID,
			Quantity:  d.Quantity,
		}
	}

	transaction := &models.Transaction{}
	result, err := h.service.Create(transaction, details)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Failed to create transaction", err.Error())
		return
	}

	SuccessResponse(w, http.StatusCreated, "Transaction created successfully", result)
}
