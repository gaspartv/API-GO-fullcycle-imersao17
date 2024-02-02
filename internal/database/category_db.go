package database

import (
	"database/sql"

	"github.com/gaspartv/API-GO-fullcycle-imersao17/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{
		db: db,
	}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	row := cd.db.QueryRow("SELECT id, name FROM category WHERE id = ?", id)
	var category entity.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Query("INSERT INTO category (id, name) VALUES ($1)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
