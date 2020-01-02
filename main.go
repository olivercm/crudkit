package main

import (
	"crudkit/kit"
	"crudkit/kit/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"os/exec"
)

var tableName = "user"

func main() {

	//先清理原有的项目结构

	//根据表名生成model

	err := generateStructure()
	if err != nil {
		log.Error(err)
		return
	}

	sweaters := model.Inventory{"Customer", "User", "customer"}
	err = generateDao(sweaters)
	if err != nil {
		log.Error(err)
		return
	}

	err = generateTableNameService(sweaters)
	if err != nil {
		log.Error(err)
		return
	}

	err = generateTableNameSchemaImpl(sweaters)
	if err != nil {
		log.Error(err)
		return
	}

	err = generateTableNameProto(sweaters)
	if err != nil {
		log.Error(err)
		return
	}

	err = generatePbGo()
	if err != nil {
		log.Error(err)
		return
	}
}

//创建项目结构
func generateStructure() error {
	var dirs = []string{
		"../crudkit/client/schemas",

		"../crudkit/server/api",
		"../crudkit/server/dao",
		"../crudkit/server/model",
		"../crudkit/server/service",
	}

	for _, dir := range dirs {
		err := kit.MkdirAll(dir)
		if err != nil {
			return err
		}
	}
	fmt.Println("generate structure")
	return nil
}

//根据模版生成dao.go
func generateDao(sweaters model.Inventory) error {
	daoTemplatePath := "kit/templates/dao.go.tpl"
	daoPath := fmt.Sprintf("server/dao/dao.go")

	err := kit.GenerateFile(sweaters, daoTemplatePath, daoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", daoPath)
	return nil
}

//根据模版生成{table_name}_service.go
func generateTableNameService(sweaters model.Inventory) error {
	tableNameServiceTemplatePath := "kit/templates/model_service.go.tpl"
	tableNameServicePath := fmt.Sprintf("server/service/%s_service.go", tableName)

	err := kit.GenerateFile(sweaters, tableNameServiceTemplatePath, tableNameServicePath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameServicePath)
	return nil
}

//根据模版生成{table_name}_schema_impl.go
func generateTableNameSchemaImpl(sweaters model.Inventory) error {
	tableNameSchemaImplTemplatePath := "kit/templates/model_schema_impl.go.tpl"
	tableNameSchemaImplPath := fmt.Sprintf("client/schemas/%s_schema_impl.go", tableName)

	err := kit.GenerateFile(sweaters, tableNameSchemaImplTemplatePath, tableNameSchemaImplPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameSchemaImplPath)
	return nil
}

//根据模版生成{table_name}.proto
func generateTableNameProto(sweaters model.Inventory) error {
	tableNameProtoTemplatePath := "kit/templates/model_proto.go.tpl"
	tableNameProtoPath := fmt.Sprintf("server/api/%s.proto", tableName)

	err := kit.GenerateFile(sweaters, tableNameProtoTemplatePath, tableNameProtoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameProtoPath)
	return nil
}

//生成对应的pb.go文件
func generatePbGo() error {
	cmd := exec.Command("bash", "-c", "make proto")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("exec: %s", out)

	tableNamePbGoPath := fmt.Sprintf("server/api/%s.pb.go", tableName)
	fmt.Printf("generate %s\n", tableNamePbGoPath)
	return nil
}
