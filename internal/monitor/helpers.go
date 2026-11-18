package monitor

import "fmt"

// 1 GB in bytes
const GB = 1024 * 1024 * 1024

func toGB(val int) string {
	// Cast to float64 to allow for decimal division and satisfy the %f verb
	gb := float64(val) / float64(GB)

	// %.2f formats the float to 2 decimal places
	return fmt.Sprintf("%.2fGB", gb)
}
