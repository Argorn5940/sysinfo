package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func ShowMemoryInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("メモリ情報の取得に失敗しました: %v\n", err)
		return
	}
	fmt.Println("===メモリ情報===")
	fmt.Printf("総メモリ: %.2f GB\n", float64(memInfo.Total)/(1024*1024*1024))
	fmt.Printf("使用中: %.2f GB\n", float64(memInfo.Used)/(1024*1024*1024))
	fmt.Printf("空き: %.2f GB\n", float64(memInfo.Free)/(1024*1024*1024))
	fmt.Printf("使用率: %.2f%%\n", memInfo.UsedPercent)
}
