package output

type OutputType string

const (
	JSON           OutputType = "json"
	YAML           OutputType = "yaml"
	HUMAN_READABLE OutputType = "human_readable"
)
