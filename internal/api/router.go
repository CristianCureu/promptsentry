package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cristiancureu/prompt-sentry/internal/config"
	"github.com/cristiancureu/prompt-sentry/internal/prompts"
	"github.com/cristiancureu/prompt-sentry/internal/scan"
	"github.com/cristiancureu/prompt-sentry/internal/ui"
	"github.com/pterm/pterm"
)

type ScanRequest struct {
	TargetURL string `json:"target"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/scan", handleScan)
	return mux
}

func (req *ScanRequest) Validate() error {
	if req.TargetURL == "" {
		return fmt.Errorf("target is required")
	}
	return nil
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg := &config.Config{
		TargetURL: req.TargetURL,
	}

	scanner := scan.NewScanner(cfg)
	prompts := prompts.LoadPromptSet()
	total := len(prompts)
	pbar, _ := pterm.DefaultProgressbar.
		WithTotal(total).
		WithTitle("Scanning...").
		WithRemoveWhenDone(true).
		Start()
	state := ui.NewScannerState(total, pbar)

	results := scanner.StartScan(prompts, state.PrintResult)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
