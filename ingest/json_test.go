package ingest

import (
	"bufio"
	"strings"
	"testing"
	"time"

	"github.com/gsteurer/level-metrics/filter"
	"github.com/gsteurer/level-metrics/model/record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngestJSON(t *testing.T) {

	t.Run("getJSONrecord can parse a single JSON object from a list", func(t *testing.T) {
		s := `[ {"timestamp": "2022-03-01T12:00:00.00Z", "level_name": "bar", "value": 3 }]`
		reader := strings.NewReader(s)
		body, err := getJSONRecord(bufio.NewReader(reader))
		require.NoError(t, err)
		r, err := record.FromJSON(body)
		require.NoError(t, err)
		assert.Equal(t, 3, r.Value)
		assert.Equal(t, "bar", r.LevelName)
	})

	t.Run("JSON reader can process records", func(t *testing.T) {
		m := mockProcess{Records: []record.Record{}}

		s := `[ {"timestamp": "2022-03-01T12:00:00.00Z", "level_name": "bar", "value": 3 }, {"timestamp": "2022-04-02T12:00:00.00Z", "level_name": "foo", "value": 9 }]`
		reader := strings.NewReader(s)
		t0, _ := time.Parse(time.RFC3339, "2022-04-01T12:00:00.00Z")
		tf, _ := time.Parse(time.RFC3339, "2022-04-29T12:00:00.00Z")
		f := filter.Filter{StartTime: t0, EndTime: tf}

		err := readJSONFile(reader, f, m.Process)

		require.NoError(t, err)
		assert.Equal(t, 1, len(m.Records))
		assert.Equal(t, 9, m.Records[0].Value)
		assert.Equal(t, "foo", m.Records[0].LevelName)
	})
}
