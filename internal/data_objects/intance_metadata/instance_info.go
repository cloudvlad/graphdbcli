package intance_metadata

type InstanceMetadata struct {
	Name        string
	Status      string
	Version     string
	CreatedAt   string
	LicenseName string
	Port        string
}

// Compact returns the data in the needed format
func (im *InstanceMetadata) Compact() []string {
	return []string{im.Name, im.Status, im.Version, im.CreatedAt, im.LicenseName, im.Port}
}
