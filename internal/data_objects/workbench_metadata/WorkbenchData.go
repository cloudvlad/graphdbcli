package workbench_metadata

type Data struct {
	Name          string
	Status        string
	WorkbenchPort string
	GraphDBPort   string
	GraphDBHost   string
	CreatedAt     string
}

// Compact returns the data in the needed format
func (wbmd *Data) Compact() []string {
	return []string{wbmd.Name, wbmd.Status, wbmd.WorkbenchPort, wbmd.GraphDBPort, wbmd.GraphDBHost, wbmd.CreatedAt}
}
