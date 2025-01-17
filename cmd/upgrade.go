package cmd

import (
	"github.com/spf13/cobra"
	"github.com/fwoawr/cf/pkg/util"
	"github.com/fwoawr/cf/pkg/util/cmdutil"
)

func init() {
	RootCmd.AddCommand(upgradeCmd)
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "更新 cf 到最新版本 (Update cf to the latest version)",
	Long:  "更新 cf 到最新版本 (Update cf to the latest version)",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Upgrade(util.GetCurrentVersion())
	},
}
