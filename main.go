package main

import (
	"fmt"
	"log"

	"music-library/api"
	"music-library/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "music-library/docs" // Импортируем сгенерированную документацию Swagger
)

// @title Music Library API
// @version 1.0
// @description This is a simple API for managing a music library.
// @termsOfService http://swagger.io/terms/
// @contact.name Alisher Alishev
// @contact.email alisheralishev4444@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	database.InitDB()
	defer database.CloseDB()
	r := gin.Default()
	api.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})

	log.Println("Starting server on :8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}