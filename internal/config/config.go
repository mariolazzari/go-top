package config

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Mode     string
	Interval time.Duration
}

// ParseFlags parses command-line arguments and validates them.
func ParseFlags() (*Config, error) {
	mode := flag.String("mode", "text", "Operational mode: 'text', 'web', or 'tui'")
	interval := flag.Duration("interval", time.Second, "Metrics collection interval (e.g., 500ms, 1s)")
	flag.Parse()

	// Validate mode
	normalizedMode := strings.ToLower(strings.TrimSpace(*mode))
	switch normalizedMode {
	case "text", "web", "tui":
		// Valid modes
	default:
		return nil, fmt.Errorf("invalid mode %q. Supported modes are 'text', 'web', 'tui'", *mode)
	}

	// Validate interval
	if *interval <= 0 {
		return nil, fmt.Errorf("interval must be greater than zero, got %v", *interval)
	}

	return &Config{
		Mode:     normalizedMode,
		Interval: *interval,
	}, nil
}
