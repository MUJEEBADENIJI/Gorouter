package nat

import (
	"fmt"
	"os/exec"
)

// StartHotspot starts the Windows hotspot using netsh.
func StartHotspot() error {
	cmd := exec.Command("netsh", "wlan", "set", "hostednetwork", "mode=allow", "ssid=GoRouteHotspot", "key=yourpassword")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set hotspot: %w", err)
	}

	startCmd := exec.Command("netsh", "wlan", "start", "hostednetwork")
	err = startCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start hotspot: %w", err)
	}

	return nil
}

// StopHotspot stops the Windows hotspot.
func StopHotspot() error {
	stopCmd := exec.Command("netsh", "wlan", "stop", "hostednetwork")
	err := stopCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to stop hotspot: %w", err)
	}
	return nil
}
