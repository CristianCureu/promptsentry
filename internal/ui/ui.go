package ui

import (
	"fmt"
	"sync"

	"github.com/cristiancureu/prompt-sentry/internal/config"
	"github.com/cristiancureu/prompt-sentry/internal/scan"
	"github.com/pterm/pterm"
)

type ScannerState struct {
	mu         sync.Mutex
	scanned    int
	vulnerable int
	total      int
	pbar       *pterm.ProgressbarPrinter
}

func NewScannerState(total int, pbar *pterm.ProgressbarPrinter) *ScannerState {
	return &ScannerState{
		total: total,
		pbar:  pbar,
	}
}

func (s *ScannerState) PrintResult(result scan.ScanResult) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if result.Vulnerable {
		if result.Severity == "high" {
			pterm.Error.Printfln("[%s] %s", result.Severity, result.Reason)
		} else {
			pterm.Warning.Printfln("[%s] %s", result.Severity, result.Reason)
		}
		s.vulnerable++
	} else {
		pterm.Success.Printfln("%s", result.Reason)
	}
	s.scanned++
	s.pbar.Increment()
}

func (s *ScannerState) Summary(duration string, cfg *config.Config) {
	fmt.Printf("\n✅ Finished scanning %d prompts in %s\n", s.total, duration)
	fmt.Printf("⚠️  Vulnerabilities found: %d\n", s.vulnerable)
	fmt.Printf("\n✅ Full scan report saved to %v.\n", cfg.OutputFile)
}
