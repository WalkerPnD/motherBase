package model

// ChildLead is a format to export to CSV
type ChildLead struct {
	CompanyName string `csv:"Company Name"`
	FirstName   string `csv:"First Name"`
	LastName    string `csv:"Last Name"`
	JobTitle    string `csv:"Job Title"`
	City        string `csv:"City"`
	LinkedIn    string `csv:"Linkedin"`
	Industry    string `csv:"Industry"`
	Email       string `csv:"Email"`
	Sheets      string `csv:"Nome da Planilha"`
	HardBounce  string `csv:"Hard Bounce"`
}
