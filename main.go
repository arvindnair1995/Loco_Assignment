package main

import (
	"LOCO_ASSIGNMENT/controllers"
	_ "LOCO_ASSIGNMENT/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Transaction Service API
// @version         1.0
// @description     Transaction Service CRUD APIs.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

    r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		txnService := v1.Group("/transactionservice")
		{
			txnService.PUT("/transaction/:transaction_id", controllers.CreateTransaction)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
