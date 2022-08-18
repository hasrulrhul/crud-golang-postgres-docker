package service

import (
	"log"

	"github.com/hasrulrhul/docker-golang-crud/app/dto"
	"github.com/hasrulrhul/docker-golang-crud/app/models"
	"github.com/hasrulrhul/docker-golang-crud/repository"
	"github.com/mashingan/smapping"
)

//TodoService is a ....
type TodoService interface {
	Index() []models.Todo
	Create(model dto.TodoDTO) models.Todo
	Show(id string) models.Todo
	Update(model dto.TodoDTO) models.Todo
	Delete(model models.Todo) models.Todo
}

type todoService struct {
	todoRepository repository.TodoRepository
}

//NewTodoService .....
func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) Index() []models.Todo {
	return service.todoRepository.Index()
}

func (service *todoService) Create(model dto.TodoDTO) models.Todo {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.Create(todo)
	return res
}

func (service *todoService) Show(id string) models.Todo {
	return service.todoRepository.Show(id)
}

func (service *todoService) Update(model dto.TodoDTO) models.Todo {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.todoRepository.Update(todo)
	return res
}

func (service *todoService) Delete(todo models.Todo) models.Todo {
	return service.todoRepository.Delete(todo)
}
