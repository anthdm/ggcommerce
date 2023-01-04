package api

import (
	"encoding/json"
	"net/http"

	"github.com/anthdm/ggcommerce/store"
	"github.com/anthdm/ggcommerce/types"

	"github.com/anthdm/weavebox"
)

type ProductHandler struct {
	store store.ProductStorer
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct(c *weavebox.Context) error {
	productReq := &types.CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}

	product, err := types.NewProductFromRequest(productReq)
	if err != nil {
		return err
	}

	if err := h.store.Insert(c.Context, product); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) HandleGetProducts(c *weavebox.Context) error {
	products, err := h.store.GetAll(c.Context)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) HandleGetProductByID(c *weavebox.Context) error {
	id := c.Param("id")

	product, err := h.store.GetByID(c.Context, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}
