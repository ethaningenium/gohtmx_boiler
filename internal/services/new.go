package services

import (
	"fmt"
	"templtest/internal/entities"

	"github.com/google/uuid"
)

type DB interface {
	CreateTodo(todo entities.Todo) error
	GetTodos(userID string) ([]entities.Todo, error)
	UpdateTodo(todo entities.Todo) error
	DeleteTodo(string, string) error
	GetUser(Email string) (entities.User, error)
	CreateUser(user entities.User) error
}

type Service struct {
	db DB
}

func New(db DB) *Service {
	return &Service{db}
}

func (s *Service) UserLogin(Email, Name, Password string) (entities.User, error) {
	usergetted, err := s.db.GetUser(Email)
	if err != nil {
		fmt.Println(err.Error())
		user := entities.User{
			ID:       uuid.New().String(),
			Email:    Email,
			Name:     Name,
			Password: Password,
		}
		err = s.db.CreateUser(user)
		if err != nil {
			fmt.Println(err.Error())
			return entities.User{}, err
		}
		return user, nil
	}
	return usergetted, nil
}

func (s *Service) User(Email string) (entities.User, error) {
	return s.db.GetUser(Email)
}

func (s *Service) Todos(UserID string) ([]entities.Todo, error) {
	return s.db.GetTodos(UserID)
}

func (s *Service) CreateTodos(Title, UserID string) (entities.Todo, error) {
	newTodo := entities.Todo{
		ID:          uuid.New().String(),
		Title:       Title,
		IsCompleted: "false",
		UserID:      UserID,
	}

	return newTodo, s.db.CreateTodo(newTodo)
}

func (s *Service) DeleteTodo(ID string, UserID string) error {
	return s.db.DeleteTodo(ID, UserID)
}
