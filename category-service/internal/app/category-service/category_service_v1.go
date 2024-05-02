package category_service

import (
	"context"

	errors "github.com/pkg/errors"
	internal_errors "github.com/romanfomindev/route-grpc/tree/master/category-service/internal/pkg/errors"
	"github.com/romanfomindev/route-grpc/tree/master/category-service/internal/service/category"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/romanfomindev/route-grpc/tree/master/category-service/pkg/category-service"
)

func (i *Implementation) GetCategoryById(
	ctx context.Context,
	req *desc.GetCategoryByIdRequest,
) (*desc.GetCategoryByIdResponse, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("GetCategoryById - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	findCategory, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, internal_errors.ErrNoCategory) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, errors.Wrap(err, "categoryService.GetCategoryByID")
	}

	return makeGetCategoryByIdResponse(findCategory), nil
}

func makeGetCategoryByIdResponse(cat *category.Category) *desc.GetCategoryByIdResponse {
	return &desc.GetCategoryByIdResponse{
		Category: &desc.Category{
			Id:   cat.ID,
			Name: cat.Name,
		},
	}
}
