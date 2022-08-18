package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/docker-golang-crud/app/dto"
	"github.com/hasrulrhul/docker-golang-crud/app/models"
	"github.com/hasrulrhul/docker-golang-crud/service"
)

//TodoController is a contract what this controller can do
type TodoController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

//NewTodoController create a new instances of TodoController
func NewTodoController(todoServ service.TodoService) TodoController {
	return &todoController{
		todoService: todoServ,
	}
}

func (c *todoController) Index(ctx *gin.Context) {
	var todos []models.Todo = c.todoService.Index()
	ctx.JSON(http.StatusOK, todos)
}

func (c *todoController) Create(ctx *gin.Context) {
	var todoReq dto.TodoDTO
	errRequest := ctx.ShouldBind(&todoReq)
	if errRequest != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errRequest.Error())
		return
	}

	todo := c.todoService.Create(todoReq)
	ctx.JSON(http.StatusCreated, todo)
}

func (c *todoController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo = c.todoService.Show(id)
	if (todo == models.Todo{}) {
		ctx.JSON(http.StatusNotFound, "Data not found")
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func (c *todoController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo = c.todoService.Show(id)
	if (todo == models.Todo{}) {
		ctx.JSON(http.StatusNotFound, "Data not found")
	} else {
		var todoReq dto.TodoDTO
		todoReq.ID = id
		errReq := ctx.ShouldBind(&todoReq)
		if errReq != nil {
			ctx.JSON(http.StatusBadRequest, errReq.Error())
		}
		result := c.todoService.Update(todoReq)
		ctx.JSON(http.StatusCreated, result)
	}
}

func (c *todoController) Delete(ctx *gin.Context) {
	var todo models.Todo = c.todoService.Show(ctx.Param("id"))
	if (todo == models.Todo{}) {
		ctx.JSON(http.StatusNotFound, "Data not found")
	} else {
		result := c.todoService.Delete(todo)
		ctx.JSON(http.StatusCreated, result)
	}
}
