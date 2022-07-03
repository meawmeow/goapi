package repository

import (
	"fiberapiv1/entity"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	InsertOrders(entity.Orders) (err error)
	Product(id int) (product *entity.Product, err error)
	OrdersAll() (orders []*entity.Orders, err error)
	DeleteOrdersById(id int) (err error)
	OrdersByUserId(id int) (orders []*entity.Orders, err error)
}

type ordersRepositoryDB struct {
	db *gorm.DB
}

func NewOrdersRepositoryDB(db *gorm.DB) OrdersRepository {
	return ordersRepositoryDB{db}
}
func (r ordersRepositoryDB) InsertOrders(orders entity.Orders) (err error) {
	err = r.db.Save(&orders).Error
	return err
}

func (r ordersRepositoryDB) Product(id int) (product *entity.Product, err error) {
	err = r.db.Where("id=?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r ordersRepositoryDB) OrdersAll() (orders []*entity.Orders, err error) {
	err = r.db.Preload("Product.ProductGroup").Preload("User").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r ordersRepositoryDB) OrdersByUserId(id int) (orders []*entity.Orders, err error) {
	err = r.db.Where("user_id=?", id).Preload("Product.ProductGroup").Preload("User").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r ordersRepositoryDB) DeleteOrdersById(id int) (err error) {
	err = r.db.Delete(&entity.Orders{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
