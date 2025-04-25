package nat

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// WiFiNetwork represents a detected WiFi network
type WiFiNetwork struct {
	SSID         string
	Signal       string
	SecurityType string
}

// EnableWiFi turns on the WiFi adapter (Admin required)
func EnableWiFi() error {
	cmd := exec.Command("cmd", "/C", "netsh interface set interface name=\"Wi-Fi\" admin=enabled")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to enable WiFi: %v", err)
	}
	fmt.Println("✅ WiFi enabled successfully.")
	return nil
}

// ScanWiFiNetworks scans for available WiFi networks
func ScanWiFiNetworks() ([]WiFiNetwork, error) {
	cmd := exec.Command("cmd", "/C", "netsh wlan show networks mode=bssid")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to scan WiFi networks: %v", err)
	}

	output := out.String()
	return parseWiFiNetworks(output), nil
}

// parseWiFiNetworks extracts WiFi details from the command output
func parseWiFiNetworks(output string) []WiFiNetwork {
	var networks []WiFiNetwork
	ssidRegex := regexp.MustCompile(`SSID\s+\d+\s+:\s(.+)`)
	signalRegex := regexp.MustCompile(`Signal\s+:\s+(\d+)%`)
	securityRegex := regexp.MustCompile(`Authentication\s+:\s+(.+)`)

	lines := strings.Split(output, "\n")
	var network WiFiNetwork

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if matches := ssidRegex.FindStringSubmatch(line); matches != nil {
			if network.SSID != "" {
				networks = append(networks, network) // Save previous network
			}
			network = WiFiNetwork{SSID: matches[1]} // Start new network
		} else if matches := signalRegex.FindStringSubmatch(line); matches != nil {
			network.Signal = matches[1] + "%"
		} else if matches := securityRegex.FindStringSubmatch(line); matches != nil {
			network.SecurityType = matches[1]
		}
	}

	if network.SSID != "" {
		networks = append(networks, network) // Save last network
	}

	return networks
}

// DisplayWiFiNetworks enables WiFi (if off) and prints available networks
func DisplayWiFiNetworks() {
	err := EnableWiFi()
	if err != nil {
		fmt.Println("⚠️ Could not enable WiFi automatically. Please enable it manually.")
	}

	networks, err := ScanWiFiNetworks()
	if err != nil {
		fmt.Println("❌ Error scanning WiFi networks:", err)
		return
	}

	if len(networks) == 0 {
		fmt.Println("⚠️ No WiFi networks found.")
		return
	}

	fmt.Println("📶 Available WiFi Networks:")
	for i, net := range networks {
		fmt.Printf("%d. SSID: %s | Signal: %s | Security: %s\n", i+1, net.SSID, net.Signal, net.SecurityType)
	}
}
