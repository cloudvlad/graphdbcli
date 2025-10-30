package backup_conf

// BackupConfigurations defines all relative data
// for backup creation
type BackupConfigurations struct {
	Repositories     []string
	BackupSystemData bool
	GraphDBLocation  string
	BackupName       string
}

// GetRepositories Splits the repositories names (comma separated) and returns them as an array of strings.
func (bc BackupConfigurations) GetRepositories() []string {
	return bc.Repositories
}

var Configurations = BackupConfigurations{}
