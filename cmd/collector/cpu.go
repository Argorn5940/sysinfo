package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func ShowCPUInfo() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("CPU情報の取得に失敗しました: %v\n", err)
		return
	}
	fmt.Println("== CPU 情報 ==")
	for _, cpu := range cpuInfo {
		fmt.Printf("モデル名: %s\n", cpu.ModelName)
		fmt.Printf("コア数: %d\n", cpu.Cores)
		fmt.Printf("クロック周波数: %.2f MHz\n", cpu.Mhz)
	}
}
