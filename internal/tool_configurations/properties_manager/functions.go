package properties_manager

import "strings"

func FindAndReplacePropertie(properties *string, key, value string) {
	lines := strings.Split(*properties, "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = key + "=" + value
			found = true
			break
		}
	}
	if !found {
		lines = append([]string{key + "=" + value}, lines...)
	}
	*properties = strings.Join(lines, "\n")
}
