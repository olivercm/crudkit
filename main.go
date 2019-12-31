package main

import (
	"fmt"
)

var dirs = []string{
	"../curdkit/client/dao",
	"../curdkit/client/model",
	"../curdkit/client/schemas",

	"../curdkit/server/api",
	"../curdkit/server/dao",
	"../curdkit/server/model",
	"../curdkit/server/service",
}

var protoPath = fmt.Sprintf("../curdkit/server/api/%s.proto", tableName)

var schemaPath = fmt.Sprintf("../curdkit/client/schemas/%s_schema_impl.go", tableName)

var tableName = "user"

func main() {
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
