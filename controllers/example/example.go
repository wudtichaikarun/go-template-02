package controllers_example

import (
	restapi_context "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/context"
	restapi_response "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/response"
	"github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
	services_example "github.com/wudtichaikarun/poc-go-template-02/services/example"
)

type exampleController struct {
	exampleSvc services_example.ExampleService
}

// @Summary      create example
// @Description  create example
// @Tags         example
// @Router       /v1/examples [post]
// @Accept       json
// @Produce      json
// @Param        request  body  request.ExampleReq  true "create example body"
// @Success      201
func (ec *exampleController) AddExample(c restapi_context.Context) {
	var example request.ExampleReq

	if err := c.Bind(&example); err != nil {
		restapi_response.BadRequest(c, "invalid request")
		return
	}

	if err := ec.exampleSvc.Create(example); err != nil {
		restapi_response.InternalServerError(c, "create error")
		return
	}

	restapi_response.Created(c)
}

// @Summary      list all example
// @Description  list all example
// @Tags         example
// @Router       /v1/examples [get]
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ExampleListRes
func (ec *exampleController) ListExample(c restapi_context.Context) {
	panic("unimplemented")
}

func New(exampleSvc services_example.ExampleService) ExampleController {
	return &exampleController{exampleSvc}
}
