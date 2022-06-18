package processor

import (
	"encoding/json"
	"fmt"

	"github.com/go-yaml/yaml"
	"github.com/gsteurer/level-metrics/model/record"
	"github.com/gsteurer/level-metrics/model/summary"
	"github.com/gsteurer/level-metrics/output"
)

type ProcessFn func(r record.Record)

type Processor interface {
	Process(r record.Record)
	Summarize(outputType output.OutputType) ([]byte, error)
}
type processor struct {
	results map[string]int
}

// allocates a processor
func Create() Processor {
	return &processor{results: map[string]int{}}
}

// updates state of processor with a record
func (p *processor) Process(r record.Record) {
	p.results[r.LevelName] += r.Value
}

// produces output of a given type such as json, yaml etc
func (p *processor) Summarize(outputType output.OutputType) ([]byte, error) {
	items := []summary.Summary{}
	for levelName, totalValue := range p.results {
		items = append(items, summary.Summary{LevelName: levelName, TotalValue: totalValue})
	}
	switch outputType {
	case output.JSON:
		return json.Marshal(items)
	case output.YAML:
		return yaml.Marshal(items)
	case output.HUMAN_READABLE:
		s := ""
		for idx, item := range items {
			s += item.ToString()
			if idx < len(items)-1 {
				s += "\n"
			}
		}
		return []byte(s), nil

	default:
		return []byte(""), fmt.Errorf("unhandled output type '%s'", outputType)
	}

}
