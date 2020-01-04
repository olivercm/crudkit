package main

import (
	"crudkit/framework/mysql"
	"crudkit/framework/utils"
	"crudkit/kit/model"
	"flag"
	"github.com/labstack/gommon/log"
)

type data struct {
	ServerName string
	User string
	Password string
	Address string
	TableName string
	Database string
}

func main() {

	var d data
	flag.StringVar(&d.ServerName, "serverName", "", "The serverName used for name of microservice.")
	flag.StringVar(&d.User, "user", "", "The user used for mysql connection.")
	flag.StringVar(&d.Password, "password", "", "The password used for mysql connection.")
	flag.StringVar(&d.Address, "address", "", "The address used for mysql connection.")
	flag.StringVar(&d.TableName, "tableName", "", "The tableName used for getting fields.")
	flag.StringVar(&d.Database, "database", "", "The database used for inquire tableName.")
	flag.Parse()

	serverName := d.ServerName
	user := d.User
	password := d.Password
	address := d.Address
	tableName := d.TableName
	database := d.Database

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

	db, err := mysql.Master(user, password, address)
	if err != nil {
		log.Error(err)
		return
	}

	var metas []model.DbMeta
	err = db.Raw(`select column_name, data_type, column_key, extra
		from information_schema.columns
		where table_schema = ?
		and table_name = ?`, database, tableName).Find(&metas).Error
	if len(metas) == 0 {
		log.Error("metas are empty; incorrect table_schema or table_name")
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
