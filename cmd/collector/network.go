package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

func ShowNetworkInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("ネットワーク情報の取得に失敗しました: %v\n", err)
		return
	}
	fmt.Println("=== ネットワーク情報 ===")
	for _, iface := range interfaces {
		fmt.Printf("\nインターフェース: %s\n", iface.Name)
		fmt.Printf("MACアドレス: %s\n", iface.HardwareAddr)
		fmt.Printf("状態: %v\n", iface.Flags)

		for _, addr := range iface.Addrs {
			fmt.Printf("IPアドレス: %s\n", addr.Addr)
		}
	}

	//ネットワーク統計情報の取得
	stats, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("ネットワーク情報の取得に失敗しました: %v\n", err)
		return
	}
	fmt.Println("=== ネットワーク統計 ===")
	for _, stat := range stats {
		fmt.Printf("\nインターフェース: %s\n", stat.Name)
		fmt.Printf("送信バイト数: %.2f MB\n", float64(stat.BytesSent)/(1024*1024))
		fmt.Printf("受信バイト数: %.2f MB\n", float64(stat.BytesRecv)/(1024*1024))
		fmt.Printf("送信パケット数: %d\n", stat.PacketsSent)
		fmt.Printf("受信パケット数: %d\n", stat.PacketsRecv)
	}
}
