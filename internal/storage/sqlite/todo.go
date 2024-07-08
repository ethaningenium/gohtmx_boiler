package sqlite

import (
	"templtest/internal/entities"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) CreateTodo(todo entities.Todo) error {
	// Создание запроса с использованием Squirrel
	query, args, err := sq.Insert("todos").
		Columns("id", "title", "is_completed", "user_id").
		Values(todo.ID, todo.Title, todo.IsCompleted, todo.UserID).
		ToSql()
	if err != nil {
		return err
	}

	// Выполнение запроса
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Функция для получения всех задач пользователя с использованием Squirrel
func (repo *Repository) GetTodos(userID string) ([]entities.Todo, error) {
	var todos []entities.Todo

	// Создание запроса с использованием Squirrel
	query, args, err := squirrel.Select("*").
		From("todos").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	// Выполнение запроса и сканирование результата в срез Todo
	err = repo.db.Select(&todos, query, args...)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Функция для обновления задачи (todo) с использованием Squirrel
func (repo *Repository) UpdateTodo(todo entities.Todo) error {
	// Создание запроса с использованием Squirrel
	query, args, err := squirrel.Update("todos").
		Set("title", todo.Title).
		Set("is_completed", todo.IsCompleted).
		Where(squirrel.Eq{"id": todo.ID}).
		ToSql()
	if err != nil {
		return err
	}

	// Выполнение запроса
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteTodo(ID string, UserID string) error {
	query, args, err := sq.Delete("todos").
		Where(squirrel.Eq{"id": ID, "user_id": UserID}).ToSql()
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
