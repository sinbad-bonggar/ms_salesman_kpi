package infra_model

import (
	"encoding/json"
	"time"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	valueobject "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/value_objects/order"

	"github.com/google/uuid"
)

type Order struct {
	ID        string    `json:"id" gorm:"column:id;primaryKey"`
	Items     string    `json:"items" gorm:"column:items"`
	Status    uint8     `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (o Order) ToEntity() (*entities.Order, error) {
	parsedUUID, err := uuid.Parse(o.ID)
	if err != nil {
		return nil, err
	}

	// since value object has unexported fields,
	// we need to unmarshall back to array of struct
	out := []struct {
		Name  string
		Qty   uint8
		Price float32
	}{}
	if err = json.Unmarshal([]byte(o.Items), &out); err != nil {
		return nil, err
	}

	items := []valueobject.OrderItem{}
	for _, v := range out {
		item := valueobject.CreateOrderItem(v.Name, v.Qty, v.Price)
		items = append(items, item)
	}

	status := valueobject.CreateOrderStatus(o.Status)
	order, _ := entities.MakeOrder(parsedUUID, items, status, o.CreatedAt, o.UpdatedAt)

	return &order, nil
}
