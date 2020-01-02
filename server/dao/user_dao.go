package dao

import (
	"crudkit/framework/db"
	"crudkit/server/model"
	"time"
)

type UserDao struct {
}

func (*UserDao) GetUserList(page, pageSize int64) ([]*model.User, error) {
	var ret []*model.User
	err := db.Master().
		Where("delete_time = 0").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&ret).
		Error
	return ret, err
}

func (*UserDao) CreateUser(e *model.User) (*model.User, error) {
	e.CreateTime = time.Now().Unix()
	err := db.Master().Create(e).Error
	return e, err
}

func (*UserDao) UpdateUser(e *model.User) error {
	err := db.Master().
		Model(&model.User{}).
		Where("id = ?", e.ID).
		Update(map[string]interface{}{
			"name":        e.Name,
			"age":         e.Age,
			"update_time": time.Now().Unix(),
		}).Error
	return err
}

func (*UserDao) DeleteUserByID(id int64) error {
	return db.Master().
		Model(&model.User{}).
		Where("id = ?", id).
		Update(map[string]interface{}{
			"delete_time": time.Now().Unix(),
		}).
		Error
}
