package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	log := AppLogger(os.Stdout)
	// Redirect output to buffer for testing
	log.Logger.SetOutput(&buf)

	t.Run("Debug", func(t *testing.T) {
		buf.Reset()
		log.Debug("debug message")
		output := buf.String()
		if !strings.Contains(output, "DEBUG") {
			t.Errorf("Expected DEBUG in output, got %s", output)
		}
		if !strings.Contains(output, "debug message") {
			t.Errorf("Expected 'debug message' in output, got %s", output)
		}
	})

	t.Run("Info", func(t *testing.T) {
		buf.Reset()
		log.Info("info message")
		output := buf.String()
		if !strings.Contains(output, "INFO") {
			t.Errorf("Expected INFO in output, got %s", output)
		}
		if !strings.Contains(output, "info message") {
			t.Errorf("Expected 'info message' in output, got %s", output)
		}
	})

	t.Run("Warning", func(t *testing.T) {
		buf.Reset()
		log.Warning("warning message")
		output := buf.String()
		if !strings.Contains(output, "WARNING") {
			t.Errorf("Expected WARNING in output, got %s", output)
		}
		if !strings.Contains(output, "warning message") {
			t.Errorf("Expected 'warning message' in output, got %s", output)
		}
	})

	t.Run("Error", func(t *testing.T) {
		buf.Reset()
		log.Error("error message")
		output := buf.String()
		if !strings.Contains(output, "ERROR") {
			t.Errorf("Expected ERROR in output, got %s", output)
		}
		if !strings.Contains(output, "error message") {
			t.Errorf("Expected 'error message' in output, got %s", output)
		}
	})
}

func TestSingleton(t *testing.T) {
	log := Logging()
	if log == nil {
		t.Fatal("Expected singleton logger to be initialized")
	}
}
