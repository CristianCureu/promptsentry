package main

import "testing"

func TestMainCommand(t *testing.T) {
	if rootCmd.Use != "promptsentry" {
		t.Fatal("rootCmd not initialized properly")
	}
}
