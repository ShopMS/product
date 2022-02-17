package repository

import (
	"context"
	"fmt"

	"github.com/shopms/product/ent"
	"github.com/shopms/product/internal/domain"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, query *domain.ProductQuery) (*domain.ProductPage, error)
	Create(ctx context.Context, products []*ent.Product) ([]*ent.Product, error)
}

type productRepositoryImpl struct {
	tx *ent.Client
}

func NewProductRepository(db *ent.Client) ProductRepository {
	return &productRepositoryImpl{
		tx: db,
	}
}

func (p *productRepositoryImpl) GetProducts(ctx context.Context, params *domain.ProductQuery) (*domain.ProductPage, error) {
	products, err := p.tx.Product.Query().All(context.Background())

	if err != nil {
		//return nil, err
		fmt.Println(err)
	}

	return &domain.ProductPage{
		Total: int64(len(products)),
		Items: products,
	}, nil
}

func (p *productRepositoryImpl) Create(ctx context.Context, products []*ent.Product) ([]*ent.Product, error) {
	bulk := make([]*ent.ProductCreate, len(products))

	for i, product := range products {
		bulk[i] = p.tx.Product.Create().SetName(product.Name).
			SetDescription(product.Description).
			SetQuantity(product.Quantity).
			SetPrice(product.Price)
	}
	result, err := p.tx.Product.CreateBulk(bulk...).Save(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
