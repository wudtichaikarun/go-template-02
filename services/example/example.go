package services_example

import (
	repositories_example "github.com/wudtichaikarun/poc-go-template-02/repositories/example"
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
)

type exampleService struct {
	exampleRepo repositories_example.ExampleRepository
}

// Create implements ExampleService.
func (*exampleService) Create(r request.ExampleReq) error {
	panic("unimplemented")
}

// List implements ExampleService.
func (es *exampleService) List() ([]models.Example, error) {
	return es.exampleRepo.List()
}

func New(exampleRepo repositories_example.ExampleRepository) ExampleService {
	return &exampleService{exampleRepo}
}
