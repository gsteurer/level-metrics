package ingest

import (
	"bufio"
	"io"

	"github.com/gsteurer/level-metrics/filter"
	"github.com/gsteurer/level-metrics/model/record"
	iprocessor "github.com/gsteurer/level-metrics/processor"
)

// getJSONRecord provides a mechaism for reading a list of JSON objects from an array
// we avoid reading the whole file into memory by reading one object at a time and moving on
func getJSONRecord(reader *bufio.Reader) (string, error) {

	start := "{"
	end := "}"
	beginRecording := false
	result := ""
	counters := map[string]int{}
	for {
		if rune, _, err := reader.ReadRune(); err == nil {

			s := string(rune)
			if s == start {
				beginRecording = true
				counters[start]++
			} else if s == end {
				counters[end]++

				if counters[start] == counters[end] {
					result += s
					return result, nil
				}
			}

			if beginRecording {
				result += s
			}

		} else {
			return result, err
		}
	}

}

func readJSONFile(file io.Reader, recordFilter filter.Filter, process iprocessor.ProcessFn) error {
	rdr := bufio.NewReader(file)

	for {
		bar, err := getJSONRecord(rdr)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		r, err := record.FromJSON(bar)
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

	return nil
}
