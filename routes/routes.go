package routes

import (
	"github.com/wudtichaikarun/poc-go-template-02/configs"
	"github.com/wudtichaikarun/poc-go-template-02/controllers"
	controllers_example "github.com/wudtichaikarun/poc-go-template-02/controllers/example"
	"github.com/wudtichaikarun/poc-go-template-02/pkg/restapi"
	restapi_context "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/context"
	restapi_response "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/response"
	repositories_example "github.com/wudtichaikarun/poc-go-template-02/repositories/example"
	services_example "github.com/wudtichaikarun/poc-go-template-02/services/example"
	"gorm.io/gorm"
)

func SetupRouters(route restapi.Router, db *gorm.DB, config configs.Configuration) {
	// for mock error
	route.GET("/error", func(ctx restapi_context.Context) {
		restapi_response.InternalServerError(ctx, "test err")
	})

	// initial repositories
	exampleRepo := repositories_example.New(db)

	// initial services
	exampleSvc := services_example.New(exampleRepo)

	// initial controllers
	systemCtrl := controllers.NewSystemController()
	exampleCtrl := controllers_example.New(exampleSvc)

	system := route.NewRouterGroup(route.Group("/system"))
	system.GET("/health", systemCtrl.HealthCheck)

	v1 := route.NewRouterGroup(route.Group("/v1"))

	// banners
	v1.GET("/examples", exampleCtrl.ListExample)
	v1.POST("/examples", exampleCtrl.AddExample)

}
