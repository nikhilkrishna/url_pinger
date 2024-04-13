package tests

import (
	"bytes"
	"os"
	"sync"
	"testing"
	"url_pinger/pkg/logger"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

func TestStdoutLogger(t *testing.T) {
	var wg sync.WaitGroup
	logger := logger.NewStdoutLogger()

	output := captureOutput(func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Log("test message")
		}()
		wg.Wait()
		logger.Close()
	})

	if expected, got := "test message\n", output; got != expected {
		t.Errorf("Expected log output %q, got %q", expected, got)
	}

}
