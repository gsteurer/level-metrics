package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gsteurer/level-metrics/config"
	"github.com/gsteurer/level-metrics/ingest"
	"github.com/gsteurer/level-metrics/output"
	iprocessor "github.com/gsteurer/level-metrics/processor"
)

func main() {
	// parse config
	cfg, err := config.Create()
	if err != nil {
		log.Fatal(fmt.Errorf("reading config: %s", err))
	}

	// read and process metrics from files
	processor := iprocessor.Create()
	metricIngest := ingest.NewFileIngest(cfg)

	if err := metricIngest.Run(processor.Process); err != nil {
		log.Fatal(fmt.Errorf("ingesting data: %s", err))
	}

	// summarize data and write to file
	levelSummary, err := processor.Summarize(output.OutputType(cfg.OutputFileType))
	if err != nil {
		log.Fatal(fmt.Errorf("summarizing metrics: %s", err))
	}

	if err := os.WriteFile(cfg.OutputFileName, append(levelSummary, []byte("\n")...), 0644); err != nil {
		log.Fatal(fmt.Errorf("writing output to file: %s", err))
	}

	// display output to terminal
	if display, err := processor.Summarize(output.HUMAN_READABLE); err != nil {
		log.Fatal(fmt.Errorf("writing output to terminal: %s", err))
	} else {
		fmt.Printf("%s\n", display)
	}
}
