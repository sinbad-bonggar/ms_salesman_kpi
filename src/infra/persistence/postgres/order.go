package postgres

import (
	"context"
	"encoding/json"
	"errors"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	repositories "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/repositories"
	infra_model "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepository{
		connection: db,
	}
}

func ToModel(order entities.Order) (*infra_model.Order, error) {
	// since value object has unexported fields,
	// we need to unmarshall back to struct
	list := []map[string]interface{}{}
	for _, v := range order.GetItems() {
		m, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}

		out := map[string]interface{}{}
		err = json.Unmarshal([]byte(m), &out)
		if err != nil {
			return nil, err
		}

		list = append(list, out)
	}

	items, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return &infra_model.Order{
		ID:        order.GetID().String(),
		Items:     string(items),
		Status:    order.GetStatus().EnumIndex(),
		CreatedAt: order.GetCreatedAt(),
		UpdatedAt: order.GetUpdatedAt(),
	}, nil
}

func (repo *orderRepository) Persist(ctx context.Context, order entities.Order) (*entities.Order, error) {
	tx := repo.connection.WithContext(ctx).Begin()

	model, err := ToModel(order)
	if err != nil {
		return nil, err
	}

	repo.connection.Save(model)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &order, nil
}

func (repo *orderRepository) Detail(ctx context.Context, ID string) (*entities.Order, error) {
	var model infra_model.Order

	err := repo.connection.WithContext(ctx).First(&model, "id = ?", ID).Error
	// check error ErrRecordNotFound, then return nil instead err
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	order, err := model.ToEntity()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (repo *orderRepository) Delete(ctx context.Context, ID string) error {
	var model infra_model.Order

	err := repo.connection.WithContext(ctx).Delete(&model, "id", ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) List(ctx context.Context, page int, perPage int, filter *repositories.OrderRepositoryFilter) (*[]entities.Order, *int64, error) {
	var models []*infra_model.Order
	var count int64
	var orderEntities []entities.Order

	q := repo.connection.WithContext(ctx).Where(filter)
	if err := q.Find(&models).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	if err := q.Scopes(Paginate(page, perPage)).Find(&models).Error; err != nil {
		return nil, nil, err
	}

	for _, row := range models {
		order, _ := row.ToEntity()

		orderEntities = append(orderEntities, *order)
	}

	return &orderEntities, &count, nil
}
