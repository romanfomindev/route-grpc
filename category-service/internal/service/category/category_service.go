package category

import (
	"context"

	"github.com/pkg/errors"
	internal_errors "github.com/romanfomindev/route-grpc/tree/master/category-service/internal/pkg/errors"
)

type Service struct {
	repository CategoryRepository
}

type CategoryRepository interface {
	FindAllCategories(context.Context) (Categories, error)
}

func New(repository CategoryRepository) *Service {
	return &Service{
		repository: repository,
	}
}

var ErrNoCategory = errors.Wrap(internal_errors.ErrNoCategory, "no category")

func (s *Service) GetCategoryByID(ctx context.Context, id uint64) (*Category, error) {
	categories, err := s.repository.FindAllCategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.FindAllCategories")
	}

	category := categories.FilterByID(id)
	if category == nil {
		return nil, ErrNoCategory
	}

	return category, nil
}
