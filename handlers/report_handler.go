package handlers

import (
	"kasir-api-go/repositories"
	"kasir-api-go/services"
	"net/http"
	"regexp"
)

type ReportHandler struct {
	service services.TransactionService
}

func NewReportHandler(service services.TransactionService) *ReportHandler {
	return &ReportHandler{service}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	// Cek query params untuk date range
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var summary *repositories.SalesSummary
	var err error

	// Jika ada date range, gunakan GetSummaryByDateRange
	if startDate != "" && endDate != "" {
		// Validasi format tanggal (YYYY-MM-DD)
		if !isValidDateFormat(startDate) || !isValidDateFormat(endDate) {
			ErrorResponse(w, http.StatusBadRequest, "Invalid date format", "Format tanggal harus YYYY-MM-DD")
			return
		}
		summary, err = h.service.GetSummaryByDateRange(startDate, endDate)
	} else {
		// Default: hari ini
		summary, err = h.service.GetTodaySummary()
	}

	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Failed to get report", err.Error())
		return
	}

	SuccessResponse(w, http.StatusOK, "Report retrieved successfully", summary)
}

func isValidDateFormat(date string) bool {
	// Regex untuk format YYYY-MM-DD
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return regex.MatchString(date)
}
