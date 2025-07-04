package cmd

import (
	"os"
	"regexp"
	"testing"
)

func TestExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"graphdbcli"}

	err := Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
}

func TestVersionCommand(t *testing.T) {
	oldArgs := os.Args
	Version = "11.0.0"
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"graphdbcli", "--version"}

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	err := Execute()
	w.Close()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	var buf [1024]byte
	n, _ := r.Read(buf[:])
	output := string(buf[:n])
	if output == "" || !testingVersionRegex(output, `graphdbcli version \d+\.\d+\.\d+`) {
		t.Errorf("Expected version output to match regex 'graphdbcli version v\\d+\\.\\d+\\.\\d+', got: %q", output)
	}
}

func testingVersionRegex(output, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(output)
}
