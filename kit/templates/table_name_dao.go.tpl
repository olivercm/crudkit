package dao

import (
	"crudkit/framework/db"
	"crudkit/crud/server/model"
	"time"
)

type {{.Model}}Dao struct {
}

func (*{{.Model}}Dao) Get{{.Model}}List(currentPage, pageSize int64) ([]*model.{{.Model}}, error) {
	var ret []*model.{{.Model}}
	err := db.Master().
		Where("del_flag = 0").
		Offset((currentPage - 1) * pageSize).Limit(pageSize).
		Find(&ret).
		Error
	return ret, err
}

func (*{{.Model}}Dao) Create{{.Model}}(e *model.{{.Model}}) (*model.{{.Model}}, error) {
	e.CreateTime = time.Now().Unix()
	err := db.Master().Create(e).Error
	return e, err
}

func (*{{.Model}}Dao) Update{{.Model}}(e *model.{{.Model}}) error {
	err := db.Master().
		Model(&model.{{.Model}}{}).
		Where("id = ?", e.Id).
		Update(map[string]interface{}{
			{{range .ModelDaoFields}}{{.}}
            {{end}}
		}).Error
	return err
}

func (*{{.Model}}Dao) Delete{{.Model}}ById(id int64) error {
	return db.Master().
		Model(&model.{{.Model}}{}).
		Where("id = ?", id).
		Update(map[string]interface{}{
			"del_flag": 1,
		}).
		Error
}

func (*{{.Model}}Dao) Count{{.Model}}() (int, error) {
	count := 0
	sql := `select count(*) from {{.Model}} where del_flag = 0 `
	err := db.Master().Raw(sql).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}