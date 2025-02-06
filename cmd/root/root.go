package root

import (
	"fmt"
	"os"

	"github.com/Argorn5940/sysinfo/cmd/collector"
	"github.com/spf13/cobra"
)

// RootCmdをエクスポート（大文字で始める）
var RootCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "システム情報を表示するCLIツール",
	Long: `sysinfoは、システムの様々な情報（cpu、メモリ、ディスク、ネットワークを
	収集して表示するCLIツールです。`,
	Run: func(cmd *cobra.Command, args []string) {
		//デフォルトの動作（すべての情報を表示）
		collector.ShowAllInfo()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	//サブコマンドの追加
	RootCmd.AddCommand(newVersionCmd())
	RootCmd.AddCommand(cpuCmd)
	RootCmd.AddCommand(memoryCmd)
	RootCmd.AddCommand(diskCmd)
	RootCmd.AddCommand(networkCmd)
	RootCmd.AddCommand(allCmd)
}

// bバージョン情報用のコマンド
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "バージョン情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysinfo v1.0.0")
	},
}

// cPU情報用のコマンド
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "CPU情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		collector.ShowCPUInfo()
	},
}

// メモリ情報用のコマンド
var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "メモリ情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		collector.ShowMemoryInfo()
	},
}

// ディスク情報用のコマンド
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "ディスク情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		collector.ShowDiskInfo()
	},
}

// ネットワーク情報用のコマンド
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "ネットワーク情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		collector.ShowNetworkInfo()
	},
}

// 全ての情報を収集するコマンド
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "全ての情報を収集",
	Run: func(cmd *cobra.Command, args []string) {
		collector.ShowAllInfo()
	},
}
