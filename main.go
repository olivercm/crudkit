package main

import (
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


	//根据表名生成model


	err = sweaters.GenerateStructure()
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
