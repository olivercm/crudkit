package model

type User struct {
	Id         int64  `gorm:"id,AUTO_INCREMENT"`
	Name       string `gorm:"column:name"`
	Age        int64  `gorm:"column:age"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
	DeleteTime int64  `gorm:"column:delete_time"`
}
