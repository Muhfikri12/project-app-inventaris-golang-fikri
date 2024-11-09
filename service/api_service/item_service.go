package apiservice

import (
	"fmt"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apirepository "github.com/Muhfikri12/project-app-inventaris-golang-fikri/repository/api_repository"
)

type ItemService struct {
	ItemRepo apirepository.ItemRepoDB
}

func NewItemService(repo apirepository.ItemRepoDB) ItemService {
	return ItemService{ItemRepo: repo}
}

func (is *ItemService) CreateItemService(item *model.Items) error{

	if err := is.ItemRepo.CreateItem(item); err != nil {
		return err
	}
	
	return nil
}

func (is *ItemService) ShowAllItems(page, limit int, pagination *model.Response) (*[]model.Items, error){

	item, err := is.ItemRepo.ShowItems(page, limit, pagination)
	if err != nil {
		return nil, err
	}
	
	return item, nil
}

func (is *ItemService) GetItemByID(id int) (*model.Items, error) {
	
	item, err := is.ItemRepo.GetItemByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get item by ID: %w", err)
	} 

	return item, err
}

func (is *ItemService) UpdateItemID(id int, item *model.Items) error {
	err := is.ItemRepo.UpdateItemID(id, item)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

func (is *ItemService) SoftDeleteItemService(id int) error {

    if err := is.ItemRepo.SoftDeleteItem(id); err != nil {
        return fmt.Errorf("failed to delete item: %w", err)
    }

    return nil
}
