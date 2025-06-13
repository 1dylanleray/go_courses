package reporter

import (
	"encoding/json"
	"gowatcher_g3/internal/checker"
	"os"
)

func ExportResultsToJsonfile(filepath string, results []checker.ReportEntry) error {
	data, err := json.MarshalIndent(results, "", "")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return err
	}
	return nil
}
