package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (m *defaultUserModel) FindUserBy(db *gorm.DB, field string, value interface{}) ([]User, error) {
	var users []User
	if res := db.Where(field+" = ?", value).Find(&users); res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(users)
	if len(users) == 0 {
		return nil, errors.New("查询为空")
	}
	return users, nil
}
