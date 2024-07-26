package domain

type Doctor struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100;unique"`
	Password string `gorm:"size:100"`
}
