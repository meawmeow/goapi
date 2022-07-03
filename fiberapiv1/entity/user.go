package entity

type User struct {
	ID       uint64 `gorm:"primaryKey:autoIncrement"`
	Username string `gorm:"type:varchar(255)"`
	Email    string `gorm:"unique"`
	Password []byte `gorm:"->;<-;not null"`
	Address  string `gorm:"type:varchar(255)"`
}
