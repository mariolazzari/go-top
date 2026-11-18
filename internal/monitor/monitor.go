package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func New(interval time.Duration) (*Monitor, error) {
	if interval <= 0 {
		return nil, fmt.Errorf("interval must be greater than zero")
	}

	return &Monitor{
		interval: interval,
	}, nil
}

func (m *Monitor) run(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}

		percentages, err := cpu.PercentWithContext(
			ctx,
			m.interval,
			false,
		)

		if err != nil {
			if ctx.Err() != nil {
				return
			}

			m.mu.Lock()
			m.snapshot.LastError = err.Error()
			m.mu.Unlock()

			continue
		}

		if len(percentages) == 0 {
			continue
		}

		m.mu.Lock()
		m.snapshot.CPUPercent = percentages[0]
		m.snapshot.UpdatedAt = time.Now()
		m.snapshot.LastError = ""
		m.mu.Unlock()
	}
}

// Start launches the background metrics collector.
func (m *Monitor) Start(ctx context.Context) {
	m.started.Do(func() {
		go m.run(ctx)
	})
}

func (m *Monitor) Snapshot() Snapshot {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.snapshot
}
