package types

import "fmt"

const minProductNameLen = 3

type Product struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}

type CreateProductRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductFromRequest(req *CreateProductRequest) (*Product, error) {
	if err := validateCreateProductRequest(req); err != nil {
		return nil, err
	}
	return &Product{
		SKU:  req.SKU,
		Name: req.Name,
	}, nil
}

func validateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product is to short")
	}

	return nil
}
