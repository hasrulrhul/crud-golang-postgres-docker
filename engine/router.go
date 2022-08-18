package engine

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/docker-golang-crud/app/controllers"
	"github.com/hasrulrhul/docker-golang-crud/config"
	"github.com/hasrulrhul/docker-golang-crud/repository"
	"github.com/hasrulrhul/docker-golang-crud/service"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	db             *gorm.DB                   = config.ConnectDatabase()
	todoRepository repository.TodoRepository  = repository.NewTodoRepository(db)
	todoService    service.TodoService        = service.NewTodoService(todoRepository)
	todoController controllers.TodoController = controllers.NewTodoController(todoService)
)

func SetupRouter() *gin.Engine {
	// defer config.CloseConnection(db)

	// Gin instance
	r := gin.Default()
	// }
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	// Routes
	v1 := r.Group("api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "CRUD Golang Postgres with Docker",
			})
		})
		todos := v1.Group("/todo")
		{
			todos.GET("", todoController.Index)
			todos.POST("/", todoController.Create)
			todos.GET("/:id", todoController.Show)
			todos.PUT("/:id", todoController.Update)
			todos.DELETE("/:id", todoController.Delete)
		}
	}
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			//c.Next()
			return
		}
		c.Next()
	}
}
