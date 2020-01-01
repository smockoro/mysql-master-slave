package repository

import "github.com/smockoro/mysql-master-slave/sample-go/domain/model"

// UserRepository : ...
type UserRepository interface {
	Create(model.User) (int64, error)
	Find(int64) (model.User, error)
	Update(model.User) (int64, error)
	Delete(int64) (int64, error)
}
