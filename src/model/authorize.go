// 用户凭证模型
package model

import "time"

type Authorize struct {
	Id        uint      `gorm:"primary_key"`
	Token     string
	UserId    uint      `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
	ExpiredAt time.Time `gorm:"expired_at"`
}

func (Authorize) TableName() string {
	return "m_authorize"
}

func (p *Authorize) BeforeCreate() error {
	p.CreatedAt = time.Now()
	return nil
}
