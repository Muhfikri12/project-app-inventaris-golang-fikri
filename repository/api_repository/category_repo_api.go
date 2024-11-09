package apirepository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
)

type CategoryRepoDB struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepoDB {
	return &CategoryRepoDB{DB: db}
}

func (c *CategoryRepoDB) ShowAllCategories() (*[]model.Categories, error) {
	query := `SELECT id, name, description FROM categories WHERE deleted_at IS NULL`
	
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Categories

	for rows.Next() {
		var category model.Categories
		// var category string
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &categories, nil
}

func (c *CategoryRepoDB) CreateCategory(category *model.Categories) error {
	query := `INSERT INTO categories(name, description) VALUES($1, $2) RETURNING id`
	err := c.DB.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	
	if err != nil {
		return err
	}

	return nil
}