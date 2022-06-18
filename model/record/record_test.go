package record

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecord(t *testing.T) {
	// initial state
	t1, _ := time.Parse(time.RFC3339, "2022-01-01T6:00:00.00Z")
	expected := Record{
		LevelName: "foo",
		Value:     3,
		Timestamp: t1,
	}

	body := string(`{"level_name": "foo", "value": 3, "timestamp": "2022-01-01T6:00:00.00Z"}`)
	row := "2022-01-01T6:00:00.00Z,foo,3"

	// tests
	t.Run("a record is successfully created from a JSON body", func(t *testing.T) {

		actual, err := FromJSON(body)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("a record is successfully created from a CSV row", func(t *testing.T) {

		actual, err := FromCSV(row)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

}
