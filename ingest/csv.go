package ingest

import (
	"bufio"
	"io"
	"strings"

	"github.com/gsteurer/level-metrics/filter"
	"github.com/gsteurer/level-metrics/model/record"
	iprocessor "github.com/gsteurer/level-metrics/processor"
)

func readCSVFile(file io.Reader, recordFilter filter.Filter, process iprocessor.ProcessFn) error {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		// skip header and empty lines
		if strings.Contains(row, "timestamp") || strings.Trim(row, "\n") == "" {
			continue
		}

		r, err := record.FromCSV(row)
		if err != nil {
			return err
		}

		if r.Timestamp.After(recordFilter.StartTime) && r.Timestamp.Before(recordFilter.EndTime) {
			process(r)
		}

		// we can jump out of the file when we encounter timestamps after the specified entime
		//   * this happens if the first row in the file is after the end time
		//   * when we've started reading records that exceed the end time
		if r.Timestamp.After(recordFilter.EndTime) {
			break
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
