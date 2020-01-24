package utils

func GetGolangType(dataType string) string {
	switch dataType {
	case "tinyint", "smallint", "mediumint", "int8", "int4", "int2":
		return "int32"

	case "bigint", "int":
		return "int64"

	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext", "varchar2", "json", "jsonb":
		return "string"

	case "decimal", "double":
		return "float64"

	case "float":
		return "float32"

	case "bool":
		return "bool"

	case "date":
		return "string"

	case "timestamp":
		return "int64"
	}

	return "unknown"
}
