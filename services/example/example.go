package services_example

import (
	repositories_example "github.com/wudtichaikarun/poc-go-template-02/repositories/example"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/response"
)

type exampleService struct {
	exampleRepo repositories_example.ExampleRepository
}

// Create implements ExampleService.
func (*exampleService) Create(r request.ExampleReq) error {
	panic("unimplemented")
}

// FindAll implements ExampleService.
func (*exampleService) FindAll() (t *response.ExampleListRes, err error) {
	panic("unimplemented")
}

func New(exampleRepo repositories_example.ExampleRepository) ExampleService {
	return &exampleService{exampleRepo}
}
