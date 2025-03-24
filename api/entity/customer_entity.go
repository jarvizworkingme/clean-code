package entity

func (u *Customer) TableName() string {
	return "customer"
}

type Customer struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"type:Text;not null;unique"`
	BaseEntity
}
