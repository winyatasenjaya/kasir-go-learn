package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

// ReportRepository - repository untuk laporan
type ReportRepository struct {
	db *sql.DB
}

// NewReportRepository - membuat instance baru ReportRepository
func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetTotalRevenue - menghitung total revenue dalam rentang tanggal
func (repo *ReportRepository) GetTotalRevenue(startDate, endDate time.Time) (int, error) {
	query := `
		SELECT COALESCE(SUM(total_amount), 0) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
	`
	var totalRevenue int
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&totalRevenue)
	if err != nil {
		return 0, err
	}
	return totalRevenue, nil
}

// GetTotalTransactions - menghitung total transaksi dalam rentang tanggal
func (repo *ReportRepository) GetTotalTransactions(startDate, endDate time.Time) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
	`
	var totalTransactions int
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&totalTransactions)
	if err != nil {
		return 0, err
	}
	return totalTransactions, nil
}

// GetBestSellingProduct - mendapatkan produk terlaris dalam rentang tanggal
func (repo *ReportRepository) GetBestSellingProduct(startDate, endDate time.Time) (*models.ProdukTerlaris, error) {
	query := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM products p
		INNER JOIN transaction_details td ON p.id = td.product_id
		INNER JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	
	var product models.ProdukTerlaris
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&product.Nama, &product.QtyTerjual)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}
