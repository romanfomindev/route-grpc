package product_service

import (
	"context"

	product_service "github.com/romanfomindev/route-grpc/tree/master/product-service/internal/service/product"
	desc "github.com/romanfomindev/route-grpc/tree/master/product-service/pkg/product-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {

	product, err := i.productService.CreateProduct(ctx, req.GetName(), int64(req.GetCategoryId()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return makeCreateProductResponse(product), nil
}

func makeCreateProductResponse(product *product_service.Product) *desc.CreateProductResponse {
	return &desc.CreateProductResponse{
		Result: &desc.Product{
			Name:       product.Name,
			Id:         uint64(product.ID),
			CategoryId: product.CategoryID,
		},
	}
}
