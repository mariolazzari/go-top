package monitor

import (
	"sync"
	"time"
)

type Snapshot struct {
	CPUPercent float64
	UpdatedAt  time.Time
	LastError  string
}

type Monitor struct {
	interval time.Duration
	mu       sync.RWMutex
	snapshot Snapshot
	started  sync.Once
}
