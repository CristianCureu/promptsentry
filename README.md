# PromptSentry

A robust scanner for detecting prompt injection vulnerabilities in LLM endpoints.

[![Go Report Card](https://goreportcard.com/badge/github.com/cristiancureu/prompt-sentry)](https://goreportcard.com/report/github.com/cristiancureu/prompt-sentry)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

PromptSentry is a command-line tool designed to scan Large Language Model (LLM) endpoints for various security vulnerabilities, with a focus on:

- Prompt injection attacks
- System prompt leaks
- LLM policy bypasses
- Unsafe prompt handling

The tool works by sending crafted prompts to target LLM APIs and analyzing the responses for signs of vulnerability.

## Features

- **Comprehensive Scanning**: Tests endpoints against a variety of prompt injection techniques
- **Parallel Processing**: Option to run scans concurrently for faster results
- **Flexible Output**: Generate reports in multiple formats (console, JSON, CSV)
- **Confidence Scoring**: Each vulnerability is rated with confidence and severity metrics
- **Pattern Recognition**: Identifies common indicators of vulnerabilities in responses

## Installation

### From Source

```bash
git clone https://github.com/cristiancureu/prompt-sentry.git
cd prompt-sentry
go build -o promptsentry .
```

### Using Go

```bash
go install github.com/cristiancureu/prompt-sentry@latest
```

## Usage

Basic scan against an LLM endpoint:

```bash
./promptsentry scan --target="http://localhost:11434" --apikey="your-api-key"
```

### Command Line Options

```
Usage:
  promptsentry scan [flags]

Flags:
      --target string     target URL (required)
      --apikey string     api key for authentication
      --output string     output file (default "report.json")
      --format string     output format: console | csv | json (default "console")
      --parallel          run scan in parallel mode
```

## Examples

### Basic Scan

```bash
./promptsentry scan --target="http://localhost:11434"
```

### Parallel Scan with JSON Output

```bash
./promptsentry scan --target="http://localhost:11434" --parallel --format=json --output=vulnerabilities.json
```

### Scan with API Key Authentication

```bash
./promptsentry scan --target="https://api.example.com/v1" --apikey="sk-yourapikeyhere" --format=csv
```

## How It Works

PromptSentry works in several stages:

1. **Prompt Loading**: Loads a set of carefully crafted test prompts designed to elicit vulnerable behavior
2. **Request Execution**: Sends each prompt to the target LLM endpoint
3. **Response Analysis**: Applies a set of rules to detect patterns indicating vulnerabilities
4. **Report Generation**: Compiles findings into a structured report

## Rule-Based Detection

The scanner uses a rule-based system to classify vulnerabilities:

- **System Prompt Leak**: Detects when an LLM reveals its system instructions
- **Policy Bypass**: Identifies when safety guidelines are circumvented
- **Unsafe Prompt Handling**: Detects when the LLM should have refused but didn't
- **Evasive Responses**: Identifies ambiguous responses to dangerous prompts

## Contribution

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Security Considerations

PromptSentry is designed for security testing by authorized personnel. Always ensure you have permission to test any LLM endpoints. The tool should be used responsibly and ethically.

## Disclaimer

This tool is provided for educational and legitimate security testing purposes only. The authors and contributors are not responsible for any misuse or damage caused by this tool.