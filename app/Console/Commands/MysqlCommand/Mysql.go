package MysqlCommand

import (
	"api-skeleton/app/ConstDir"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBModel struct {
}

type TableColumn struct {
	gorm.Model
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
	"char":      "string",
}

func (m *DBModel) GetColumns(dbName, tableName string) []*TableColumn {

	var results []*TableColumn
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	if dbName == "" {
		dbName = ConstDir.SCHEMA
	}

	db, err := ConnectPoolFactory.GetMysql(dbName)
	if err != nil {
		panic("db链接获取异常")
	}

	err = db.Raw(query, dbName, tableName).Scan(&results).Error

	fmt.Println(results)

	return results
}
