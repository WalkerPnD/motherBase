package model

// Irregular people with no linkedin links
type Irregular struct {
	CompanyName string `csv:"Company Name" gorm:"type:varchar(63)"`
	FullName    string `csv:"Full Name" gorm:"type:varchar(127)"`
	JobTitle    string `csv:"Job Title" gorm:"type:varchar(63)"`
	City        string `csv:"City" gorm:"type:varchar(127)"`
	LinkedIn    string `csv:"Linkedin" gorm:"-"`
	Industry    string `csv:"Industry" gorm:"type:varchar(63)"`
	Email       string `csv:"email" gorm:"type:varchar(127);primaty_key:true"`
	Sheets      string `csv:"Nome Da Planilha" gorm:"type:varchar(63)"`
	HardBounce  string `csv:"hardbounce"`
}

// IrregularToLead change Lead to Irregular
func IrregularToLead(irr *Irregular) *Lead {
	return &Lead{
		CompanyName: irr.CompanyName,
		Industry:    irr.Industry,
		FullName:    irr.FullName,
		JobTitle:    irr.JobTitle,
		City:        irr.City,
		LinkedIn:    irr.LinkedIn,
		Email:       irr.Email,
		Sheets:      []Sheet{Sheet{Name: irr.Sheets}},
		HardBounce:  false,
	}
}
