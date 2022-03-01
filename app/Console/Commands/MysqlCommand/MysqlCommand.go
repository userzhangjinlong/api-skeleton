package MysqlCommand

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	dataBaseName string
	tableName    string
	dbModel      DBModel
)

var SqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var TableColumnToStructCmd = &cobra.Command{
	Use:   "struct",
	Short: "表字段转为结构体类",
	Long:  "表字段转为结构体类",
	Run: func(cmd *cobra.Command, args []string) {
		columns := dbModel.GetColumns(dataBaseName, tableName)

		template := NewDatabaseTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err := template.Generate(dataBaseName, tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}

	},
}

func init() {
	SqlCmd.AddCommand(TableColumnToStructCmd)
	TableColumnToStructCmd.Flags().StringVarP(&dataBaseName, "schema", "", "", "请输如tableSchema名称")
	TableColumnToStructCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
