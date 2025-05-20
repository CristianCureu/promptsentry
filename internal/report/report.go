package report

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cristiancureu/prompt-sentry/internal/config"
	"github.com/cristiancureu/prompt-sentry/internal/scan"
)

func GenerateCSVReport(results []scan.ScanResult, cfg *config.Config) {
	file, err := os.Create(cfg.OutputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Prompt", "Response", "Vulnerable", "Reason", "Confidence"})
	for _, r := range results {
		writer.Write([]string{r.Prompt, r.Response, fmt.Sprintf("%v", r.Vulnerable), r.Reason, r.Confidence})
	}
}

func GenerateJSONReport(results []scan.ScanResult, cfg *config.Config) {
	file, err := os.Create(cfg.OutputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		panic(err)
	}
}
