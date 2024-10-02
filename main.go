package main

import (
	"ApiDoc/docs"
	"ApiDoc/handlers"
	"ApiDoc/models"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

var users = []models.Users{}

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	//creating a defult router using the GIN framework in go that will be  user to handle different type of route requests

	router.POST("/adduser", func(c *gin.Context) {
		handlers.AddUserHandler(&users, c)
	})

	router.GET("/getUser/:id", func(c *gin.Context) {
		handlers.GetUserInfo(&users, c)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Println("Could't run on the server", err)
	} else {
		fmt.Println("Server is up and running listening to the routes")
	}

}
