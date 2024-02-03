package data

import (
	"MyEcommerce/features/product"
	"errors"
	"log"

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
func (repo *productQuery) SelectAll(page, limit int, category string) ([]product.Core, int, error) {
	var products []Product
	query := repo.db.Order("created_at desc")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	
	var totalData int64
	err := query.Model(&products).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	// var products []Product
    // if category == "" {
    //     err := repo.db.Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&products).Error
    //     if err != nil {
    //         return nil, err
    //     }
    // } else {
    //     err := repo.db.Order("created_at desc").Limit(limit).Offset((page-1)*limit).Where("category = ?", category).Find(&products).Error
    //     if err != nil {
    //         return nil, err
    //     }
    // }

	err = query.Limit(limit).Offset((page - 1) * limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}


	var productCores []product.Core
	for _, p := range products {
		productCores = append(productCores, p.ModelToCore())
	}

	return productCores, totalPage, nil
}

// SelectById implements product.ProductDataInterface.
func (repo *productQuery) SelectById(IdProduct int) (*product.Core, error) {
	var productDataGorm Product
	tx := repo.db.Preload("User").Where("id = ?", IdProduct).First(&productDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := productDataGorm.ModelToCore()
	return &result, nil
}

// Update implements product.ProductDataInterface.
func (repo *productQuery) Update(userIdLogin int, input product.Core) error {
	product := Product{}
	tx := repo.db.Where("id = ? AND user_id = ?", input.ID, userIdLogin).First(&product)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errors.New("you do not have permission to edit this product")
		}
		return tx.Error
	}

	productInputGorm := CoreToModel(input)

	tx = repo.db.Model(&product).Updates(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Delete implements product.ProductDataInterface.
func (repo *productQuery) Delete(userIdLogin, IdProduct int) error {
	tx := repo.db.Where("id = ? AND user_id = ?", IdProduct, userIdLogin).Delete(&Product{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

// SelectByUserId implements product.ProductDataInterface.
func (repo *productQuery) SelectByUserId(userIdLogin int) ([]product.Core, error) {
	var productDataGorms []Product
	tx := repo.db.Where("user_id = ?", userIdLogin).Find(&productDataGorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var results []product.Core
	for _, productDataGorm := range productDataGorms {
		result := productDataGorm.ModelToCore()
		results = append(results, result)
	}
	return results, nil
}

// Search implements product.ProductDataInterface.
func (repo *productQuery) Search(query string) ([]product.Core, error) {
	var productDataGorms []Product
	log.Println("query", query)
	tx := repo.db.Where("name LIKE ?", "%"+query+"%").Find(&productDataGorms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var results []product.Core
	for _, productDataGorm := range productDataGorms {
		result := productDataGorm.ModelToCore()
		results = append(results, result)
	}
	return results, nil
}
