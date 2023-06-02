package types

import (
	"fmt"
	"strings"
)

const minProductNameLen = 3

const minProductPriceLen = 1

type Product struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	SKU   string  `bson:"sku" json:"sku"`
	Name  string  `bson:"name" json:"name"`
	Slug  string  `bson:"slug" json:"slug"`
	Price float64 `bson:"price" json:"price"`
}

type CreateProductRequest struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProductFromRequest(req *CreateProductRequest) (*Product, error) {
	if err := validateCreateProductRequest(req); err != nil {
		return nil, err
	}

	parts := strings.Split(strings.ToLower(req.Name), " ")
	slug := strings.Join(parts, "-")

	return &Product{
		SKU:   req.SKU,
		Name:  req.Name,
		Slug:  slug,
		Price: req.Price,
	}, nil
}

func validateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.SKU) < 3 {
		return fmt.Errorf("the SKU of the product is to short")
	}
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product is to short")
	}
	if (req.Price) < 1 {
		return fmt.Errorf("the Price of the product is too low")
	}
	return nil
}
