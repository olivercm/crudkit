package kit

import (
	"crudkit/kit/model"
	"os"
	"text/template"
)

func CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func GenerateFile(sweaters model.Inventory, templatePath, filePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, sweaters)
	if err != nil {
		return err
	}
	return nil
}
