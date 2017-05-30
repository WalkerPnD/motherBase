package model

// Lead is a potential sales contact
type Lead struct {
	CompanyName string `csv:"Company Name" gorm:"type:varchar(63)"`
	FullName    string `csv:"Full Name" gorm:"type:varchar(127)"`
	JobTitle    string `csv:"Job Title" gorm:"type:varchar(63)"`
	City        string `csv:"City" gorm:"type:varchar(127)"`
	LinkedIn    string `csv:"Linkedin" gorm:"primary_key" gorm:"type:varchar(511)"`
	Industry    string `csv:"Industry" gorm:"type:varchar(63)"`
	Email       string `csv:"Email" gorm:"type:varchar(127)"`
	Sheets      string `csv:"Nome Da Planilha" gorm:"type:varchar(63)"`
	HardBounce  bool   `csv:"Hardbounce"`
}

// ToChildLead change Lead to Irregular
func (l *Lead) ToChildLead() *ChildLead {
	hb := "não"
	if l.HardBounce {
		hb = "sim"
	}

	return &ChildLead{
		CompanyName: l.CompanyName,
		Industry:    l.Industry,
		FullName:    l.FullName,
		JobTitle:    l.JobTitle,
		City:        l.City,
		LinkedIn:    l.LinkedIn,
		Email:       l.Email,
		Sheets:      l.Sheets,
		HardBounce:  hb,
	}
}
