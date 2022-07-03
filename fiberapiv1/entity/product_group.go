package entity

type ProductGroup struct {
	ID       uint64 `gorm:"primaryKey:autoIncrement"`
	Name     string `gorm:"type:varchar(255)"`
	ImageUri string
}
