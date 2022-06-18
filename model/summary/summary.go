package summary

import "fmt"

type Summary struct {
	LevelName  string `json:"level_name"`
	TotalValue int    `json:"total_value"`
}

func (s *Summary) ToString() string {
	return fmt.Sprintf("Level Name: %s\n\tTotal Value: %v", s.LevelName, s.TotalValue)
}
