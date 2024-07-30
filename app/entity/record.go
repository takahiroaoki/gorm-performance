package entity

type Record struct {
	ID   uint   `gorm:"primaryKey"`
	Test string `gorm:"type:varchar(10);not null"`
}
