package scan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cristiancureu/prompt-sentry/internal/config"
)

type Scanner struct {
	Client *http.Client
	Config *config.Config
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func NewScanner(cfg *config.Config) *Scanner {
	return &Scanner{
		Client: &http.Client{},
		Config: cfg,
	}
}

func (s *Scanner) StartScan(cfg *config.Config, prompts []string) []ScanResult {
	var results []ScanResult

	for _, prompt := range prompts {
		ollamaResp, err := s.SendPrompt(prompt)
		if err != nil {
			fmt.Printf("Error scanning prompt: %v\n", err)
			continue
		}

		scanResult := AnalyzeResponse(prompt, ollamaResp.Response)
		results = append(results, scanResult)
	}

	return results
}

func (s *Scanner) SendPrompt(prompt string) (OllamaResponse, error) {
	reqBody := OllamaRequest{
		Model:  "tinyllama",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return OllamaResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.Client.Post(s.Config.TargetURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return OllamaResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OllamaResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	var result OllamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return OllamaResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}
