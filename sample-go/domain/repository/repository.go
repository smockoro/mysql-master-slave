package repository

import (
	"context"

	"github.com/smockoro/mysql-master-slave/sample-go/domain/model"
)

// UserRepository : ...
type UserRepository interface {
	Create(context.Context, *model.User) (int64, error)
	Find(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) (int64, error)
	Delete(context.Context, int64) (int64, error)
}
