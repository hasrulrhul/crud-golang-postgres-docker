package repository

import (
	"github.com/hasrulrhul/docker-golang-crud/app/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Index() []models.Todo
	Create(model models.Todo) models.Todo
	Show(id string) models.Todo
	Update(model models.Todo) models.Todo
	Delete(user models.Todo) models.Todo
}

type userConnection struct {
	connection *gorm.DB
}

//NewTodoRepository is creates a new instance of TodoRepository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Index() []models.Todo {
	var user []models.Todo
	db.connection.Find(&user)
	return user
}

func (db *userConnection) Create(model models.Todo) models.Todo {
	db.connection.Save(&model)
	return model
}

func (db *userConnection) Show(id string) models.Todo {
	var user models.Todo
	db.connection.Find(&user).Where("id = ?", id)
	return user
}

func (db *userConnection) Update(model models.Todo) models.Todo {
	db.connection.Updates(&model)
	db.connection.Find(&model)
	return model
}

func (db *userConnection) Delete(user models.Todo) models.Todo {
	db.connection.Delete(&user)
	return user
}
