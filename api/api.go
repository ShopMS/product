package api

import "github.com/shopms/product/internal/service"

type ProductApi struct {
	ProductService service.ProductService
}

func NewProductApi(productService service.ProductService) *ProductApi {
	return &ProductApi{
		ProductService: productService,
	}
}
