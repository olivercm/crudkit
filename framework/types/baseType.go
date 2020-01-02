package types

import (
	"fmt"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
)

var DefaultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdminUserType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func InitType(object *graphql.Object) *graphql.Object {
	var baseType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "baseType",
			Fields: graphql.Fields{
				"success": &graphql.Field{
					Type: graphql.Boolean,
				},
				"total": &graphql.Field{
					Type: graphql.Int,
				},
				"body": &graphql.Field{
					Type: object,
				},
				"bodys": &graphql.Field{
					Type: graphql.NewList(object),
				},
				"errorCode": &graphql.Field{
					Type: graphql.Int,
				},
				"errorMsg": &graphql.Field{
					Type: graphql.String,
				},
				"extra": &graphql.Field{
					Type: graphql.String,
				},
				"current": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
	return baseType
}

func InitBaseType(input interface{}) *graphql.Object {
	baseType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "baseType",
			Fields: InitFields(input),
		},
	)
	return baseType
}

func InitFields(input interface{}) graphql.Fields {
	graphqlFields := graphql.Fields{}
	getValue := reflect.ValueOf(input)
	getType := reflect.TypeOf(input)
	for i := 0; i < getType.NumField(); i++ {
		rfield :=getType.Field(i)
		jsonTag := rfield.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}
		value := getValue.Field(i).Interface()
		graphqlField := &graphql.Field{}
		switch value.(type) {
		case int64, int32, int16, int8, int, uint64:
			graphqlField.Type = graphql.Int

		case string:
			graphqlField.Type = graphql.String

		case bool:
			graphqlField.Type = graphql.Boolean

		case float64, float32:
			graphqlField.Type = graphql.Float

		default:
			graphqlField.Type = JSONObjectType
		}
		fieldKey := getJSONTagName(jsonTag, Capitalize(rfield.Name))
		graphqlFields[fieldKey] = graphqlField
	}
	return graphqlFields
}
func getJSONTagName(tag, defaultValue string) string {
	attrs := strings.Split(tag, ",")
	var name string
	if len(attrs) > 0 {
		name = attrs[0]
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return defaultValue
	}
	return name
}

var JSONObjectType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSONObjectType",
	Description: "JSONObjectType 不对对象进行任何 parse，原始结构将直接被转换为json对象",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return valueAST.GetValue()
	},
})

func NewGraphQLListTypeFromRPCType(listItemType interface{}) *graphql.List {
	return graphql.NewList(InitType(InitBaseType(listItemType)))
}

func NewGraphQLTypeFromRPCType(itemType interface{}) *graphql.Object {
	return InitType(InitBaseType(itemType))
}

func NewGraphQLArgsFromRPCType(itemType interface{}) graphql.FieldConfigArgument {
	return InitFieldConfigArgument(itemType)
}

func InitFieldConfigArgument(input interface{}) graphql.FieldConfigArgument {
	graphqlFields := graphql.FieldConfigArgument{}
	getValue := reflect.ValueOf(input)
	getType := reflect.TypeOf(input)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i).Name
		if strings.Contains(field, "XXX") {
			continue
		}
		value := getValue.Field(i).Interface()
		graphqlField := &graphql.ArgumentConfig{}

		switch value.(type) {
		case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
			graphqlField.Type = graphql.Int
			graphqlField.DefaultValue = 0

		case string:
			graphqlField.Type = graphql.String
			graphqlField.DefaultValue = ""

		case bool:
			graphqlField.Type = graphql.Boolean
			graphqlField.DefaultValue = false

		case float64, float32:
			graphqlField.Type = graphql.Float
			graphqlField.DefaultValue = 0.0
		case []int64, []int32, []int16, []int8, []int, []uint64, []uint32, []uint16, []uint8, []uint:
			graphqlField.Type = graphql.NewList(graphql.Int)
			graphqlField.DefaultValue = make([]int64, 0)
		case []string:
			graphqlField.Type = graphql.NewList(graphql.String)
			graphqlField.DefaultValue = make([]string, 0)
		default:
			panic(fmt.Sprintf("InitFieldConfigArgument,unknown type: %s", getType.Field(i).Name))
		}
		jsonTag := getType.Field(i).Tag.Get("json")
		fieldKey := getJSONTagName(jsonTag, Capitalize(field))
		graphqlFields[fieldKey] = graphqlField
		// if fieldKey != Capitalize(field) {
		//	log.Infoln(fieldKey, "!=", Capitalize(field))
		// }
		// graphqlFields[Capitalize(field)] = graphqlField
	}
	return graphqlFields
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[i] += 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
