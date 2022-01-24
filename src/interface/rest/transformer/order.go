package transformer

import (
	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	valueobject "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"
)

type OrderTransform struct {
	ID        string                  `json:"id"`
	Items     []valueobject.OrderItem `json:"items"`
	Status    string                  `json:"status"`
	CreatedAt string                  `json:"createdAt"`
	UpdatedAt string                  `json:"updatedAt"`
}

type OrderTransformCreateUpdate struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type OrderTransformList struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func Transform(o *entities.Order) *OrderTransform {
	return &OrderTransform{
		ID:        o.GetID().String(),
		Status:    o.GetStatus().ToString(),
		Items:     o.GetItems(),
		CreatedAt: o.GetCreatedAt().String(),
		UpdatedAt: o.GetUpdatedAt().String(),
	}
}

func TransformCreateUpdate(o *entities.Order) *OrderTransformCreateUpdate {
	return &OrderTransformCreateUpdate{
		ID:        o.GetID().String(),
		CreatedAt: o.GetCreatedAt().String(),
		UpdatedAt: o.GetUpdatedAt().String(),
	}
}

func TransformList(o *[]entities.Order) *[]OrderTransformList {
	var orders []OrderTransformList
	if len(*o) == 0 {
		orders = make([]OrderTransformList, 0)
	}

	for _, order := range *o {
		or := &OrderTransformList{
			ID:        order.GetID().String(),
			Status:    order.GetStatus().ToString(),
			CreatedAt: order.GetCreatedAt().String(),
			UpdatedAt: order.GetUpdatedAt().String(),
		}

		orders = append(orders, *or)
	}

	return &orders
}
