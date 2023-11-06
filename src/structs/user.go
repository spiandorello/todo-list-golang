package structs

type User struct {
	Metadata
	Name     string `gorm:"type:varchar(32);not null"`
	Username string `gorm:"type:varchar(100);not null;unique"`
}
