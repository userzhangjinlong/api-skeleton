package MysqlCommand

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Global"
	"fmt"
	"gorm.io/plugin/dbresolver"
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
	"int":       "int64",
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

func (m *DBModel) GetColumns(tableSchema string, tableName string) []*TableColumn {
	var results []*TableColumn
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"

	//todo::已知目前集群链接使用方式；文档DBResolver 会根据工作表、struct 自动切换连接暂未生效发现待处理
	err := Global.DB.Clauses(dbresolver.Use(ConstDir.SCHEMA)).
		Raw(query, tableSchema, tableName).Scan(&results).Error

	fmt.Println("results结果", results)
	fmt.Println("异常信息：", err)
	if err != nil {
		return nil
	}

	return results
}
