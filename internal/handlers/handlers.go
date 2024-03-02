package handlers

import (
	"github.com/exxxception/restful-service/internal/services"
	"github.com/gin-gonic/gin"

	_ "github.com/exxxception/restful-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	services   *services.Services
	apiAuthKey string
}

func NewHandlers(services *services.Services, apiAuthKey string) *Handlers {
	return &Handlers{
		services:   services,
		apiAuthKey: apiAuthKey,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api", h.userIdentity)
	{
		accounts := api.Group("/accounts")
		{
			accounts.POST("/", h.Create)
			accounts.GET("/:id", h.GetInfo)
			accounts.PUT("/deposit/:id", h.Deposit)
			accounts.PUT("/withdraw/:id", h.Withdraw)
		}
	}

	return router
}
