package domain_repositories

import (
	"context"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	order_vo "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"
)

type OrderRepositoryFilter struct {
	Status order_vo.OrderStatus
}

type OrderRepository interface {
	Persist(context context.Context, order entities.Order) (*entities.Order, error)
	Detail(context context.Context, ID string) (*entities.Order, error)
	Delete(context context.Context, ID string) error
	List(context context.Context, page int, perPage int, filter *OrderRepositoryFilter) (*[]entities.Order, *int64, error)
}
