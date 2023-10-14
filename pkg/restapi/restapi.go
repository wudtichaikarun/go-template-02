package restapi

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wudtichaikarun/poc-go-template-02/configs"
	restapi_gin "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/gin"
)

type Router struct {
	*restapi_gin.GinRouter
}

func (r *Router) NewRouterGroup(g *gin.RouterGroup) *restapi_gin.GinRouterGroup {
	return &restapi_gin.GinRouterGroup{RouterGroup: g}
}

type RestServer struct {
	Router *Router
	Server *http.Server
	Config *configs.Configuration
}

func New(config *configs.Configuration) *RestServer {
	router := &Router{restapi_gin.NewGinRouter()}
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: router,
	}
	return &RestServer{Config: config, Router: router, Server: server}
}

func (r *RestServer) Start() error {
	fmt.Printf("Starting server on %s:%d", r.Config.Server.Host, r.Config.Server.Port)
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := r.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil
}

func (r *RestServer) WaitForStop() error {
	//  Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	fmt.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown:: %v\n", err)
		return err
	}
	fmt.Println("Server stopped.")

	return nil
}
