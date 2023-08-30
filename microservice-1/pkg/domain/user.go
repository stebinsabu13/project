package domain

type User struct {
	ID        uint   `json:"id" gorm:"primarykey;auto_increment"`
	FirstName string `json:"firstname" gorm:"not null"`
	LastName  string `json:"lastname" gorm:"not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	MobileNum string `json:"mobilenum" gorm:"uniqueIndex;not null"`
}
