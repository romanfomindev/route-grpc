package category_repository

import (
	"context"

	"github.com/romanfomindev/route-grpc/tree/master/category-service/internal/service/category"
)

var categories = category.Categories{
	{
		ID:   1,
		Name: "Toys",
	},
	{
		ID:   2,
		Name: "Laptops",
	},
	{
		ID:   3,
		Name: "Auto",
	},
}

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r Repository) FindAllCategories(_ context.Context) (category.Categories, error) {
	return categories, nil
}
