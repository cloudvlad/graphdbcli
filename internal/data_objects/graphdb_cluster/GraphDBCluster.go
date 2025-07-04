// Package graphdb_cluster defines the cluster object.
package graphdb_cluster

var PropertyOverridesFilepath string

var Instance = GraphDBInstance{}

// GraphDBInstance defines all relative data for an instance management.
type GraphDBInstance struct {
	Name                  string
	Version               string
	StoredLicenseFilename string
	IsActive              bool
	PropertyOverrides     string
	Port                  string
}
