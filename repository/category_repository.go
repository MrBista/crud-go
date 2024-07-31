package repository

import (
	"context"
	"database/sql"

	"github.com/MrBista/crud-go/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, cateogry domain.Category) bool
	FindById(ctx context.Context, tx *sql.Tx, cateogryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
