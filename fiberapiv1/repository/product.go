package repository

import (
	"fiberapiv1/entity"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProductGroup(entity.ProductGroup) (err error)
	InsertProduct(entity.Product) (err error)
	GetAllProductGroup() (productGroups []entity.ProductGroup, err error)
	GetAllProductV1() (products []entity.Product, err error) //bad style
	GetAllProductV2() (products []entity.Product, err error) //good style
	GetAllProductByGroupID(groupId int) (products []entity.Product, err error)
}

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	return productRepositoryDB{db}
}

func (r productRepositoryDB) InsertProductGroup(productGroup entity.ProductGroup) (err error) {
	err = r.db.Save(&productGroup).Error
	return err
}
func (r productRepositoryDB) InsertProduct(product entity.Product) (err error) {
	err = r.db.Save(&product).Error
	return err
}

func (r productRepositoryDB) GetAllProductGroup() (productGroups []entity.ProductGroup, err error) {
	err = r.db.Find(&productGroups).Error
	return productGroups, err
}

// query v1
func (r productRepositoryDB) GetAllProductV1() (products []entity.Product, err error) {
	//err = r.db.Find(&products).Error
	var results []map[string]interface{}

	sql := "SELECT " +
		" products.id as product_id ," +
		" products.name as product_name," +
		" products.price as product_price ," +
		" products.image_uri as product_image, " +
		" product_groups.id as product_groups_id, " +
		" product_groups.name as product_groups_name, " +
		" product_groups.image_uri as product_groups_image " +
		" FROM products" +
		" INNER JOIN product_groups" +
		" ON product_groups.id = products.product_group_id"

	r.db.Raw(sql).Scan(&results)

	for _, pro := range results {
		pro_id, _ := strconv.Atoi(fmt.Sprintf("%s", pro["product_id"]))
		pro_name := fmt.Sprintf("%s", pro["product_name"])
		pro_image := fmt.Sprintf("%s", pro["product_image"])

		proGroup_id, _ := strconv.Atoi(fmt.Sprintf("%s", pro["product_groups_id"]))
		proGroup_name := fmt.Sprintf("%s", pro["product_groups_name"])
		proGroup_image := fmt.Sprintf("%s", pro["product_groups_image"])

		proGroup := entity.ProductGroup{
			ID:       uint64(proGroup_id),
			Name:     proGroup_name,
			ImageUri: proGroup_image,
		}
		product := entity.Product{
			ID:           uint64(pro_id),
			Name:         pro_name,
			ImageUri:     pro_image,
			ProductGroup: proGroup,
		}
		products = append(products, product)
	}
	return products, err
}

// query v2
func (r productRepositoryDB) GetAllProductV2() (products []entity.Product, err error) {
	err = r.db.Preload("ProductGroup").Find(&products).Error
	return products, err
}

func (r productRepositoryDB) GetAllProductByGroupID(groupId int) (products []entity.Product, err error) {
	err = r.db.Preload("ProductGroup").Where("product_group_id=?", groupId).Find(&products).Error
	return products, err
}
