package model

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type Inventory struct {
	UCFirstServer string
	Model         string
	Server        string
}

//先清理原有的项目结构
func (i *Inventory) CleanStructure() error {
	var dirs = []string{
		"../crudkit/client",
		"../crudkit/server",
	}

	for _, dir := range dirs {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	fmt.Println("clean structure")
	return nil
}

//创建项目结构
func (i *Inventory) GenerateStructure() error {
	var dirs = []string{
		"../crudkit/client/schemas",

		"../crudkit/server/api",
		"../crudkit/server/dao",
		"../crudkit/server/model",
		"../crudkit/server/service",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	fmt.Println("generate structure")
	return nil
}

//根据模版生成dao.go
func (i *Inventory) GenerateDao() error {
	daoTemplatePath := "kit/templates/dao.go.tpl"
	daoPath := fmt.Sprintf("server/dao/dao.go")

	err := i.GenerateFile(daoTemplatePath, daoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", daoPath)
	return nil
}

//根据模版生成{table_name}_dao.go
func (i *Inventory) GenerateTableNameDao(tableName string) error {
	tableNameDaoTemplatePath := "kit/templates/table_name_dao.go.tpl"
	tableNameDaoPath := fmt.Sprintf("server/dao/%s_dao.go", tableName)

	err := i.GenerateFile(tableNameDaoTemplatePath, tableNameDaoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameDaoPath)
	return nil
}

//根据模版生成{table_name}_service.go
func (i *Inventory) GenerateTableNameService(tableName string) error {
	tableNameServiceTemplatePath := "kit/templates/table_name_service.go.tpl"
	tableNameServicePath := fmt.Sprintf("server/service/%s_service.go", tableName)

	err := i.GenerateFile(tableNameServiceTemplatePath, tableNameServicePath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameServicePath)
	return nil
}

//根据模版生成{table_name}_schema_impl.go
func (i *Inventory) GenerateTableNameSchemaImpl(tableName string) error {
	tableNameSchemaImplTemplatePath := "kit/templates/table_name_schema_impl.go.tpl"
	tableNameSchemaImplPath := fmt.Sprintf("client/schemas/%s_schema_impl.go", tableName)

	err := i.GenerateFile(tableNameSchemaImplTemplatePath, tableNameSchemaImplPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameSchemaImplPath)
	return nil
}

//根据模版生成{table_name}.proto
func (i *Inventory) GenerateTableNameProto(tableName string) error {
	tableNameProtoTemplatePath := "kit/templates/table_name_proto.go.tpl"
	tableNameProtoPath := fmt.Sprintf("server/api/%s.proto", tableName)

	err := i.GenerateFile(tableNameProtoTemplatePath, tableNameProtoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameProtoPath)
	return nil
}

//生成对应的pb.go文件
func (i *Inventory) GeneratePbGo(tableName string) error {
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

func (i *Inventory) GenerateFile(templatePath, filePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, i)
	if err != nil {
		return err
	}
	return nil
}
