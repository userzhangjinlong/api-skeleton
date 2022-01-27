package MysqlCommand

import (
	"api-skeleton/app/ConstDir"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"fmt"
)

type DBModel struct {
}

type TableColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
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

func (m *DBModel) GetColumns(dbName, tableSchema string, tableName string) []*TableColumn {
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

	fmt.Println("mysqlgo的链接db：", db)
	err = db.Raw(query, tableSchema, tableName).Scan(&results).Error

	fmt.Println("results结果", results)
	fmt.Println("异常信息：", err)

	return results
}
