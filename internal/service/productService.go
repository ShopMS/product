package service

import (
	"context"

	repo "github.com/shopms/product/internal/repository"
	pb "github.com/shopms/product/rpc/product"
)

type ProductService interface {
	GetProducts(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error)
}

type productServiceImpl struct {
	ProductRepository repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: repository,
	}
}
