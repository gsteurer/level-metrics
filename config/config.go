package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gsteurer/level-metrics/input"
	"github.com/gsteurer/level-metrics/output"
)

const (
	DIRECTORY_PATH_DESCRIPTION   = "A directory containing files to be processed. Subdirectories are skipped."                                             // nolint
	INPUT_FILE_TYPE_DESCRIPTION  = "The input file type. Valid arguments are JSON and CSV."                                                                // nolint
	START_TIME_DESCRIPTION       = "An RFC3339 timestamp that filters out all records before this time."                                                   // nolint
	END_TIME_DESCRIPTION         = "An RFC3339 timestamp that filters out all records after this time."                                                    // nolint
	OUTPUT_FILE_TYPE_DESCRIPTION = "The summary results are written to an output file in this format. Valid arguments are JSON and YAML. Default is JSON." // nolint
	OUTPUT_FILE_NAME_DESCRIPTION = "The name of the file to which summary results are written. Default is current working directory."                      // nolint
)

type Config struct {
	DirectoryPath  string
	InputFileType  string
	StartTime      time.Time
	EndTime        time.Time
	OutputFileType string
	OutputFileName string
	startTimeRaw   string
	endTimeRaw     string
}

func (cfg *Config) parse() {

	flag.StringVar(&cfg.DirectoryPath, "directory", "", DIRECTORY_PATH_DESCRIPTION)
	flag.StringVar(&cfg.DirectoryPath, "d", "", DIRECTORY_PATH_DESCRIPTION)

	flag.StringVar(&cfg.InputFileType, "type", "", INPUT_FILE_TYPE_DESCRIPTION)
	flag.StringVar(&cfg.InputFileType, "t", "", INPUT_FILE_TYPE_DESCRIPTION)

	flag.StringVar(&cfg.startTimeRaw, "startTime", "", START_TIME_DESCRIPTION)

	flag.StringVar(&cfg.endTimeRaw, "endTime", "", END_TIME_DESCRIPTION)

	flag.StringVar(&cfg.OutputFileType, "outputFileType", string(output.JSON), OUTPUT_FILE_TYPE_DESCRIPTION)

	flag.StringVar(&cfg.OutputFileName, "outputFileName", "", OUTPUT_FILE_NAME_DESCRIPTION)

	flag.Parse()

}

func (cfg *Config) polish() error {
	var err error
	cfg.StartTime, err = time.Parse(time.RFC3339, cfg.startTimeRaw)
	if err != nil {
		return fmt.Errorf("could not parse startTime: %s ", err.Error())
	}

	cfg.EndTime, err = time.Parse(time.RFC3339, cfg.endTimeRaw)
	if err != nil {
		return fmt.Errorf("could not parse endTime: %s", err.Error())
	}

	if cfg.OutputFileName == "" {
		cfg.OutputFileName = "./out." + string(cfg.OutputFileType)
	}

	return nil
}

func (cfg *Config) validate() error {

	if cfg.startTimeRaw == "" {
		return fmt.Errorf("a valid RFC3339 timestamp is required for startTime.")
	}

	if cfg.endTimeRaw == "" {
		return fmt.Errorf("a valid RFC3339 timestamp is required for endTime.")
	}

	if cfg.DirectoryPath == "" || cfg.DirectoryPath == "/" {
		return fmt.Errorf("invalid input directory '%s'", cfg.DirectoryPath)
	}

	if _, err := os.Stat(cfg.DirectoryPath); err != nil {
		return fmt.Errorf("parsing input directory '%s': %s", cfg.DirectoryPath, err)
	}

	if output.OutputType(cfg.OutputFileType) != output.JSON && output.OutputType(cfg.OutputFileType) != output.YAML {
		return fmt.Errorf("invalid output type '%s'", cfg.OutputFileType)
	}

	if input.InputType(cfg.InputFileType) != input.JSON && input.InputType(cfg.InputFileType) != input.CSV {
		return fmt.Errorf("invalid input type '%s'", cfg.OutputFileType)
	}

	return nil
}

func Create() (Config, error) {

	cfg := Config{}
	cfg.parse()

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	if err := cfg.polish(); err != nil {
		return cfg, err
	}

	return cfg, nil
}
