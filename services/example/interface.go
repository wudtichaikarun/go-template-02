package services_example

import (
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
)

type ExampleService interface {
	Create(r request.ExampleReq) error
	List() ([]models.Example, error)
}
