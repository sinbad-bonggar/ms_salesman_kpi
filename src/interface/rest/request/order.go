package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	dto "github.com/sinbad-bonggar/ms_salesman_kpi/src/app/data_transfer_object"
	common_error "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/errors"
)

// IOrderRequest ...
type IOrderRequest interface {
	Validate(r *http.Request) (dto.IOrderDTO, error)
}

type Order struct {
	Name  string
	Qty   float32
	Price float32
}

// List
type OrderList struct {
	Page    int
	PerPage int
	Status  uint8
}

func (req *OrderList) Validate(r *http.Request) (dto.IOrderDTO, error) {
	if page, err := strconv.Atoi(r.URL.Query().Get("page")); err != nil {
		req.Page = 1
	} else {
		req.Page = page
	}

	if perPage, err := strconv.Atoi(r.URL.Query().Get("perPage")); err != nil {
		req.PerPage = 10
	} else {
		req.PerPage = perPage
	}

	if status, err := strconv.Atoi(r.URL.Query().Get("status")); err != nil {
		req.Status = uint8(0)
	} else {
		req.Status = uint8(status)
	}

	d := dto.NewOrderListDTO(req.Page, req.PerPage, req.Status)
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d, nil
}

// Detail
type OrderDetail struct {
	ID string
}

func (req *OrderDetail) Validate(r *http.Request) (dto.IOrderDTO, error) {
	req.ID = chi.URLParam(r, "id")

	d := dto.NewOrderDetailDTO(req.ID)
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d, nil
}

// Create
type OrderCreate struct {
	Items []Order
}

func (req *OrderCreate) Validate(r *http.Request) (dto.IOrderDTO, error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg := err.Error()
		if strings.Contains(msg, "request.OrderCreate") {
			msg = "payload body must be an object"
		} else if strings.Contains(msg, "OrderCreate.Items") {
			msg = "items must be array of object"
		} else {
			msg = "invalid JSON format"
		}

		return nil, common_error.NewError(common_error.INVALID_PAYLOAD_CREATE_ORDER, fmt.Errorf(msg))
	}

	var items []dto.Order

	for _, item := range req.Items {
		items = append(items, dto.Order{
			Name:  item.Name,
			Qty:   item.Qty,
			Price: item.Price,
		})
	}

	d := dto.NewOrderCreateDTO(items)
	if err := d.Validate(); err != nil {
		return nil, err
	}

	return d, nil
}

// Update
type OrderUpdate struct {
	ID     string
	Status uint8
	Items  []Order
}

func (req *OrderUpdate) Validate(r *http.Request) (dto.IOrderDTO, error) {
	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg := err.Error()
		if strings.Contains(msg, "request.OrderUpdate") {
			msg = "payload body must be an object"
		} else if strings.Contains(msg, "OrderUpdate.Status") {
			msg = "status must be integer"
		} else if strings.Contains(msg, "OrderUpdate.Items") {
			msg = "items must be array of object"
		} else {
			msg = "invalid JSON format"
		}

		return nil, common_error.NewError(common_error.INVALID_PAYLOAD_UPDATE_ORDER, fmt.Errorf(msg))
	}

	var items []dto.Order
	for _, item := range req.Items {
		items = append(items, dto.Order{
			Name:  item.Name,
			Qty:   item.Qty,
			Price: item.Price,
		})
	}

	d := dto.NewOrderUpdateDTO(req.ID, req.Status, items)
	if err := d.Validate(); err != nil {
		return nil, err
	}

	return d, nil
}

// Delete
type OrderDelete struct {
	ID string
}

func (req *OrderDelete) Validate(r *http.Request) (dto.IOrderDTO, error) {
	req.ID = chi.URLParam(r, "id")

	d := dto.NewOrderDeleteDTO(req.ID)
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d, nil
}
