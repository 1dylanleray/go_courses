package checker

import (
	"net/http"
	"time"
)

// CheckResult en majuscule pour exporter le type
type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckURL(url string) CheckResult {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			Target: url,
			Err:    &UnreachableURLError{URL: url, Err: err},
		}

	}
	defer resp.Body.Close()

	return CheckResult{
		Target: url,
		Status: resp.Status,
	}
}
