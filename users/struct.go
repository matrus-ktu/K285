package users

type Users struct {
	Id        int          `gorm:"primary_key;auto_increment;not_null"`
	SessionID string       `gorm:"type:varchar(255);not_null"`
	FirstName string       `gorm:"type:varchar(255);not_null"`
	LastName  string       `gorm:"type:varchar(255);not_null"`
	Email     string       `gorm:"type:varchar(255);not_null"`
	Password  string       `gorm:"type:varchar(255);not_null"`
	Org       Organization `gorm:"foreignKey:Id"`
}

type Organization struct {
	Id      int    `gorm:"primary_key;auto_increment;not_null"`
	Name    string `gorm:"type:varchar(255)"`
	Code    string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Phone   string `gorm:"type:varchar(255)"`
}
