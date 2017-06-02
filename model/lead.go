package model

// Lead is a potential sales contact
type Lead struct {
	CompanyName string `csv:"Company Name" gorm:"type:varchar(63)"`
	FirstName   string `csv:"First Name" gorm:"type:varchar(127)"`
	LastName    string `csv:"Last Name" gorm:"type:varchar(127)"`
	JobTitle    string `csv:"Job Title" gorm:"type:varchar(63)"`
	City        string `csv:"City" gorm:"type:varchar(127)"`
	LinkedIn    string `csv:"Linkedin" gorm:"primary_key" gorm:"type:varchar(511)"`
	Industry    string `csv:"Industry" gorm:"type:varchar(63)"`
	Email       string `csv:"Email" gorm:"type:varchar(127)"`
	Sheets      string `csv:"Nome Da Planilha" gorm:"type:varchar(63)"`
	HardBounce  bool   `csv:"Hardbounce"`
}
