package services

import (
	"github.com/Project/webProject2019/entity"
	"github.com/Project/webProject2019/payments"
)

// RoomServiceImpl implements rooms.RoomService interface
type PaymentServiceImpl struct {
	paymentservices payments.PaymentServices
}

// NewPaymentServiceImpl will create new PaymentsService object
func NewPaymentServiceImpl(pay payments.PaymentServices) *PaymentServiceImpl {
	return &PaymentServiceImpl{paymentservices: pay}
}


// PaymentReservation persists new reservations information
func (pservice *PaymentServiceImpl) PaymentReservation(pay entity.Payment) error {

	err := pservice.paymentservices.PaymentReservation(pay)

	if err != nil {
		return err
	}

	return nil
}


func (pservice *PaymentServiceImpl) IncreamentRate(rate entity.Rating) error {

	err := pservice.paymentservices.IncreamentRate(rate)

	if err != nil {
		return err
	}

	return nil
}
func (pservice *PaymentServiceImpl) RetrieveHotelRateValue() (entity.Rating,error) {

	data,err := pservice.paymentservices.RetrieveHotelRateValue()

	if err != nil {
		return data,err
	}

	return data, nil
}


//func (pservice *PaymentServiceImpl) RetrieveAccountFromBank(v int) (entity.Balance,error) {
//
//	data,err := pservice.paymentservices.RetrieveAccountFromBank(v)
//
//	if err != nil {
//		return data,err
//	}
//
//	return data, nil
//}

