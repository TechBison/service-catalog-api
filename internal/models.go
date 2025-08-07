package internal

type Version struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Number    string `json:"number"`
	ServiceID int    `json:"-"`
}

type Service struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Versions    []Version `json:"versions" gorm:"foreignKey:ServiceID"`
}
