package model

type User struct {
    Id int64 `gorm:"column:id;primary_key" json:"id"`
    Name string `gorm:"column:name" json:"name"`
    Age int64 `gorm:"column:age" json:"age"`
    CreateTime int64 `gorm:"column:create_time" json:"create_time"`
    UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
    DeleteTime int64 `gorm:"column:delete_time" json:"delete_time"`
    
}

