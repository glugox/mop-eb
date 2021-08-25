package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID       uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:100;not null;unique" json:"name"`
	Price    int       `gorm:"not null" json:"price"`
}


func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("Name Empty")
	}

	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {

	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Product{}, err
	}

	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	ps := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&ps).Error
	if err != nil {
		return &[]Product{}, err
	}

	return &ps, nil
}

func (p *Product) FindProductByID(db *gorm.DB, uid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}

	return p, err
}