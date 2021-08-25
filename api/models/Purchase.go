package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Purchase   struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint32    `json:"user_id"`
	ProductID   uint32    `json:"product_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Purchase) Validate() error {

	if p.UserID == 0 {
		return errors.New("User Empty")
	}
	if p.ProductID == 0 {
		return errors.New("Product Empty")
	}

	// TODO: Validate foreign keys?

	return nil
}

func (p *Purchase) SavePurchase(db *gorm.DB) (*Purchase, error) {

	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Purchase{}, err
	}
	return p, nil
}

func (p *Purchase) FindAllPurchases(db *gorm.DB, userID uint32) (*[]Purchase, error) {
	var err error
	ps := []Purchase{}
	if userID != 0{
		err = db.Debug().Model(&Purchase{}).Where("user_id = ?", userID).Find(&ps).Error
	} else {
		err = db.Debug().Model(&Purchase{}).Find(&ps).Error
	}
	if err != nil {
		return &[]Purchase{}, err
	}
	return &ps, nil
}

func (p *Purchase) FindPurchaseByID(db *gorm.DB, uid uint32) (*Purchase, error) {
	var err error
	err = db.Debug().Model(Purchase{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Purchase{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Purchase{}, errors.New("Purchase Not Found")
	}
	return p, err
}