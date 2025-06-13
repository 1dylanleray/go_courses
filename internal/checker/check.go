package checker

import (
	"gowatcher_g3/config"
	"net/http"
	"time"
)

// CheckResult en majuscule pour exporter le type
type CheckResult struct {
	InputTarget string
	Status      string
	Err         error
}

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string
	ErrMsg string
}

func CheckURL(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target.URL,
			Err:         &UnreachableURLError{URL: target.URL, Err: err},
		}

	}
	defer resp.Body.Close()

	return CheckResult{
		InputTarget: target.URL,
		Status:      resp.Status,
	}
}
