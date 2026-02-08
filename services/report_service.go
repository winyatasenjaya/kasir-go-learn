package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.ReportResponse, error) {
	// Get today's date range (start of day to start of next day)
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	return s.GetReportByDateRange(startOfDay, endOfDay)
}

// GetReportByDateRange - mendapatkan laporan dalam range tanggal tertentu
func (s *ReportService) GetReportByDateRange(startDate, endDate time.Time) (*models.ReportResponse, error) {
	// Get total revenue
	totalRevenue, err := s.repo.GetTotalRevenue(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get total transactions
	totalTransactions, err := s.repo.GetTotalTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	bestProduct, err := s.repo.GetBestSellingProduct(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &models.ReportResponse{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransactions,
		ProdukTerlaris: bestProduct,
	}, nil
}
