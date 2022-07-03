package entity

type Product struct {
	ID             uint64 `gorm:"primaryKey:autoIncrement"`
	Name           string `gorm:"type:varchar(255)"`
	Price          string
	ImageUri       string
	ProductGroupID uint64
	ProductGroup   ProductGroup
}
