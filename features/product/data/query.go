package data

import (
	"MyEcommerce/features/product"
	"errors"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) Insert(userIdLogin int, input product.Core) error {

	productInputGorm := CoreToModel(input)
	productInputGorm.UserID = uint(userIdLogin)

	tx := repo.db.Create(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectAll implements product.ProductDataInterface.
func (repo *productQuery) SelectAll() ([]product.Core, error) {
	var products []Product
	err := repo.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	var productCores []product.Core
	for _, p := range products {
		productCores = append(productCores, p.ModelToCore())
	}

	return productCores, nil
}
