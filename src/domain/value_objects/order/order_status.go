package domain_value_object

type OrderStatus uint8

const (
	StatusDraft uint8 = iota + 1
	StatusDone
)

func CreateOrderStatus(s uint8) OrderStatus {
	return OrderStatus(s)
}

func (w OrderStatus) ToString() string {
	return [...]string{"Draft", "Done"}[w-1]
}

func (w OrderStatus) EnumIndex() uint8 {
	return uint8(w)
}
