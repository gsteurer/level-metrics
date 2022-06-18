package ingest

import (
	"github.com/gsteurer/level-metrics/model/record"
)

type mockProcess struct {
	Records []record.Record
}

func (m *mockProcess) Process(r record.Record) {
	m.Records = append(m.Records, r)
}
