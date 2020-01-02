package main

import (
	"fmt"
)

var dirs = []string{
	"../crudkit/client/schemas",

	"../crudkit/server/api",
	"../crudkit/server/dao",
	"../crudkit/server/model",
	"../crudkit/server/service",
}

var protoPath = fmt.Sprintf("../crudkit/server/api/%s.proto", tableName)

var schemaPath = fmt.Sprintf("../crudkit/client/schemas/%s_schema_impl.go", tableName)

var tableName = "user"

func main() {

	//先清理原有的项目结构



	//创建项目结构
	//for _, dir := range dirs {
	//	err := kit.MkdirAll(dir)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//}

	//根据表名生成model


	//根据表名生成proto
	//err := kit.CreateFile(protoPath)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//根据表名生成schema
	//err := kit.CreateFile(schemaPath)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//根据model生成proto文件的curd


	//根据model生成对应dao方法


}
