package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func ShowDiskInfo() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("ディスク情報の取得に失敗しました: %v\n", err)
		return
	}

	fmt.Println("=== ディスク情報 ===")
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("パーティション %s の情報取得に失敗: %v\n", partition.Mountpoint, err)
			continue
		}

		fmt.Printf("\nマウントポイント: %s\n", partition.Mountpoint)
		fmt.Printf("ファイルシステム: %s\n", partition.Fstype)
		fmt.Printf("総容量: %.2f GB\n", float64(usage.Total)/(1024*1024*1024))
		fmt.Printf("使用中: %.2f GB\n", float64(usage.Used)/(1024*1024*1024))
		fmt.Printf("空き: %.2f GB\n", float64(usage.Free)/(1024*1024*1024))
		fmt.Printf("使用中: %.2f%%\n", usage.UsedPercent)
	}

}
