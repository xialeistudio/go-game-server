// 用户业务类
package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"model"
	"strings"
	"util/hash"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// 注册
func (p User) Register(username, password string) (user *model.User, err error) {
	user = new(model.User)
	// 检测用户是否存在
	if err = p.db.Where("username=?", username).Take(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	} else if err == nil {
		err = errors.New("用户名已存在")
		return
	}
	// 开始注册
	user.Username = username
	user.Password = hash.Md5(strings.NewReader(password))
	err = p.db.Create(user).Error
	return
}
