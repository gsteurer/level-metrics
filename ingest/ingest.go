package ingest

import "github.com/gsteurer/level-metrics/processor"

type Ingest interface {
	Run(processor.ProcessFn) error
}
