package services

import (
	"fiberapiv1/entity"
	"fiberapiv1/errs"
	"fiberapiv1/logs"
	"fiberapiv1/models"
	"fiberapiv1/repository"
	"strconv"
)

type OrdersService interface {
	CreateOrders(models.OrdersRequest) (err error)
	OrdersAll() ([]models.OrdersResponse, error)
	OrdersByUserId(int) ([]models.OrdersResponse, error)
	DeleteOrdersById(int) error
}

type ordersService struct {
	ordersRepo repository.OrdersRepository
}

func NewOrdersService(ordersRepo repository.OrdersRepository) OrdersService {
	return ordersService{ordersRepo: ordersRepo}
}

func (r ordersService) CreateOrders(request models.OrdersRequest) (err error) {
	product, err := r.ordersRepo.Product(request.ProductID)
	if err != nil {
		return errs.NewError("productby id record not found")
	}
	price, err := strconv.Atoi(product.Price)
	if err != nil {
		return err
	}
	netPrice := (price * request.Count)

	orders := entity.Orders{
		Count:     request.Count,
		NetPrices: uint64(netPrice),
		ProductID: uint64(request.ProductID),
		UserID:    uint64(request.UserID),
	}
	//fmt.Println("Orders Result : ", orders)
	err = r.ordersRepo.InsertOrders(orders)
	return err
}

func (r ordersService) OrdersAll() ([]models.OrdersResponse, error) {

	orders, err := r.ordersRepo.OrdersAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	ordersResponses := []models.OrdersResponse{}
	for _, item := range orders {
		ordersResponse := models.OrdersResponse{
			ID:        int(item.ID),
			Count:     item.Count,
			NetPrices: item.NetPrices,
			Product:   item.Product,
			User:      item.User,
		}
		ordersResponses = append(ordersResponses, ordersResponse)
	}

	return ordersResponses, nil
}

func (r ordersService) OrdersByUserId(id int) ([]models.OrdersResponse, error) {

	orders, err := r.ordersRepo.OrdersByUserId(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	ordersResponses := []models.OrdersResponse{}
	for _, item := range orders {
		ordersResponse := models.OrdersResponse{
			ID:        int(item.ID),
			Count:     item.Count,
			NetPrices: item.NetPrices,
			Product:   item.Product,
			User:      item.User,
		}
		ordersResponses = append(ordersResponses, ordersResponse)
	}

	return ordersResponses, nil
}

func (r ordersService) DeleteOrdersById(id int) (err error) {

	err = r.ordersRepo.DeleteOrdersById(id)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	return nil
}
