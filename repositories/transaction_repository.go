package repositories

import (
	"errors"
	"fmt"
	"kasir-api-go/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction, details []models.TransactionDetail) (*models.Transaction, error)
	GetTodaySummary() (*SalesSummary, error)
	GetSummaryByDateRange(startDate, endDate string) (*SalesSummary, error)
}

type SalesSummary struct {
	TotalRevenue   int                `json:"total_revenue"`
	TotalTransaksi int                `json:"total_transaksi"`
	ProdukTerlaris *ProdukTerlaris    `json:"produk_terlaris"`
}

type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *models.Transaction, details []models.TransactionDetail) (*models.Transaction, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Validasi stock dan hitung total
		total := 0
		for i := range details {
			// Ambil product dengan lock untuk avoid race condition
			var product models.Product
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, details[i].ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("product dengan ID %d tidak ditemukan", details[i].ProductID)
				}
				return err
			}

			// Validasi stock
			if product.Stock < details[i].Quantity {
				return fmt.Errorf("stock untuk product '%s' tidak tersedia (tersisa: %d, diminta: %d)", 
					product.Name, product.Stock, details[i].Quantity)
			}

			// Hitung subtotal
			details[i].Subtotal = product.Price * details[i].Quantity
			total += details[i].Subtotal

			// Kurangi stock
			product.Stock -= details[i].Quantity
			if err := tx.Save(&product).Error; err != nil {
				return err
			}
		}

		// 2. Set total transaksi
		transaction.Total = total

		// 3. Insert transaction
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// 4. Insert transaction details
		for i := range details {
			details[i].TransactionID = transaction.ID
			if err := tx.Create(&details[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Load details dengan product untuk response
	if err := r.db.Preload("Details.Product").First(transaction, transaction.ID).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetTodaySummary() (*SalesSummary, error) {
	var summary SalesSummary

	// Hitung total revenue dan total transaksi hari ini
	type Result struct {
		TotalRevenue   int
		TotalTransaksi int
	}
	var result Result
	
	today := time.Now().Format("2006-01-02")
	
	err := r.db.Model(&models.Transaction{}).
		Select("COALESCE(SUM(total), 0) as total_revenue, COUNT(*) as total_transaksi").
		Where("DATE(created_at) = ?", today).
		Scan(&result).Error
	
	if err != nil {
		return nil, err
	}
	
	summary.TotalRevenue = result.TotalRevenue
	summary.TotalTransaksi = result.TotalTransaksi

	// Cari produk terlaris hari ini
	var produkTerlaris ProdukTerlaris
	err = r.db.Model(&models.TransactionDetail{}).
		Select("products.name as nama, SUM(transaction_details.quantity) as qty_terjual").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Joins("JOIN transactions ON transactions.id = transaction_details.transaction_id").
		Where("DATE(transactions.created_at) = ?", today).
		Group("products.id, products.name").
		Order("qty_terjual DESC").
		Limit(1).
		Scan(&produkTerlaris).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if produkTerlaris.Nama != "" {
		summary.ProdukTerlaris = &produkTerlaris
	}

	return &summary, nil
}

func (r *transactionRepository) GetSummaryByDateRange(startDate, endDate string) (*SalesSummary, error) {
	var summary SalesSummary

	type Result struct {
		TotalRevenue   int
		TotalTransaksi int
	}
	var result Result
	
	err := r.db.Model(&models.Transaction{}).
		Select("COALESCE(SUM(total), 0) as total_revenue, COUNT(*) as total_transaksi").
		Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate).
		Scan(&result).Error
	
	if err != nil {
		return nil, err
	}
	
	summary.TotalRevenue = result.TotalRevenue
	summary.TotalTransaksi = result.TotalTransaksi

	// Cari produk terlaris di range tanggal
	var produkTerlaris ProdukTerlaris
	err = r.db.Model(&models.TransactionDetail{}).
		Select("products.name as nama, SUM(transaction_details.quantity) as qty_terjual").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Joins("JOIN transactions ON transactions.id = transaction_details.transaction_id").
		Where("DATE(transactions.created_at) BETWEEN ? AND ?", startDate, endDate).
		Group("products.id, products.name").
		Order("qty_terjual DESC").
		Limit(1).
		Scan(&produkTerlaris).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if produkTerlaris.Nama != "" {
		summary.ProdukTerlaris = &produkTerlaris
	}

	return &summary, nil
}
