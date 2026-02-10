package services

import (
	"kasir-api-go/models"
	"kasir-api-go/repositories"
)

type TransactionService interface {
	Create(transaction *models.Transaction, details []models.TransactionDetail) (*models.Transaction, error)
	GetTodaySummary() (*repositories.SalesSummary, error)
	GetSummaryByDateRange(startDate, endDate string) (*repositories.SalesSummary, error)
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) Create(transaction *models.Transaction, details []models.TransactionDetail) (*models.Transaction, error) {
	return s.repo.Create(transaction, details)
}

func (s *transactionService) GetTodaySummary() (*repositories.SalesSummary, error) {
	return s.repo.GetTodaySummary()
}

func (s *transactionService) GetSummaryByDateRange(startDate, endDate string) (*repositories.SalesSummary, error) {
	return s.repo.GetSummaryByDateRange(startDate, endDate)
}
