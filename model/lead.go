package model

// Lead is a potential sales contact
type Lead struct {
	CompanyName string  `csv:"Company Name" gorm:"type:varchar(63)"`
	FullName    string  `csv:"Full Name" gorm:"type:varchar(127)"`
	JobTitle    string  `csv:"Job Title" gorm:"type:varchar(63)"`
	City        string  `csv:"City" gorm:"type:varchar(127)"`
	LinkedIn    string  `csv:"Linkedin" gorm:"primary_key" gorm:"type:varchar(511)"`
	Industry    string  `csv:"Industry" gorm:"type:varchar(63)"`
	Email       string  `csv:"Email" gorm:"type:varchar(127)"`
	Sheets      []Sheet `gorm:"many2many:lead_sheets;"`
	HardBounce  bool    `csv:"hardbounce"`
}
