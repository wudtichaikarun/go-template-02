package controllers

import (
	restapi_context "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/context"
	restapi_response "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/response"
)

type SuccessResponse struct {
	Status  int                      `json:"status"`
	Data    SystemControllerResponse `json:"data"`
	Message string                   `json:"message"`
}

type SystemControllerResponse struct {
	Live string `json:"live"`
}

type SystemController struct{}

func NewSystemController() *SystemController {
	return &SystemController{}
}

// @Summary      system health check
// @Description  system health check
// @Tags         system
// @Router       /system/health [get]
// @Accept       json
// @Produce      json
// @Success      200  {object}  controllers.SuccessResponse
func (sc *SystemController) HealthCheck(ctx restapi_context.Context) {
	restapi_response.Success(ctx, &restapi_response.SuccessResponse{
		Data: map[string]interface{}{
			"live": "ok",
		},
	})
}
