package service

import (
	"github.com/lukinairina90/REST_API_TODO_list"
	"github.com/lukinairina90/REST_API_TODO_list/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item REST_API_TODO_list.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]REST_API_TODO_list.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (REST_API_TODO_list.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input REST_API_TODO_list.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
