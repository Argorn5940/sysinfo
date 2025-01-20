package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "v1.0.0"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("sysinfo version %s\n", Version)
		},
	}
}
