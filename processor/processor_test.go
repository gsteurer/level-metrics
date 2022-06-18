package processor

import (
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/go-yaml/yaml"
	"github.com/gsteurer/level-metrics/model/record"
	"github.com/gsteurer/level-metrics/model/summary"
	"github.com/gsteurer/level-metrics/output"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessor(t *testing.T) {
	expected := []summary.Summary{
		{
			LevelName:  "foo",
			TotalValue: 9,
		},
		{
			LevelName:  "bar",
			TotalValue: 6,
		},
	}

	// sort to guarantees the order of the list as long as the TotalValue values are unique
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].TotalValue < expected[j].TotalValue
	})

	r1 := record.Record{LevelName: "foo", Value: 4}
	r2 := record.Record{LevelName: "foo", Value: 5}
	r3 := record.Record{LevelName: "bar", Value: 6}
	p := Create()
	p.Process(r1)
	p.Process(r2)
	p.Process(r3)

	t.Run("expected summary, YAML", func(t *testing.T) {

		// summary is successful
		s, err := p.Summarize(output.YAML)
		require.NoError(t, err)

		// the result of summarize is usable YAML
		var actual []summary.Summary
		err = yaml.Unmarshal(s, &actual)
		require.NoError(t, err)

		sort.Slice(actual, func(i, j int) bool {
			return actual[i].TotalValue < actual[j].TotalValue
		})

		// and the expected results match the actual results
		assert.Equal(t, expected, actual)

	})

	t.Run("expected summary, JSON", func(t *testing.T) {

		// summary is successful
		s, err := p.Summarize(output.JSON)
		require.NoError(t, err)

		// the result of summarize is usable JSON
		var actual []summary.Summary
		err = json.Unmarshal(s, &actual)
		require.NoError(t, err)

		sort.Slice(actual, func(i, j int) bool {
			return actual[i].TotalValue < actual[j].TotalValue
		})

		// and the expected results match the actual results
		assert.Equal(t, expected, actual)

	})

	t.Run("expected summary, human readable", func(t *testing.T) {

		// summary is successful
		actual, err := p.Summarize(output.HUMAN_READABLE)
		require.NoError(t, err)

		// and the expected results match the actual results
		assert.True(t, strings.Contains(string(actual), "foo"))
		assert.True(t, strings.Contains(string(actual), "9"))
		assert.True(t, strings.Contains(string(actual), "bar"))
		assert.True(t, strings.Contains(string(actual), "6"))

	})

}
