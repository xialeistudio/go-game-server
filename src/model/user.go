// 用户模型
package model

import "time"

type User struct {
	Id        uint      `gorm:"primary_key"`
	Username  string
	Password  string
	Banned    bool
	CreatedAt time.Time `gorm:"created_at"`
}

func (User) TableName() string {
	return "m_user"
}

func (p *User) BeforeCreate() error {
	p.CreatedAt = time.Now()
	return nil
}
