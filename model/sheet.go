package model

// Lead is a potential sales contact
type Sheet struct {
	ID   int    `gorm:"primary_key;index:idx_name"`
	Name string `gorm:"type:varchar(63)"`
}
