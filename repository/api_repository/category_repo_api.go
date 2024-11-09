package apirepository

import (
	"database/sql"
	"fmt"

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

func (c *CategoryRepoDB) GetCatgoryByID(id int) (*model.Categories, error) {
	category := model.Categories{}

	query := `SELECT id, name, description
		FROM categories 
		WHERE id=$1 
		AND deleted_at IS NULL`

	if err := c.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}
	
	return &category, nil
}

func (c *CategoryRepoDB) UpdateCategory(id int, category *model.Categories) error {
	query := `
		UPDATE categories
		SET name = $1, description = $2
		WHERE id = $3 AND deleted_at IS NULL
	`
	result, err := c.DB.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("category not found or no changes made")
	}

	return nil
}
