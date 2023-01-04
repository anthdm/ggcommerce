package store

import (
	"context"

	"github.com/anthdm/ggcommerce/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetByID(context.Context, string) (*types.Product, error)
	GetAll(context.Context) ([]*types.Product, error)
}
