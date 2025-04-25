package traffic_monitor

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// StartMonitoring captures network traffic
func StartMonitoring() {
	fmt.Println("📡 Traffic Monitoring Started")

	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("❌ Failed to open network interface:", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println("📦 Captured Packet:", packet)
	}
}
