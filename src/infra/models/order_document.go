package infra_model

import (
	"time"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	order_vo "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"

	"github.com/google/uuid"
)

type OrderDocument struct {
	ID        string              `bson:"id"`
	Items     []OrderItemDocument `bson:"items"`
	Status    uint8               `bson:"status"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
}

type OrderItemDocument struct {
	Name  string  `bson:"name"`
	Qty   uint8   `bson:"qty"`
	Price float32 `bson:"price"`
	Total float32 `bson:"total"`
}

func (o OrderDocument) ToEntity() (*entities.Order, error) {
	var itemsVo []order_vo.OrderItem
	var statusVo order_vo.OrderStatus

	parsed, err := uuid.Parse(o.ID)
	if err != nil {
		return &entities.Order{}, err
	}

	for _, item := range o.Items {
		or, _ := item.ToEntity()
		itemsVo = append(itemsVo, or)
	}

	statusVo = order_vo.CreateOrderStatus(o.Status)

	order, err := entities.MakeOrder(parsed, itemsVo, statusVo, o.CreatedAt, o.UpdatedAt)
	if err != nil {
		return &entities.Order{}, err
	}

	return &order, nil
}

func (o OrderItemDocument) ToEntity() (order_vo.OrderItem, error) {
	items := order_vo.CreateOrderItem(o.Name, o.Qty, o.Price)

	return items, nil
}
