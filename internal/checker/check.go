package checker

import (
	"errors"
	"fmt"
	"gowatcher_g3/config"
	"net/http"
	"time"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

type ReportEntry struct {
	Name   string
	URl    string
	Owner  string
	Status string
	ErrMsg string // Omis si vide
}

func CheckURL(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target,
			Err: &UnreachableURLError{
				URL: target.URL,
				Err: err,
			},
		}
	}
	defer resp.Body.Close()

	return CheckResult{
		InputTarget: target,
		Status:      resp.Status,
	}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   res.InputTarget.Name,
		URl:    res.InputTarget.URL,
		Owner:  res.InputTarget.Owner,
		Status: res.Status,
	}

	if res.Err != nil {
		var unreachable *UnreachableURLError
		if errors.As(res.Err, &unreachable) {
			report.Status = "Inaccessible"
			report.ErrMsg = fmt.Sprintf("Unreachable URL: %v", unreachable.Err)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("Erreur générique: %v", res.Err)
		}
	}
	return report
}
