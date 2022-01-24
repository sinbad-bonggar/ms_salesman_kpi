package handler

import (
	"errors"
	"net/http"

	dto "github.com/sinbad-bonggar/ms_salesman_kpi/src/app/data_transfer_object"
	usecases "github.com/sinbad-bonggar/ms_salesman_kpi/src/app/use_cases"
	common_error "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/errors"
	request "github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/request"
	response "github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/response"
	transformer "github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/transformer"
)

// IOrderHandler ...
type IOrderHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	response response.IResponseClient
	usecase  usecases.IOrderUseCase
}

// NewOrderHandler ...
func NewOrderHandler(r response.IResponseClient, u usecases.IOrderUseCase) IOrderHandler {
	return &orderHandler{
		response: r,
		usecase:  u,
	}
}

// List ...
func (h *orderHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := request.OrderList{}
	dto, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, err)
		return
	}

	if orders, count, err := h.usecase.List(ctx, dto); err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_ORDER, err))
		return
	} else {
		h.response.JSON(
			w,
			"",
			transformer.TransformList(orders),
			h.response.BuildMeta(req.Page, req.PerPage, *count),
		)
	}
}

// Detail ...
func (h *orderHandler) Detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := request.OrderDetail{}
	dto, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.INVALID_ORDER, err))
		return
	}

	order, err := h.usecase.Detail(ctx, dto)
	if order == nil && err == nil {
		h.response.HttpError(w, common_error.NewError(common_error.ORDER_NOT_FOUND, errors.New("record not found")))
		return
	}
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_ORDER, err))
		return
	}

	h.response.JSON(w, "", transformer.Transform(order), nil)
}

// Create ...
func (h *orderHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := request.OrderCreate{}
	dto, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, err)
		return
	}

	if order, err := h.usecase.Create(ctx, dto); err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_ORDER, err))
		return
	} else {
		if order == nil {
			h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_ORDER, nil))
			return
		}

		h.response.JSON(w, "Successfully created data", transformer.TransformCreateUpdate(order), nil)
	}
}

// Update ...
func (h *orderHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := request.OrderUpdate{}
	orderUpdateDTO, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_UPDATE_ORDER, err))
		return
	}

	// because the ID we got from URL
	// we need to return 404 instead return 500
	orderDetailDTO := dto.NewOrderDetailDTO(req.ID)
	order, err := h.usecase.Detail(ctx, orderDetailDTO)
	if order == nil && err == nil {
		h.response.HttpError(w, common_error.NewError(common_error.ORDER_NOT_FOUND, errors.New("record not found")))
		return
	}
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_UPDATE_ORDER, err))
		return
	}

	if order, err := h.usecase.Update(ctx, orderUpdateDTO); err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_UPDATE_ORDER, err))
		return
	} else {
		h.response.JSON(w, "Successfully updated data", transformer.TransformCreateUpdate(order), nil)
	}
}

// Delete ...
func (h *orderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := request.OrderDelete{}
	orderDeleteDTO, err := req.Validate(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.INVALID_ORDER, err))
		return
	}

	orderDetailDTO := dto.NewOrderDetailDTO(req.ID)
	order, err := h.usecase.Detail(ctx, orderDetailDTO)
	if order == nil && err == nil {
		h.response.HttpError(w, common_error.NewError(common_error.ORDER_NOT_FOUND, errors.New("record not found")))
		return
	}
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_DELETE_ORDER, err))
		return
	}

	if err := h.usecase.Delete(ctx, orderDeleteDTO); err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_DELETE_ORDER, err))
		return
	} else {
		h.response.JSON(w, "Successfully deleted data", nil, nil)
	}
}
