package sqlite

import (
	"templtest/internal/entities"

	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) GetUser(Email string) (entities.User, error) {
	var user entities.User

	query, args, err := sq.Select("*").
		From("users").
		Where(sq.Eq{"email": Email}).
		ToSql()
	if err != nil {
		return user, err
	}

	err = repo.db.Get(&user, query, args...)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *Repository) CreateUser(user entities.User) error {

	query, args, err := sq.Insert("users").
		Columns("id", "email", "name", "password").
		Values(user.ID, user.Email, user.Name, user.Password).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
