package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("config is created successfully", func(t *testing.T) {
		cfg := Config{
			DirectoryPath:  "/tmp", // this will fail on windows
			InputFileType:  "json",
			startTimeRaw:   "2022-04-01T12:00:00.00Z",
			endTimeRaw:     "2022-04-29T12:00:00.00Z",
			OutputFileType: "json",
			OutputFileName: "/tmp/out.json",
		}

		err := cfg.polish()
		require.NoError(t, err)

		err = cfg.validate()
		require.NoError(t, err)
	})

	t.Run("config validations", func(t *testing.T) {
		var err error
		expected := Config{
			DirectoryPath:  "/tmp",
			InputFileType:  "json",
			startTimeRaw:   "2022-04-01T12:00:00.00Z",
			endTimeRaw:     "2022-04-29T12:00:00.00Z",
			OutputFileType: "json",
			OutputFileName: "/tmp/out.json",
		}

		// error if input directory does not exist
		test := expected
		test.DirectoryPath = ""
		_ = test.polish()
		err = test.validate()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input directory")

		// empty start time is invalid
		test = expected
		test.startTimeRaw = ""
		_ = test.polish()
		err = test.validate()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "a valid RFC3339 timestamp is required for")

		// empty end time is invalid
		test = expected
		test.endTimeRaw = ""
		_ = test.polish()
		err = test.validate()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "a valid RFC3339 timestamp is required for")

		// empty output file type is invalid
		test = expected
		test.OutputFileType = ""
		_ = test.polish()
		err = test.validate()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid output type")

		// empty input file type is invalid
		test = expected
		test.InputFileType = ""
		_ = test.polish()
		err = test.validate()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input type")

	})
}
