package main

import (
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/cristiancureu/prompt-sentry/internal/config"
	"github.com/cristiancureu/prompt-sentry/internal/prompts"
	"github.com/cristiancureu/prompt-sentry/internal/report"
	"github.com/cristiancureu/prompt-sentry/internal/scan"
	"github.com/cristiancureu/prompt-sentry/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "promptsentry",
	Short: "Prompt injection scanner for LLM endpoints",
	Long:  `PromptSentry scans LLMs for prompt injection vulnerabilities and system prompt leaks.`,
}

func Execute() error {
	cfg := config.Config{}

	rootCmd.AddCommand(NewScanCmd(&cfg))

	return rootCmd.Execute()
}

func NewScanCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Prompt injection scanner for LLM endpoints",
		Long:  `PromptSentry scans LLMs for prompt injection vulnerabilities and system prompt leaks.`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig(cmd)
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}

			scanner := scan.NewScanner(&cfg)
			prompts := prompts.LoadPromptSet()
			total := len(prompts)

			s := spinner.New(spinner.CharSets[14], 120*time.Millisecond)
			s.Suffix = " Scanning prompts for vulnerabilities...\n"
			s.Start()

			pbar, _ := pterm.DefaultProgressbar.
				WithTotal(total).
				WithTitle("Scanning...").
				WithRemoveWhenDone(true).
				Start()

			state := ui.NewScannerState(total, pbar)

			start := time.Now()
			var results []scan.ScanResult

			s.Stop()

			if cfg.Parallel {
				results = scanner.StartParallelScan(prompts, state.PrintResult)
			} else {
				results = scanner.StartScan(prompts, state.PrintResult)
			}

			duration := time.Since(start)

			state.Summary(duration.String(), &cfg)

			switch cfg.Format {
			case "json":
				report.GenerateJSONReport(results, &cfg)
			case "csv":
				report.GenerateCSVReport(results, &cfg)
			default:
				report.GenerateJSONReport(results, &cfg)
			}

		},
	}

	cmd.Flags().String("target", "", "target URL")
	cmd.Flags().String("apikey", "", "api key")
	cmd.Flags().String("output", "report.json", "output value")
	cmd.Flags().String("format", "console", "output format: console | csv | json")
	cmd.Flags().Bool("parallel", false, "run in parallel")

	_ = cmd.MarkFlagRequired("target")

	return cmd
}
