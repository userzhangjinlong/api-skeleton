package MysqlCommand

import (
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
)

type DBModel struct {
}

type TableColumn struct {
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

func (m *DBModel) GetColumns(dbName string, tableName string) []*TableColumn {

	var results *TableColumn
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	db, err := ConnectPoolFactory.GetMysql()
	if err != nil {
		panic("db链接获取异常")
	}

	db.Raw(query, dbName, tableName).Scan(&tableName)

	tableColum := make([]*TableColumn, 0)
	tableColum = append(tableColum, results)
	return tableColum
}
