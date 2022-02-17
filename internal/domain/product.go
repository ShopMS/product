package domain

import (
	"github.com/shopms/product/ent"
	pb "github.com/shopms/product/rpc/product"
)

const (
	ProductTable = "Product"
)

func NewProduct(product *pb.Product) *ent.Product {
	if product == nil {
		return nil
	}
	return &ent.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       float64(product.Price),
		Quantity:    int(product.Quantity),
	}
}

func GetProto(p *ent.Product) *pb.Product {
	if p == nil {
		return nil
	}
	return &pb.Product{
		Id:          int64(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Quantity:    int64(p.Quantity),
		Price:       float32(p.Price),
	}
}

type ProductPage struct {
	Total int64
	Items []*ent.Product
}

type ProductQuery struct {
	Sort, Order   string
	Offset, Limit int64
}
