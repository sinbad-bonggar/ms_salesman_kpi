package domain_models

import (
	"time"

	valueobject "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"

	"github.com/google/uuid"
)

type Order struct {
	id        uuid.UUID
	items     []valueobject.OrderItem
	status    valueobject.OrderStatus
	createdAt time.Time
	updatedAt time.Time
}

func (m *Order) GetID() uuid.UUID {
	return m.id
}

func (m *Order) GetItems() []valueobject.OrderItem {
	return m.items
}

func (m *Order) GetStatus() valueobject.OrderStatus {
	return m.status
}

func (m *Order) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *Order) GetUpdatedAt() time.Time {
	return m.updatedAt
}

func (m *Order) UpdateStatus(status valueobject.OrderStatus) {
	m.status = status
}

func (m *Order) UpdateItems(items []valueobject.OrderItem) {
	m.items = items
}

func CreateOrder(i []valueobject.OrderItem, s valueobject.OrderStatus) (Order, error) {
	return Order{
		id:        uuid.New(),
		items:     i,
		status:    s,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func MakeOrder(
	id uuid.UUID,
	items []valueobject.OrderItem,
	status valueobject.OrderStatus,
	createdAt time.Time,
	updatedAt time.Time,
) (Order, error) {
	return Order{
		id:        id,
		items:     items,
		status:    status,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
