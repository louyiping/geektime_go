package service

import (
	v1 "github.com/louyiping/geektime_go/week4/api/payment/service/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/louyiping/geektime_go/week4/app/payment/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPaymentService)

type PaymentService struct {
	v1.UnimplementedPaymentServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewPaymentService(bc *biz.BeerUseCase, logger log.Logger) *PaymentService {
	return &PaymentService{

		bc:  bc,
		log: log.NewHelper("service/payment", logger)}
}