package models

import (
	"log"
)

func (order *Order) AddOrder() error {
	err := DBInstance().Create(order).Error
	return err
}

func (order *Order) FetchAllOrders(userId uint) ([]Order, error) {
	var orders []Order
	err := DBInstance().Where("user_id=?", userId).Find(&orders).Error
	return orders, err
}

func (order *Order) RemoveOrder(orderId uint) error {
	removeOrder := &Order{}
	err := DBInstance().Where("id=?", orderId).Find(removeOrder).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return DBInstance().Delete(removeOrder).Error
}

func (order *Order) RemoveAllOrders(userId uint) error {
	var orders []Order
	return DBInstance().Where("user_id=?", userId).Find(&orders).Delete(&orders).Error
}