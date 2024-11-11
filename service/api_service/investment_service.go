package apiservice

import (
	"time"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
)

type InvestmentService struct {
	Repo *apirepository.InvestmentRepoDB
}

func NewInvestmentService(repo *apirepository.InvestmentRepoDB) *InvestmentService {
	return &InvestmentService{Repo: repo}
}

func (is *InvestmentService) GetTotalInvestment() (*model.Investment, error) {
	items, err := is.Repo.CalculateDepresiation()
	if err != nil {
		return nil, err
	}

	var totalInvestment int
	var totalDepreciatedValue int

	for _, item := range *items {
		
		totalInvestment += item.Price

		months := is.CalculateDurationMonths(item.TransactionDate, time.Now())
	
		currentPrice := item.Price

		for month := 0; month < months; month++ {
			monthlyDepreciation := (currentPrice * item.Depreciation / 100)
			currentPrice -= monthlyDepreciation
		}

		itemDepreciation := item.Price - currentPrice
		totalDepreciatedValue += itemDepreciation
	}

	// Membuat respons
	item := model.Investment{
		TotalInvestment:  totalInvestment,
		DepreciatedValue: totalDepreciatedValue,
	}

	return &item, nil
}


func (is *InvestmentService) CalculateDurationMonths(start time.Time, end time.Time) int {
	years := end.Year() - start.Year()
	months := int(end.Month() - start.Month())
	return years*12 + months
}