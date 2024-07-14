package model

type AuthorInfo struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty" gorm:"primary_key"`
}

// TableName returns the table name for the AuthorInfo struct
func (AuthorInfo) TableName() string {
	return "authors"
}
