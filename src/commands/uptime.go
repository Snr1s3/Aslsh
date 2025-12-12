package commands

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

func Uptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	days := uptime / (24 * 3600)
	hours := (uptime % (24 * 3600)) / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60
	return fmt.Sprintf("Uptime: %d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
}
