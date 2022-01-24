package usecases

import (
	"context"
	"errors"

	dto "github.com/sinbad-bonggar/ms_salesman_kpi/src/app/data_transfer_object"
	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	repositories "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/repositories"
	valueobjects "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"
	infra_error "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/errors"
)

type IOrderUseCase interface {
	List(ctx context.Context, do dto.IOrderDTO) (*[]entities.Order, *int64, error)
	Detail(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error)
	Create(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error)
	Update(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error)
	Delete(ctx context.Context, do dto.IOrderDTO) error
}

type orderUseCase struct {
	repository repositories.OrderRepository
}

func NewOrderUseCase(or repositories.OrderRepository) IOrderUseCase {
	return &orderUseCase{
		repository: or,
	}
}

func (u *orderUseCase) List(ctx context.Context, do dto.IOrderDTO) (*[]entities.Order, *int64, error) {
	orderList, ok := do.(*dto.OrderList)
	if !ok {
		return nil, nil, errors.New("type assertion failed to dto.OrderList")
	}

	filter := repositories.OrderRepositoryFilter{
		Status: valueobjects.OrderStatus(orderList.Status),
	}

	result, count, err := u.repository.List(ctx, orderList.Page, orderList.PerPage, &filter)
	if err != nil {
		return nil, nil, infra_error.NewError(infra_error.FAILED_RETRIEVE_ORDER, err)
	}

	return result, count, nil
}

func (u *orderUseCase) Detail(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error) {
	orderDetail, ok := do.(*dto.OrderDetail)
	if !ok {
		return nil, errors.New("type assertion failed to dto.OrderDetail")
	}

	result, err := u.repository.Detail(ctx, orderDetail.ID)
	if err != nil {
		return nil, infra_error.NewError(infra_error.FAILED_RETRIEVE_ORDER, err)
	}

	return result, nil
}

func (u *orderUseCase) Create(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error) {
	orderCreate, ok := do.(*dto.OrderCreate)
	if !ok {
		return nil, errors.New("type assertion failed to dto.OrderCreate")
	}

	var items []valueobjects.OrderItem
	for _, item := range orderCreate.Items {
		item := valueobjects.CreateOrderItem(item.Name, uint8(item.Qty), item.Price)
		items = append(items, item)
	}

	status := valueobjects.CreateOrderStatus(1)

	order, err := entities.CreateOrder(items, status)
	if err != nil {
		return nil, infra_error.NewError(infra_error.FAILED_CREATE_ORDER, err)
	}

	result, err := u.repository.Persist(ctx, order)
	if err != nil {
		return nil, infra_error.NewError(infra_error.FAILED_CREATE_ORDER, err)
	}
	// TODO: trigger event to message broker

	return result, nil
}

func (u *orderUseCase) Update(ctx context.Context, do dto.IOrderDTO) (*entities.Order, error) {
	orderUpdate, ok := do.(*dto.OrderUpdate)
	if !ok {
		return nil, errors.New("type assertion failed to dto.OrderUpdate")
	}

	order, err := u.repository.Detail(ctx, orderUpdate.ID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("record not found")
	}

	var items []valueobjects.OrderItem
	for _, item := range orderUpdate.Items {
		item := valueobjects.CreateOrderItem(item.Name, uint8(item.Qty), item.Price)
		items = append(items, item)
	}

	order.UpdateStatus(valueobjects.OrderStatus(orderUpdate.Status))
	order.UpdateItems(items)

	result, err := u.repository.Persist(ctx, *order)
	if err != nil {
		return nil, err
	}

	// TODO: trigger event to message broker

	return result, nil
}

func (u *orderUseCase) Delete(ctx context.Context, do dto.IOrderDTO) error {
	orderDelete, ok := do.(*dto.OrderDelete)
	if !ok {
		return errors.New("type assertion failed to dto.OrderDelete")
	}

	err := u.repository.Delete(ctx, orderDelete.ID)
	if err != nil {
		return err
	}

	// TODO: trigger event to message broker

	return nil
}
