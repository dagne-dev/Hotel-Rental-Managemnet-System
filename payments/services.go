package payments

import "github.com/Project/webProject2019/entity"

type PaymentServices interface {
	PaymentReservation(pay entity.Payment)  error
	IncreamentRate(rate entity.Rating)  error
	RetrieveHotelRateValue() (entity.Rating,error)
        GetUserId(updaterate *entity.Rating) ([]entity.Rating,[]error)
	GetUserRateValue(user_id uint) (*entity.Rating,[]error)
	UpdateUserRateValue(user_id uint,u int) ([]error)  
	//RetrieveAccountFromBank(v int) (entity.Balance,error)

}
