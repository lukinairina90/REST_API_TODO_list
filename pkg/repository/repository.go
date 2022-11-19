package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lukinairina90/REST_API_TODO_list"
)

type Authorization interface {
	CreateUser(user REST_API_TODO_list.User) (int, error)
	GetUser(username, password string) (REST_API_TODO_list.User, error)
}

type TodoList interface {
	Create(userId int, list REST_API_TODO_list.TodoList) (int, error)
	GetAll(userId int) ([]REST_API_TODO_list.TodoList, error)
	GetById(userId, listId int) (REST_API_TODO_list.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input REST_API_TODO_list.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item REST_API_TODO_list.TodoItem) (int, error)
	GetAll(userId, listId int) ([]REST_API_TODO_list.TodoItem, error)
	GetById(userId, itemId int) (REST_API_TODO_list.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input REST_API_TODO_list.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
