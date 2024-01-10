package ports

import "github.com/ForbiddenR/micros/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
