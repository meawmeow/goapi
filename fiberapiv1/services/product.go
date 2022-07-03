package services

import (
	"fiberapiv1/entity"
	"fiberapiv1/errs"
	"fiberapiv1/logs"
	"fiberapiv1/models"
	"fiberapiv1/repository"
)

type ProductService interface {
	CreateProduct(models.ProductRequest) (err error)
	CreateProductGroup(models.ProductGroupRequest) (err error)
	Products() ([]models.ProductResponse, error)
	ProductsByGroupId(groupId int) ([]models.ProductResponse, error)
	ProductGroups() ([]models.ProductGroupResponse, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}

func (r productService) CreateProduct(request models.ProductRequest) (err error) {
	product := entity.Product{
		Name:           request.Name,
		Price:          request.Price,
		ImageUri:       request.ImageUri,
		ProductGroupID: uint64(request.ProductGroupID),
	}
	err = r.productRepo.InsertProduct(product)
	return err
}

func (r productService) CreateProductGroup(request models.ProductGroupRequest) (err error) {
	productGroup := entity.ProductGroup{
		Name:     request.Name,
		ImageUri: request.ImageUri,
	}
	err = r.productRepo.InsertProductGroup(productGroup)
	return err
}

func (r productService) ProductGroups() (response []models.ProductGroupResponse, err error) {

	pGroups, err := r.productRepo.GetAllProductGroup()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	pGroupResponses := []models.ProductGroupResponse{}
	for _, pGroup := range pGroups {
		pGroupResponse := models.ProductGroupResponse{
			ID:       pGroup.ID,
			Name:     pGroup.Name,
			ImageUri: pGroup.ImageUri,
		}
		pGroupResponses = append(pGroupResponses, pGroupResponse)
	}

	return pGroupResponses, nil
}

func (r productService) Products() (response []models.ProductResponse, err error) {

	products, err := r.productRepo.GetAllProductV2()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	productResponses := []models.ProductResponse{}
	for _, product := range products {
		productResponse := models.ProductResponse{
			ID:           int(product.ID),
			Name:         product.Name,
			Price:        product.Price,
			ImageUri:     product.ImageUri,
			ProductGroup: product.ProductGroup,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (r productService) ProductsByGroupId(groupId int) (response []models.ProductResponse, err error) {

	products, err := r.productRepo.GetAllProductByGroupID(groupId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	productResponses := []models.ProductResponse{}
	for _, product := range products {
		productResponse := models.ProductResponse{
			ID:           int(product.ID),
			Name:         product.Name,
			Price:        product.Price,
			ImageUri:     product.ImageUri,
			ProductGroup: product.ProductGroup,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
