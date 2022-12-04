package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" valid:"required~Name of name social media"`
	SocialMediaUrl string `gorm:"not null" valid:"required~Url of your photo url is required"`
	UserID         uint
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil

	return

}
