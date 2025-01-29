package collector

import (
	"fmt"
)

func ShowAllInfo() {
	fmt.Println("\n============システム情報=========")

	fmt.Println("\n[CPU情報]")
	ShowCPUInfo()

	fmt.Println("\n[メモリ情報]")
	ShowMemoryInfo()

	fmt.Println("\n[ディスク情報]")
	ShowDiskInfo()

	fmt.Println("\n[ネットワーク情報]")
	ShowNetworkInfo()
}
