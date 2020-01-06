package model

import (
	"crudkit/framework/utils"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Inventory struct {
	UCFirstServer     string
	Model             string
	Server            string
	Fields            []string
	ProtoFields       []string
	ModelDaoFields    []string
	ServiceFields     []string
	ServiceListFields []string
}

//先清理原有的项目结构
func (i *Inventory) CleanStructure() error {
	var dirs = []string{
		"../crudkit/crud/client",
		"../crudkit/crud/server",
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
		"../crudkit/crud/client/schemas",

		"../crudkit/crud/server/api",
		"../crudkit/crud/server/dao",
		"../crudkit/crud/server/model",
		"../crudkit/crud/server/service",
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
	daoPath := fmt.Sprintf("crud/server/dao/dao.go")

	err := i.generateFile(daoTemplatePath, daoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", daoPath)
	return nil
}

//根据模版生成{table_name}_dao.go
func (i *Inventory) GenerateTableNameDao(tableName string) error {
	tableNameDaoTemplatePath := "kit/templates/table_name_dao.go.tpl"
	tableNameDaoPath := fmt.Sprintf("crud/server/dao/%s_dao.go", tableName)

	tableNamePath := fmt.Sprintf("crud/server/model/%s.go", tableName)
	i.ModelDaoFields = getFields(tableNamePath, func(s []*ast.Field) []string {
		fields := make([]string, 0)
		var field string
		for _, tmpField := range s {
			name := tmpField.Names[0].Name
			if utils.LowerCase(name) != "create_time" &&
				utils.LowerCase(name) != "update_time" &&
				utils.LowerCase(name) != "delete_time" {
				field = fmt.Sprintf("\"%s\": e.%s,", utils.LowerCase(name), name)
			} else {
				field = fmt.Sprintf("\"%s\": time.Now().Unix(),", utils.LowerCase(name))
			}
			fields = append(fields, field)
		}
		return fields
	})

	err := i.generateFile(tableNameDaoTemplatePath, tableNameDaoPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameDaoPath)
	return nil
}

//根据模版生成{table_name}_service.go
func (i *Inventory) GenerateTableNameService(tableName string) error {
	tableNameServiceTemplatePath := "kit/templates/table_name_service.go.tpl"
	tableNameServicePath := fmt.Sprintf("crud/server/service/%s_service.go", tableName)

	tableNamePath := fmt.Sprintf("crud/server/model/%s.go", tableName)
	i.ServiceFields = getFields(tableNamePath, func(s []*ast.Field) []string {
		fields := make([]string, 0)
		for _, tmpField := range s {
			name := tmpField.Names[0].Name
			field := fmt.Sprintf("%s: req.%s,", name, name)
			fields = append(fields, field)
		}
		return fields
	})

	i.ServiceListFields = getFields(tableNamePath, func(s []*ast.Field) []string {
		fields := make([]string, 0)
		for _, tmpField := range s {
			name := tmpField.Names[0].Name
			field := fmt.Sprintf("%s: data.%s,", name, name)
			fields = append(fields, field)
		}
		return fields
	})

	err := i.generateFile(tableNameServiceTemplatePath, tableNameServicePath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameServicePath)
	return nil
}

//根据模版生成{table_name}_schema_impl.go
func (i *Inventory) GenerateTableNameSchemaImpl(tableName string) error {
	tableNameSchemaImplTemplatePath := "kit/templates/table_name_schema_impl.go.tpl"
	tableNameSchemaImplPath := fmt.Sprintf("crud/client/schemas/%s_schema_impl.go", tableName)

	err := i.generateFile(tableNameSchemaImplTemplatePath, tableNameSchemaImplPath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNameSchemaImplPath)
	return nil
}

//根据模版生成{table_name}.proto
func (i *Inventory) GenerateTableNameProto(tableName string) error {
	tableNameProtoTemplatePath := "kit/templates/table_name.proto.tpl"
	tableNameProtoPath := fmt.Sprintf("crud/server/api/%s.proto", tableName)

	tableNamePath := fmt.Sprintf("crud/server/model/%s.go", tableName)
	i.ProtoFields = getFields(tableNamePath, func(s []*ast.Field) []string {
		fields := make([]string, 0)
		for i, tmpField := range s {
			name := tmpField.Names[0].Name
			field := fmt.Sprintf("%s %s = %d;", tmpField.Type, utils.LowerCaseFirst(name), i+1)
			fields = append(fields, field)
		}
		return fields
	})

	err := i.generateFile(tableNameProtoTemplatePath, tableNameProtoPath)
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

	tableNamePbGoPath := fmt.Sprintf("crud/server/api/%s.pb.go", tableName)
	fmt.Printf("generate %s\n", tableNamePbGoPath)
	return nil
}

func (i *Inventory) generateFile(templatePath, filePath string) error {
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

//根据表名生成model
func (i *Inventory) GenerateStruct(tableName string, columns []DbMeta, jsonAnnotation bool) error {
	fields := generateFields(columns, jsonAnnotation)
	i.Fields = fields

	tableNameTemplatePath := "kit/templates/table_name.go.tpl"
	tableNamePath := fmt.Sprintf("crud/server/model/%s.go", tableName)

	err := i.generateFile(tableNameTemplatePath, tableNamePath)
	if err != nil {
		return err
	}
	fmt.Printf("generate %s\n", tableNamePath)
	return nil
}

//生成struct fields
func generateFields(columns []DbMeta, jsonAnnotation bool) []string {
	var fields []string
	var field = ""
	for i, c := range columns {
		columnName := c.ColumnName
		dataType := utils.GetGolangType(c.DataType)
		fieldName := utils.UpperCamel(columnName)

		var annotations []string
		if i == 0 {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s;primary_key\"", columnName))
		} else {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", columnName))
		}

		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", columnName))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`", fieldName, dataType,
				strings.Join(annotations, " "))
		} else {
			field = fmt.Sprintf("%s %s", fieldName, dataType)
		}

		fields = append(fields, field)
	}
	return fields
}

type convert func([]*ast.Field) []string

//根据传入fn生成不同的fields
func getFields(filePath string, fn convert) []string {
	fields := make([]string, 0)
	src := utils.ReadFile(filePath)

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "fields", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(file, func(x ast.Node) bool {
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}

		fields = fn(s.Fields.List)
		return false
	})
	return fields
}




