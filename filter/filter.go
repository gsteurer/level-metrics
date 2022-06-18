package filter

import (
	"time"

	"github.com/gsteurer/level-metrics/config"
)

type Filter struct {
	StartTime time.Time
	EndTime   time.Time
}

func Create(cfg config.Config) Filter {
	return Filter{
		StartTime: cfg.StartTime,
		EndTime:   cfg.EndTime,
	}
}
