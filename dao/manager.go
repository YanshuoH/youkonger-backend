package dao

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type manager struct {
	*gorm.DB // alias
}

func (m *manager) Event() *event {
	return &event{m.DB}
}

func GetManager(tx ...*gorm.DB) *manager {
	if len(tx) > 0 {
		return &manager{tx[0]}
	}

	return &manager{Conn}
}


