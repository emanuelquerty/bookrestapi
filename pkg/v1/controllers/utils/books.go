package utils

import (
	"fmt"
	"reflect"
)

func CreateSelectQuery(table interface{}) string {
	tableValue := reflect.ValueOf(table).Elem()
	tableType := reflect.TypeOf(table)
	tableName := tableType.Name()
	fieldCount := tableValue.NumField()
	query := "SELECT "
	if tableValue.Kind() == reflect.Struct {
		for i := 0; i < fieldCount; i++ {
			switch tableValue.Field(i).Kind() {
			case reflect.Int:
				query = fmt.Sprintf("%s%d, ", query, tableValue.Field(i).Int())
			case reflect.String:
				query = fmt.Sprintf("%s%s, ", query, tableValue.Field(i).String())
			default:
				fmt.Print("Unsuported Type")
			}
		}
		query = fmt.Sprintf("%s FROM %s, ", query[:len(query)-2], tableName)
	}
	return query
}
