package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/smockoro/mysql-master-slave/sample-go/domain/model"
	repo "github.com/smockoro/mysql-master-slave/sample-go/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository : ...
func NewUserRepository(db *sqlx.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *model.User) (int64, error) {
	res, err := ur.db.NamedExecContext(ctx,
		"INSERT INTO users(`name`, `age`) VALUES(:name, :age)",
		user)
	if err != nil {
		return -1, fmt.Errorf("failed to insert user" + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("failed to retrieve user id" + err.Error())
	}

	return id, nil
}

func (ur *userRepository) Find(ctx context.Context, id int64) (*model.User, error) {
	res, err := ur.db.QueryxContext(ctx,
		"SELECT `id`, `name`, `age` FROM users WHERE `id` = ?",
		id)
	if err != nil {
		return nil, fmt.Errorf("failed to select operation" + err.Error())
	}
	defer res.Close()

	if !res.Next() {
		if err := res.Err(); err != nil {
			return nil, fmt.Errorf("failed to get data " + err.Error())
		}
		return nil, fmt.Errorf(fmt.Sprintf("ID='%d' is not found", id))
	}

	var user model.User
	if err := res.StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *model.User) (int64, error) {
	res, err := ur.db.NamedExecContext(ctx,
		"UPDATE users SET `name`=:name, `age`=:age WHERE `id`=:id",
		user)
	if err != nil {
		return -1, fmt.Errorf("failed to update user" + err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf(fmt.Sprintf("user id %d is not found", user.ID))
	}

	return rows, nil
}

func (ur *userRepository) Delete(ctx context.Context, id int64) (int64, error) {
	res, err := ur.db.ExecContext(ctx, "DELETE FROM users WHERE `id`= ?", id)
	if err != nil {
		return -1, fmt.Errorf("failed to delete " + err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf(fmt.Sprintf("ID='%d' is not found", id))
	}

	return rows, nil
}
