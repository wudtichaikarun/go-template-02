package controllers_example

import restapi_context "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/context"

type ExampleController interface {
	AddExample(c restapi_context.Context)
	ListExample(c restapi_context.Context)
}
