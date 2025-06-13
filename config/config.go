package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Owner string `json:"owner"`
}

func LoadTargetsFromFile(filepath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file %s : %w", filepath, err)
	}

	var targets []InputTarget
	if err = json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal file %s : %w", filepath, err)
	}
	return targets, nil
}

func SaveTargetsToFile(filepath string, targets []InputTarget) error {
	data, err := json.MarshalIndent(targets, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to umarshal %s : %w", filepath, err)
	}
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("Failed to write %s : %w", filepath, err)
	}
	return nil
}
