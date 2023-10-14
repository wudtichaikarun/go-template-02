package services_example

import (
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/response"
)

type ExampleService interface {
	Create(r request.ExampleReq) error
	FindAll() (t *response.ExampleListRes, err error)
}
