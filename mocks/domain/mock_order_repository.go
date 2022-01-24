package mock_domain_repositories

import (
	"context"

	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"
	repos "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/repositories"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{}
}

// implement interface
var _ repos.OrderRepository = &MockOrderRepository{}

func (m *MockOrderRepository) Persist(context context.Context, order entities.Order) (*entities.Order, error) {
	args := m.Called(context, order)
	var or *entities.Order
	var err error

	if n, ok := args.Get(0).(*entities.Order); ok {
		or = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return or, err
}

func (m *MockOrderRepository) Detail(context context.Context, ID string) (*entities.Order, error) {
	args := m.Called(context, ID)
	var or *entities.Order
	var err error

	if n, ok := args.Get(0).(*entities.Order); ok {
		or = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return or, err
}

func (m *MockOrderRepository) Delete(context context.Context, ID string) error {
	args := m.Called(context, ID)
	if err, ok := args.Get(0).(error); ok {
		return err
	}

	return nil
}

func (m *MockOrderRepository) List(context context.Context, page int, perPage int, filter *repos.OrderRepositoryFilter) (*[]entities.Order, *int64, error) {
	args := m.Called(context, page, perPage, filter)
	var or *[]entities.Order
	var count *int64
	var err error

	if n, ok := args.Get(0).(*[]entities.Order); ok {
		or = n
	}

	if n, ok := args.Get(1).(*int64); ok {
		count = n
	}

	if n, ok := args.Get(2).(error); ok {
		err = n
	}

	return or, count, err
}
