package models

// User defines the user object
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}
