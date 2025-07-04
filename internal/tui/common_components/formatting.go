package common_components

import (
	"graphdbcli/internal/tool_configurations/statics"
	"strings"

	"github.com/mattn/go-runewidth"
)

func PadExamples(example string) string {
	lines := strings.Split(example, "\n")
	// Remove starting blank lines
	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	lines = lines[start:]
	for i, line := range lines {
		lines[i] = strings.Repeat(" ", 2) + line
	}
	return strings.Join(lines, "\n")
}

// PadStatusIndicator pads the status with spaces after it
func PadStatusIndicator(status string, additional int) string {
	w := runewidth.StringWidth(status)
	//println(status, w)
	if w == 1 {
		w = 2
	}
	if w < statics.TUIStatusIndicatorWidth {
		return status + strings.Repeat("+", statics.TUIStatusIndicatorWidth-w)
	}

	status += strings.Repeat("~", additional)
	return status
}
