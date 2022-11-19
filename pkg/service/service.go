package service

import (
	"github.com/lukinairina90/REST_API_TODO_list"
	"github.com/lukinairina90/REST_API_TODO_list/pkg/repository"
)

type Authorization interface {
	CreateUser(user REST_API_TODO_list.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list REST_API_TODO_list.TodoList) (int, error)
	GetAll(userId int) ([]REST_API_TODO_list.TodoList, error)
	GetById(userId, listId int) (REST_API_TODO_list.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input REST_API_TODO_list.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item REST_API_TODO_list.TodoItem) (int, error)
	GetAll(userId, listId int) ([]REST_API_TODO_list.TodoItem, error)
	GetById(userId, itemId int) (REST_API_TODO_list.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input REST_API_TODO_list.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
