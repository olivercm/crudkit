package main

import (
	"crudkit/framework/db"
	"crudkit/framework/utils"
	"crudkit/kit/model"
	"github.com/labstack/gommon/log"
)

var tableName = "user"
var serverName = "customer"

func main() {

	sweaters := &model.Inventory{
		UCFirstServer: utils.UpperCaseFirst(serverName),
		Model:         utils.UpperCamel(tableName),
		Server:        serverName,
	}

	err := sweaters.CleanStructure()
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateStructure()
	if err != nil {
		log.Error(err)
		return
	}

	//还需要调整
	var metas []model.DbMeta
	err = db.Master().Raw(`select column_name, data_type, column_key, extra 
		from information_schema.columns
		where table_schema = 'db'
		and table_name = ?`, tableName).Find(&metas).Error
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateStruct(tableName, metas, true)
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateDao()
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateTableNameDao(tableName)
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateTableNameService(tableName)
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateTableNameSchemaImpl(tableName)
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GenerateTableNameProto(tableName)
	if err != nil {
		log.Error(err)
		return
	}

	err = sweaters.GeneratePbGo(tableName)
	if err != nil {
		log.Error(err)
		return
	}
}
