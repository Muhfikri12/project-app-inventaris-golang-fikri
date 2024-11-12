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

		months := is.calculateDurationMonths(item.TransactionDate, time.Now())
	
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


func (is *InvestmentService) calculateDurationMonths(start time.Time, end time.Time) int {
	years := end.Year() - start.Year()
	months := int(end.Month() - start.Month())
	return years*12 + months
}

func (is *InvestmentService) GetItemWithDepreciationID(id int) (*model.Items,error) {
	
	item, err := is.Repo.GetItemWithDepreciation(id)
	if err != nil {
		return nil, err
	}

	var depreciationValue int

	months := is.calculateDurationMonths(item.TransactionDate, time.Now())

	currentPrice := item.Price

	for month := 0; month < months; month++ {
		monthlyDepreciation := (currentPrice * item.Depreciation / 100)
		currentPrice -= monthlyDepreciation
	}
	
	itemDepreciation := item.Price - currentPrice
	depreciationValue += itemDepreciation

	item = &model.Items {
		ID: id,
		Name: item.Name,
		Price: item.Price,
		TransactionDate: item.TransactionDate,
		DepreciatedValue: depreciationValue,
		Depreciation: item.Depreciation,
	}

	return item, nil
}

func (is *InvestmentService) GetIDShouldBeReplaced() (*[]model.Items, error) {
	
	items , err := is.Repo.CalculateDepresiation()
	if err != nil {
		return nil, err
	}

	var itemsToBeReplaced []model.Items

	for _, item := range *items {
		months := is.calculateDurationMonths(item.TransactionDate, time.Now())
		if months > 5 {
			item.Replaced = true 
			item.TotalUsage = months
			itemsToBeReplaced = append(itemsToBeReplaced, item)
		}
	}


	return &itemsToBeReplaced, nil
}