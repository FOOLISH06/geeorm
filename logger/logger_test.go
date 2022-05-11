package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("Failed to set log level as ErrorLevel")
	}

	SetLevel(InfoLevel)
	if infoLog.Writer() != os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("Failed to set log level as InfoLevel")
	}

	SetLevel(Disable)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("Failed to set log level as Disable")
	}
}
