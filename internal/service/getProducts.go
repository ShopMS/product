package service

import (
	"context"

	"github.com/shopms/common/logger"
	"github.com/shopms/product/internal/domain"
	pb "github.com/shopms/product/rpc/product"
)

func (p *productServiceImpl) GetProducts(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	result, err := p.ProductRepository.GetProducts(ctx, toProductQuery(request))

	if err != nil {
		logger.Debugf("1")
		return nil, err
	}

	var items []*pb.Product

	for _, v := range result.Items {
		items = append(items, domain.GetProto(v))
	}

	return &pb.GetProductsResponse{
		Products: items,
	}, nil
}

func toProductQuery(productRequest *pb.GetProductsRequest) *domain.ProductQuery {
	return &domain.ProductQuery{
		Limit:  productRequest.Limit,
		Offset: productRequest.Offset,
	}
}
