package tencent

import (
	"github.com/spf13/cobra"
	"github.com/fwoawr/cf/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(tencentCmd)
}

var tencentCmd = &cobra.Command{
	Use:   "tencent",
	Short: "执行与腾讯云相关的操作 (Perform Tencent Cloud related operations)",
	Long:  "执行与腾讯云相关的操作 (Perform Tencent Cloud related operations)",
}
