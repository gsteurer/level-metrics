package ingest

import (
	"strings"
	"testing"
	"time"

	"github.com/gsteurer/level-metrics/filter"
	"github.com/gsteurer/level-metrics/model/record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngestCSV(t *testing.T) {

	t.Run("CSV reader can process records", func(t *testing.T) {

		m := mockProcess{Records: []record.Record{}}

		s := "timestamp,level_name,value\n2022-03-01T12:00:00.00Z,bar,3\n2022-04-02T12:00:00.00Z,foo,9"
		reader := strings.NewReader(s)
		t0, _ := time.Parse(time.RFC3339, "2022-04-01T12:00:00.00Z")
		tf, _ := time.Parse(time.RFC3339, "2022-04-29T12:00:00.00Z")
		f := filter.Filter{StartTime: t0, EndTime: tf}

		err := readCSVFile(reader, f, m.Process)

		require.NoError(t, err)
		assert.Equal(t, 1, len(m.Records))
		assert.Equal(t, 9, m.Records[0].Value)
		assert.Equal(t, "foo", m.Records[0].LevelName)
	})
}
