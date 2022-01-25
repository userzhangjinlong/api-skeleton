package Console

import (
	MysqlCommand "api-skeleton/app/Console/Commands/MysqlCommand"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	//注册需要执行的命令
	rootCmd.AddCommand(MysqlCommand.SqlCmd)
}

//Execute 执行相应命令
func Execute() error {
	return rootCmd.Execute()
}
