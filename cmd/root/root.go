package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: " A system information coolection tool",
} // 修正: rootCmdを定義
func init() {
	rootCmd.AddCommand(newVersionCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
