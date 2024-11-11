package apiservice

import (
	"fmt"

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

func (sc *ServiceCategory) GetCategoryByID(id int) (*model.Categories, error) {
	
	category, err := sc.RepoService.GetCatgoryByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by ID: %w", err)
	} 

	return category, nil
}

func (sc *ServiceCategory) UpdateCategory(id int, category *model.Categories) error {
	
	err := sc.RepoService.UpdateCategory(id, category)
	if err != nil {
		return fmt.Errorf("failed to update category by ID: %w", err)
	} 

	return nil
}

func (sc *ServiceCategory) SoftDeleteCatService(id int) error {

    if err := sc.RepoService.SoftDeleteCat(id); err != nil {
        return fmt.Errorf("failed to delete category: %w", err)
    }

    return nil
}