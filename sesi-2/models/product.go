package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(trx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) BeforeUpdate(trx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
