package models

import "github.com/wudtichaikarun/poc-go-template-02/pkg/model"

type Example struct {
	model.Model
	Name        string `gorm:"type:varchar(255);" json:"name" form:"name"`
	Description string `gorm:"type:varchar(255);" json:"description" form:"description"`
}
