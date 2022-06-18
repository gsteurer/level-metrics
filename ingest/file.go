package ingest

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gsteurer/level-metrics/config"
	"github.com/gsteurer/level-metrics/filter"
	"github.com/gsteurer/level-metrics/input"
	iprocessor "github.com/gsteurer/level-metrics/processor"
)

type fileIngest struct {
	cfg config.Config
}

func NewFileIngest(cfg config.Config) Ingest {
	return &fileIngest{cfg: cfg}
}

func (f *fileIngest) Run(process iprocessor.ProcessFn) error {
	directoryPath := f.cfg.DirectoryPath
	inputType := f.cfg.InputFileType
	filter := filter.Create(f.cfg)

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			log.Printf("'%s' is a directory", directoryPath)
			continue
		}

		f, err := os.Open(filepath.Join(directoryPath, file.Name()))
		if err != nil {
			return err
		}

		switch input.InputType(inputType) {
		case input.CSV:
			if err := readCSVFile(f, filter, process); err != nil {
				return err
			}
		case input.JSON:
			if err := readJSONFile(f, filter, process); err != nil {
				return err
			}

		}
	}

	return nil
}
