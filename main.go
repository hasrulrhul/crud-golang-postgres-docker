package main

import (
	"os"

	"github.com/hasrulrhul/docker-golang-crud/engine"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	config.ConnectDatabase()
// }

func main() {

	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "CRUD Golang Postgres with Docker",
	// 	})
	// })

	// router.GET("todo", controllers.Index)
	// router.POST("todo", controllers.Create)
	// router.GET("todo/:id", controllers.Show)
	// router.PUT("todo/:id", controllers.Update)
	// router.DELETE("todo/:id", controllers.Delete)

	// router.Run(":" + os.Getenv("APP_PORT"))
	r := engine.SetupRouter()

	r.Run(":" + os.Getenv("APP_PORT"))
}
