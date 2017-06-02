package model

import "strings"

// Irregular people with no linkedin links
type Irregular struct {
	CompanyName string `csv:"Company Name" gorm:"type:varchar(63)"`
	FirstName   string `csv:"First Name" gorm:"type:varchar(127)"`
	LastName    string `csv:"Last Name" gorm:"type:varchar(127)"`
	JobTitle    string `csv:"Job Title" gorm:"type:varchar(63)"`
	City        string `csv:"City" gorm:"type:varchar(127)"`
	LinkedIn    string `csv:"Linkedin" gorm:"-"`
	Industry    string `csv:"Industry" gorm:"type:varchar(63)"`
	Email       string `csv:"Email" gorm:"type:varchar(127);primaty_key:true"`
	Sheets      string `csv:"Nome Da Planilha" gorm:"type:varchar(63)"`
	HardBounce  string `csv:"hardbounce" gorm:"type:varchar(8)"`
}

// CleanDatas removes spaces and does de conversion of datas
func (irr *Irregular) CleanDatas() {
	irr.CompanyName = cleanData(irr.CompanyName)
	irr.FirstName = cleanData(irr.FirstName)
	irr.LastName = cleanData(irr.LastName)
	irr.JobTitle = cleanData(irr.JobTitle)
	irr.City = strings.TrimSpace(irr.City)
	irr.LinkedIn = strings.TrimSpace(irr.LinkedIn)
	irr.Industry = strings.TrimSpace(irr.Industry)
	irr.Email = strings.TrimSpace(irr.Email)
	irr.Sheets = strings.TrimSpace(irr.Sheets)
	irr.HardBounce = strings.TrimSpace(irr.HardBounce)

	if irr.Sheets == "" {
		irr.Sheets = "-"
	}

	if irr.HardBounce == "" {
		irr.HardBounce = "não"
	}
	irr.Sheets = strings.Title(irr.Sheets)
}

// ToLead change Lead to Irregular
func (irr *Irregular) ToLead() *Lead {
	return &Lead{
		CompanyName: irr.CompanyName,
		Industry:    irr.Industry,
		FirstName:   irr.FirstName,
		LastName:    irr.LastName,
		JobTitle:    irr.JobTitle,
		City:        irr.City,
		LinkedIn:    irr.LinkedIn,
		Email:       irr.Email,
		Sheets:      irr.Sheets,
		HardBounce:  irr.HardBounceToBool(),
	}
}

// HardBounceToBool returns true / false of from Irregular HardBouns string
func (irr *Irregular) HardBounceToBool() bool {
	val := strings.TrimSpace(irr.HardBounce)
	val = string(strings.ToLower(val)[0])
	if val == "s" || val == "y" || val == "1" {
		return true
	}
	return false
}
