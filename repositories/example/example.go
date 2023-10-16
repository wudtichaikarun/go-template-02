package repositories_example

import (
	"github.com/wudtichaikarun/poc-go-template-02/pkg/utils"
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
	"gorm.io/gorm"
)

type exampleRepo struct {
	conn *gorm.DB
}

func (*exampleRepo) List() ([]models.Example, error) {
	panic("unimplemented")
}

func (repo *exampleRepo) Create(r request.ExampleReq) error {
	var m models.Example
	utils.ConvertInterfaceToStruct(r, &m)
	return repo.conn.Create(&m).Error
}

func New(conn *gorm.DB) ExampleRepository {
	return &exampleRepo{conn}
}
