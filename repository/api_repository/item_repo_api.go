package apirepository

import (
	"database/sql"
	"fmt"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
)	

type ItemRepoDB struct {
	DB *sql.DB
}

func NewItemRepo(db *sql.DB) ItemRepoDB {
	return ItemRepoDB{DB: db}
}

func (i *ItemRepoDB) CreateItem(item *model.Items) error {
	query := `INSERT INTO items(name, category_id, image, price, transaction_date, deprisiasi) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := i.DB.QueryRow(query, item.Name, item.CategoryId, item.Image, item.Price, item.TransactionDate, item.Deprisiasi).Scan(&item.ID)
	
	if err != nil {
		return err
	}

	return nil
}

func (i *ItemRepoDB) ShowItems(page, limit int, pagination *model.Response) (*[]model.Items, error) {

	var totalItems int
	queryCount := `SELECT count(id) from items`
	err := i.DB.QueryRow(queryCount).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	query := `
		SELECT p.id, p.name, c.name AS category_name, p.image, p.price, p.transaction_date, p.deprisiasi
		FROM items p
		LEFT JOIN categories c ON p.category_id = c.id
		LIMIT $1 
		OFFSET $2`

	rows, err := i.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Items

	for rows.Next() {
		var item model.Items
		// var category string
		err := rows.Scan(&item.ID, &item.Name, &item.CategoryName, &item.Image, &item.Price, &item.TransactionDate, &item.Deprisiasi)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	pagination.TotalItems = totalItems
	pagination.TotalPages = (totalItems + limit - 1) / limit

	return &items, nil
}

func (i *ItemRepoDB) GetItemByID(id int) (*model.Items, error) {
	item := model.Items{}

	query := `SELECT i.id, i.name, c.name, i.image, i.price, i.transaction_date 
	FROM items i 
	JOIN categories c ON i.category_id = c.id
	WHERE i.id=$1`

	if err := i.DB.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.CategoryName, &item.Image, &item.Price, &item.TransactionDate); err != nil {
		return nil, err
	}
	
	return &item, nil
}

func (i *ItemRepoDB) UpdateItemID(id int, item *model.Items) error {
	existingItem := &model.Items{}
	checkQuery := `SELECT id, image FROM items WHERE id = $1`
	err := i.DB.QueryRow(checkQuery, id).Scan(&existingItem.ID, &existingItem.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("item with id %d not found", id)
		}
		return err
	}

	if item.Image == "" {
		item.Image = existingItem.Image
	}

	updateQuery := `UPDATE items SET name=$1, category_id=$2, image=$3, price=$4, transaction_date=$5 WHERE id=$6`
	_, err = i.DB.Exec(updateQuery, item.Name, item.CategoryId, item.Image, item.Price, item.TransactionDate, id)
	if err != nil {
		return err
	}

	return nil
}
