package model

type GlobalRegistrar struct {
	Pk
	Name    string `gorm:"type:varchar" json:"name"`
	Url     string `gorm:"type:varchar" json:"url"`
	Country string `gorm:"type:varchar" json:"country"`
	Number  string `gorm:"type:varchar" json:"number"`
	Email   string `gorm:"type:varchar" json:"email"`
	CreatedAt
	UpdatedAt
}

