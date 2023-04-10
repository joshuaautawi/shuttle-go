package user

type User struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(100)" json:"first_name"`
	LastName  string `gorm:"type:varchar(100)" json:"last_name"`
}
