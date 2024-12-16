package internal

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
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
			renderer := &ConsoleRenderer{}
			display := NewConsoleDisplay(renderer)
			view := NewView(36, 6)

			// Create transformer and transform input
			transformer, _ := NewTransformer("character")
			coordinates := transformer.Transform(tt.input)
			view.TurnOn(coordinates)
			display.AddView(view)

			// Capture output
			var buf bytes.Buffer
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Show display
			display.Show()

			// Restore stdout
			w.Close()
			os.Stdout = old
			io.Copy(&buf, r)

			// Compare output
			output := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")

			// Debug output
			t.Logf("Got output:\n%s", buf.String())

			if len(output) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(output))
				return
			}

			for i, expectedLine := range tt.expected {
				if output[i] != expectedLine {
					t.Errorf("Line %d:\nexpected: %q\ngot:      %q", i, expectedLine, output[i])
				}
			}
		})
	}
}