package main

import (
	"log"
	"os"

	"github.com/wudtichaikarun/poc-go-template-02/configs"
	_ "github.com/wudtichaikarun/poc-go-template-02/docs"
	"github.com/wudtichaikarun/poc-go-template-02/pkg/database"
	database_mysql "github.com/wudtichaikarun/poc-go-template-02/pkg/database/mysql"
	"github.com/wudtichaikarun/poc-go-template-02/pkg/restapi"
	"github.com/wudtichaikarun/poc-go-template-02/routes"
)

// @title [Go template]
// @version 1.0
// @description poc go template
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	// load configs
	configuration, err := configs.GetConfig()
	if err != nil {
		panic("Failed to get configurations: " + err.Error())
	}

	// ===== database =====
	var db database.Database
	// use mysql can switch to other
	mysqlDB := &database_mysql.MySQLDatabase{}
	if err := mysqlDB.Connect(configs.DBConnection(), configuration.Server.Debug); err != nil {
		panic("Failed to connect to MySQL database: " + err.Error())
	}
	db = mysqlDB
	defer db.Close()
	// ===== database =====

	// rest api
	restApiInstance := restapi.New(configuration)
	routes.SetupRouters(*restApiInstance.Router, db.GetDB(), *configuration)
	restApiInstance.Start()
	restApiInstance.WaitForStop() // graceful-shutdown

}
