package payments

import "github.com/Project/webProject2019/entity"

type PaymentRepository interface {
	PaymentReservation(pay entity.Payment)  error
	IncreamentRate(rate entity.Rating)  error
	RetrieveHotelRateValue() (entity.Rating,error)
	//RetrieveAccountFromBank(v int) (entity.Balance,error)


}