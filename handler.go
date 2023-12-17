package handler

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"GO-LANG/datastore"
	"GO-LANG/model"
)

type handler struct {
	store datastore.Inventory
}

func New(s datastore.Inventory) handler {
	return handler{store: s}
}

func validateProductID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (h handler) GetProductByID(ctx *gofr.Context) (interface{}, error) {
	ID := ctx.PathParam("ID")
	if ID == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateProductID(ID); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "product",
			ID:     ID,
		}
	}

	return resp, nil
}

func (h handler) AddProduct(ctx *gofr.Context) (interface{}, error) {
	var product model.Product

	if err := ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.AddProduct(ctx, &product)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) UpdateProduct(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateProductID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var product model.Product
	if err = ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	product.ID = id

	resp, err := h.store.UpdateProduct(ctx, &product)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) RemoveProduct(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateProductID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.RemoveProduct(ctx, id); err != nil {
		return nil, err
	}

	return "Product deleted successfully", nil
}
