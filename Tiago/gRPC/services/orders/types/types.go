package types

type OrderService interface {
	CreateOrder(context.Context) error
}
