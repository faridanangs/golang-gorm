package model

type GuestBook struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Message   string `gorm:"column:message"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt string `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}
