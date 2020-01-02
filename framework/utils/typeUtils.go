package utils

import (
	"crudkit/framework/model"
	"encoding/json"
	"strings"
)

func removeRPCErrPrefix(s string) string {
	prefix := "rpc error: code = Unknown desc = "
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

func BindType(m map[string]interface{}, model interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, model)
}

func BaseTypeError(err error) (model.BaseType, error) {
	baseType := model.BaseType{
		Success:  false,
		ErrorMsg: removeRPCErrPrefix(err.Error()),
		Extra:    err.Error(),
	}
	return baseType, nil
}

func BaseTypeErrors(err error) ([]model.BaseType, error) {
	baseType := model.BaseType{
		Success:  false,
		ErrorMsg: removeRPCErrPrefix(err.Error()),
		Extra:    err.Error(),
	}
	return []model.BaseType{baseType}, nil
}

func BaseTypeOK(data interface{}) (model.BaseType, error) {
	baseType := model.BaseType{
		Success: true,
		Data:    data,
	}
	return baseType, nil
}

func BaseTypesManyOK(datas interface{}, totalNum, curretPage int64) ([]model.BaseType, error) {
	baseType := model.BaseType{
		Success: true,
		Datas:   datas,
		Total:   int(totalNum),
		Current: int(curretPage),
	}
	return []model.BaseType{baseType}, nil
}
