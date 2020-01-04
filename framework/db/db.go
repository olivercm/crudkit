package db

import (
	"github.com/jinzhu/gorm"
	"sync/atomic"
)

var master struct {
	db    atomic.Value
	debug bool
}

//并不能执行，只是为了生成的代码不报错
func Master() *gorm.DB {
	return master.db.Load().(*gorm.DB).New()
}
