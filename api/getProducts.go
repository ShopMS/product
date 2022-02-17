package api

import (
	"context"
	"fmt"

	pb "github.com/shopms/product/rpc/product"
)

func (p *ProductApi) GetProducts(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	result, err := p.ProductService.GetProducts(ctx, request)
	if err != nil {
		fmt.Printf("Cound not get products on error %v", err)
		return nil, err
	}
	return result, nil
}
