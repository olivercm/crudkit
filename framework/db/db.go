package db

import (
	"github.com/jinzhu/gorm"
	"sync/atomic"
)

var master struct {
	db    atomic.Value
	debug bool
}

/*
并不能获取数据库信息
这个工具的目的是自动生成对应结构的crud代码
 */
func Master() *gorm.DB {
	return master.db.Load().(*gorm.DB).New()
}


