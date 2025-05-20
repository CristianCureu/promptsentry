package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

type OutputFormat string

const (
	FormatCSV  OutputFormat = "csv"
	FormatJSON OutputFormat = "json"
	FormatNone OutputFormat = "none"
)

type Config struct {
	TargetURL  string
	APIKey     string
	OutputFile string
	Format     OutputFormat
	Parallel   bool
	API        bool
}

func LoadConfig(cmd *cobra.Command) (Config, error) {
	target, err := cmd.Flags().GetString("target")
	if err != nil {
		return Config{}, fmt.Errorf("error reading target: %v", err)
	}
	format, _ := cmd.Flags().GetString("format")
	apikey, _ := cmd.Flags().GetString("apikey")
	output, _ := cmd.Flags().GetString("output")
	parallel, _ := cmd.Flags().GetBool("parallel")

	return Config{
		TargetURL:  target,
		APIKey:     apikey,
		OutputFile: output,
		Format:     OutputFormat(format),
		Parallel:   parallel,
	}, nil
}
