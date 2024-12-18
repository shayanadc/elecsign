package main

import (
	"elecsign/cmd"
	"elecsign/internal/display"
	"testing"
)

// Setup helper function
func setupDisplay() display.Display {
	renderer := display.NewConsoleRenderer()
	return display.NewConsoleDisplay(renderer)
}

// Benchmark adding pixel views
func BenchmarkAddPixelView(b *testing.B) {
	d := setupDisplay()
	handler := cmd.NewCommandHandler(d)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = handler.HandleAdd([]string{"pixel", "A0", "B1", "C2"})
	}
}

// Benchmark adding character views
func BenchmarkAddCharacterView(b *testing.B) {
	d := setupDisplay()
	handler := cmd.NewCommandHandler(d)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = handler.HandleAdd([]string{"character", "ABC123"})
	}
}

// Benchmark display rendering
func BenchmarkDisplayShow(b *testing.B) {
	d := setupDisplay()
	handler := cmd.NewCommandHandler(d)

	// Setup some views
	_ = handler.HandleAdd([]string{"pixel", "A0", "B1", "C2"})
	_ = handler.HandleAdd([]string{"character", "ABC"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.HandleShow()
	}
}

// Benchmark complete workflow
func BenchmarkCompleteWorkflow(b *testing.B) {
	d := setupDisplay()
	handler := cmd.NewCommandHandler(d)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = handler.HandleAdd([]string{"pixel", "A0", "B1"})
		_ = handler.HandleAdd([]string{"character", "ABC"})
		handler.HandleShow()
		handler.HandleClear()
	}
}

// Benchmark memory usage
func BenchmarkMemoryUsage(b *testing.B) {
	d := setupDisplay()
	handler := cmd.NewCommandHandler(d)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = handler.HandleAdd([]string{"pixel", "A0", "B1"})
		handler.HandleShow()
		handler.HandleClear()
	}
}
