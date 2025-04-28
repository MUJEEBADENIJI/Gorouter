package bandwidth_allocator

import (
	"fmt"
)

// StartAllocation manages bandwidth distribution
func StartAllocation() {
	fmt.Println("📊 Bandwidth Allocation Started")

	allocateBandwidth("192.168.1.2", "10Mbps")
}

func allocateBandwidth(ip string, limit string) {
	fmt.Printf("🔹 Allocating %s to %s\n", limit, ip)
}
