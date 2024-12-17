package display

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"elecsign/internal/transformer"
	"elecsign/internal/view"
)

func TestConsoleRenderer_Render(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "pattern A",
			input: "A",
			expected: []string{
				"  **" + strings.Repeat(" ", 32),
				" *  *" + strings.Repeat(" ", 31),
				"*****" + strings.Repeat(" ", 31),
				"*   *" + strings.Repeat(" ", 31),
				"*   *" + strings.Repeat(" ", 31),
				"*   *" + strings.Repeat(" ", 31),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			renderer := NewConsoleRenderer()
			display := NewConsoleDisplay(renderer)
			view := view.NewView()

			// Create transformer and transform input
			transformer, err := transformer.NewTransformer("character")
			if err != nil {
				t.Fatalf("failed to create transformer: %v", err)
			}
			coordinates := transformer.Transform(tt.input, 0)
			view.TurnOn(coordinates)
			display.AddView(view)

			// Capture output
			var buf bytes.Buffer
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() { os.Stdout = old }() // Restore stdout after the test

			// Show display
			display.Show()

			// Close the writer and copy the output
			w.Close()
			_, _ = io.Copy(&buf, r)

			// Compare output
			output := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
			compareOutput(t, output, tt.expected)
		})
	}
}

// compareOutput checks if the actual output matches the expected output
func compareOutput(t *testing.T, output, expected []string) {
	if len(output) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(output))
		return
	}

	for i, expectedLine := range expected {
		if output[i] != expectedLine {
			t.Errorf("Line %d:\nexpected: %q\ngot: %q", i, expectedLine, output[i])
		}
	}
}
