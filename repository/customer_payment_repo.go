package repository

import (
	"go-wmb/model"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	Payment(newPayment model.Payment) ([]model.Payment, error)
}

type paymentRepo struct {
	db *gorm.DB
}

func (p *paymentRepo) Payment(newPayment model.Payment) ([]model.Payment, error) {
	invoice := make([]model.Payment, 0)
	err := p.db.Where("Table Number LIKE ?", newPayment.TableNumber).Find(&invoice).Error
	p.db.Model(&newPayment).Where("TableNumber = ?", newPayment.TableNumber).Update("IsReserved", true)
	if err != nil {
		panic(err)
	}
	return invoice, nil
}

func NewCustomerPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db: db}
}
