package mock_app_use_cases

import (
	"context"

	"github.com/sinbad-bonggar/ms_salesman_kpi/src/app/data_transfer_object"
	entities "github.com/sinbad-bonggar/ms_salesman_kpi/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockOrderDetailUseCase struct {
	mock.Mock
}

func (m *MockOrderDetailUseCase) OrderDetail(context context.Context, dto data_transfer_object.IOrderDTO) (*entities.Order, error) {
	args := m.Called(dto)
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
