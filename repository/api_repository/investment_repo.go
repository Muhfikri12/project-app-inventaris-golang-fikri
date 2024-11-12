package apirepository

import (
	"database/sql"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
)

type InvestmentRepoDB struct {
	db *sql.DB
}

func NewInvestmentRepo(db *sql.DB) *InvestmentRepoDB {
	return &InvestmentRepoDB{db}
}

func (i *InvestmentRepoDB) CalculateDepresiation() (*[]model.Items, error) {
	query := `SELECT id, price, deprisiasi, transaction_date FROM items WHERE deleted_at IS NULL`

	items := []model.Items{}

	rows, err := i.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := model.Items{}

		err := rows.Scan(&item.ID, &item.Price, &item.Depreciation, &item.TransactionDate)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}


	return &items, err
}

func (i *InvestmentRepoDB) GetItemWithDepreciation(id int) (*model.Items, error) {
	item := model.Items{}
	query := `SELECT id, name, price, deprisiasi, transaction_date FROM items WHERE id=$1 AND deleted_at IS NULL`

	err := i.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Price, &item.Depreciation, &item.TransactionDate)
	if err != nil {
		return nil, err
	}
	
	return &item, nil
}