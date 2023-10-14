package repositories_example

import (
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
	"gorm.io/gorm"
)

type exampleRepo struct {
	conn *gorm.DB
}

// Create implements ExampleRepository.
func (*exampleRepo) Create(r request.ExampleReq) error {
	panic("unimplemented")
}

// FindAll implements ExampleRepository.
func (*exampleRepo) FindAll() (r *[]*models.Example, err error) {
	panic("unimplemented")
}

func New(conn *gorm.DB) ExampleRepository {
	return &exampleRepo{conn}
}
