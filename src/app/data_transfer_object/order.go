package data_transfer_object

import (
	validation "github.com/go-ozzo/ozzo-validation"
	valueobject "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"
	infra_error "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/errors"
)

type IOrderDTO interface {
	Validate() error
}

type Order struct {
	Name  string
	Qty   float32
	Price float32
}

type OrderList struct {
	Page    int
	PerPage int
	Status  uint8
}

func NewOrderListDTO(page, perPage int, status uint8) IOrderDTO {
	return &OrderList{
		Page:    page,
		PerPage: perPage,
		Status:  status,
	}
}

func (dto *OrderList) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Page),
		validation.Field(&dto.PerPage),
		validation.Field(&dto.Status, validation.In(valueobject.StatusDraft, valueobject.StatusDone)),
	); err != nil {
		retErr := infra_error.NewError(infra_error.INVALID_REQUEST_RETRIEVE_ORDER, err)
		retErr.SetValidationMessage(err)

		return retErr
	}

	return nil
}

type OrderDetail struct {
	ID string
}

func NewOrderDetailDTO(id string) IOrderDTO {
	return &OrderDetail{
		ID: id,
	}
}

func (dto *OrderDetail) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		retErr := infra_error.NewError(infra_error.INVALID_REQUEST_RETRIEVE_ORDER, err)
		retErr.SetValidationMessage(err)

		return retErr
	}
	return nil
}

type OrderCreate struct {
	Items []Order
}

func NewOrderCreateDTO(items []Order) IOrderDTO {
	return &OrderCreate{
		Items: items,
	}
}

func (dto *OrderCreate) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Items, validation.Required),
	); err != nil {
		retErr := infra_error.NewError(infra_error.INVALID_REQUEST_CREATE_ORDER, err)
		retErr.SetValidationMessage(err)
		return retErr
	}
	return nil
}

type OrderUpdate struct {
	ID     string
	Status uint8
	Items  []Order
}

func NewOrderUpdateDTO(id string, status uint8, items []Order) IOrderDTO {
	return &OrderUpdate{
		ID:     id,
		Status: status,
		Items:  items,
	}
}

func (dto *OrderUpdate) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
		validation.Field(&dto.Status, validation.Required, validation.In(valueobject.StatusDraft, valueobject.StatusDone)),
		validation.Field(&dto.Items, validation.Required),
	); err != nil {
		retErr := infra_error.NewError(infra_error.INVALID_REQUEST_UPDATE_ORDER, err)
		retErr.SetValidationMessage(err)

		return retErr
	}
	return nil
}

type OrderDelete struct {
	ID string
}

func NewOrderDeleteDTO(id string) IOrderDTO {
	return &OrderDelete{
		ID: id,
	}
}

func (dto *OrderDelete) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		retErr := infra_error.NewError(infra_error.INVALID_REQUEST_DELETE_ORDER, err)
		retErr.SetValidationMessage(err)

		return retErr
	}
	return nil
}
