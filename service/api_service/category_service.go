package apiservice

import (
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
)

type ServiceCategory struct {
	RepoService *apirepository.CategoryRepoDB
}

func NewServiceCategory(repo *apirepository.CategoryRepoDB) *ServiceCategory {
	return &ServiceCategory{RepoService: repo}
}

func (sc *ServiceCategory) GetAllCategory() (*[]model.Categories, error) {
	
	category, err := sc.RepoService.ShowAllCategories()
	if err != nil {
		return nil, err
	}
	
	return category, nil
}

func (sc *ServiceCategory) CreateCategoryService(category *model.Categories) error{

	if err := sc.RepoService.CreateCategory(category); err != nil {
		return err
	}
	
	return nil
}