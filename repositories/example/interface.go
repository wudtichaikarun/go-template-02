package repositories_example

import (
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
)

type ExampleRepository interface {
	List() ([]models.Example, error)
	Create(request.ExampleReq) error
}
