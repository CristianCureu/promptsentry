package main

import (
	"log"

	"github.com/cristiancureu/prompt-sentry/internal/config"
	"github.com/cristiancureu/prompt-sentry/internal/prompts"
	"github.com/cristiancureu/prompt-sentry/internal/report"
	"github.com/cristiancureu/prompt-sentry/internal/scan"
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
			results := scanner.StartScan(&cfg, prompts)

			if cfg.OutputFile == "" {
				report.PrintToConsole(results)
			} else {
				switch cfg.Format {
				case "json":
					report.GenerateJSONReport(results, &cfg)
				case "csv":
					report.GenerateCSVReport(results, &cfg)
				default:
					log.Fatalf("Invalid format '%s' for file output. Use 'csv' or 'json'.", cfg.Format)
				}
			}

		},
	}

	cmd.Flags().String("target", "", "target URL")
	cmd.Flags().String("apikey", "", "api key")
	cmd.Flags().String("output", "", "output value")
	cmd.Flags().String("format", "console", "output format: console | csv | json")

	_ = cmd.MarkFlagRequired("target")

	return cmd
}
