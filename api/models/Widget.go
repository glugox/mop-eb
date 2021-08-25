package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type WidgetType int

const (
	TypeTable WidgetType = iota + 1
	TypeChart
)

type Widget   struct {
	ID        uint32         `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32         `json:"user_id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Alias     string         `gorm:"size:255;not null" json:"alias"`
	Ordering  int            `gorm:"not null" json:"ordering"`
	Type      int            `gorm:"not null" json:"type"`
	IsPrivate bool           `gorm:"not null" json:"is_private"`
}


func (w *Widget) FindAllWidgets(db *gorm.DB, userID uint32) (*[]Widget, error) {
	var err error
	ws := []Widget{}

	// TODO: Make only one query builder
	if userID == 0 {
		err = db.Debug().Model(&Widget{}).Where("user_id = 0").Where("is_private = 0").Order("ordering").Find(&ws).Error
	} else {
		err = db.Debug().Model(&Widget{}).Where("user_id = ?", userID).Where("is_private = 1").Order("ordering").Find(&ws).Error
	}
	if err != nil {
		return &[]Widget{}, err
	}
	return &ws, nil
}

func (w *Widget) FindWidgetByID(db *gorm.DB, uid uint32) (*Widget, error) {
	var err error
	err = db.Debug().Model(Widget{}).Where("id = ?", uid).Take(&w).Error
	if err != nil {
		return &Widget{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Widget{}, errors.New("Widget Not Found")
	}
	return w, err
}