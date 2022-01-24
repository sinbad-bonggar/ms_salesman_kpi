package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	commonError "github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/errors"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/response"
)

const (
	XSellerId string = "x-seller-id"
	XBuyerId  string = "x-buyer-id"
	XUserId   string = "x-user-id"
)

type sinbadContextKey int

const (
	CtxSinbadHeader sinbadContextKey = iota + 1
)

type ContexSinbadHeader struct {
	SellerId *uint64
	BuyerId  *uint64
	UserId   *uint64
}

func CheckSinbadAppHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		buyerId := r.Header.Get(XBuyerId)
		userId := r.Header.Get(XUserId)

		if buyerId == "" {
			err := errors.New("buyerId should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_BUYER_ID, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		if userId == "" {
			err := errors.New("userId should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		bi, err := strconv.ParseUint(buyerId, 10, 32)
		if err != nil {
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_BUYER_ID, err)

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		ui, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		val := ContexSinbadHeader{
			BuyerId: &bi,
			UserId:  &ui,
		}

		ctx = context.WithValue(ctx, CtxSinbadHeader, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func CheckAgentAppHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId := r.Header.Get(XUserId)

		if userId == "" {
			err := errors.New("userId should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		ui, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		val := ContexSinbadHeader{
			UserId: &ui,
		}

		ctx = context.WithValue(ctx, CtxSinbadHeader, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func CheckSSCWebHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sellerId := r.Header.Get(XSellerId)
		userId := r.Header.Get(XUserId)

		if sellerId == "" {
			err := errors.New("sellerId should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_SELLER_ID, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		if userId == "" {
			err := errors.New("userId should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		si, err := strconv.ParseUint(sellerId, 10, 32)
		if err != nil {
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_SELLER_ID, err)

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		ui, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_USER_ID, err)

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		val := ContexSinbadHeader{
			SellerId: &si,
			UserId:   &ui,
		}

		ctx = context.WithValue(ctx, CtxSinbadHeader, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func CheckAPWebHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//	still not defined
	}

	return http.HandlerFunc(fn)
}
