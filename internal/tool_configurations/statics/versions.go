// Package statics /*
//
// Contains the released versions, their URLs and other
// information about them
package statics

// Version is mapped to a released version of GraphDB.
type Version struct {
	Version     string // Version that is following the Semantic versions standards
	ReleaseDate string // ReleaseDate that the version was published
	Url         string // Url is the distribution file locator
}

// Versions contains all released versions.
var Versions = []Version{
	{
		Version:     "11.0.1",
		ReleaseDate: "07-05-2025",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/11.0.1/graphdb-11.0.1-dist.zip",
	},
	{
		Version:     "11.0.0",
		ReleaseDate: "08-04-2025",
		Url:         "https://download.ontotext.com/owlim/fd0e7994-13c2-11f0-87b0-42843b1b6b38/graphdb-11.0.0-dist.zip",
	},
	{
		Version:     "10.8.7",
		ReleaseDate: "27-05-2025",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.8.7/graphdb-10.8.7-dist.zip",
	},
	{
		Version:     "10.8.6",
		ReleaseDate: "16-05-2025",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.8.6/graphdb-10.8.6-dist.zip",
	},
	{
		Version:     "10.8.5",
		ReleaseDate: "08-04-2025",
		Url:         "https://download.ontotext.com/owlim/6b4bd20e-0efe-11f0-8854-42843b1b6b38/graphdb-10.8.5-dist.zip",
	},
	{
		Version:     "10.8.4",
		ReleaseDate: "25-02-2025",
		Url:         "https://download.ontotext.com/owlim/25801c00-f36d-11ef-8f23-42843b1b6b38/graphdb-10.8.4-dist.zip",
	},
	{
		Version:     "10.8.3",
		ReleaseDate: "28-01-2025",
		Url:         "https://download.ontotext.com/owlim/b757bb32-dd6b-11ef-9c0f-42843b1b6b38/graphdb-10.8.3-dist.zip",
	},
	{
		Version:     "10.8.2",
		ReleaseDate: "13-12-2024",
		Url:         "https://download.ontotext.com/owlim/6122490c-b954-11ef-9c1d-42843b1b6b38/graphdb-10.8.2-dist.zip",
	},
	{
		Version:     "10.8.1",
		ReleaseDate: "02-12-2024",
		Url:         "https://download.ontotext.com/owlim/72817dc2-ae59-11ef-834c-42843b1b6b38/graphdb-10.8.1-dist.zip",
	},
	{
		Version:     "10.8.0",
		ReleaseDate: "11-11-2024",
		Url:         "https://download.ontotext.com/owlim/3efdafa8-a039-11ef-b262-42843b1b6b38/graphdb-10.8.0-dist.zip",
	},
	{
		Version:     "10.7.6",
		ReleaseDate: "15-10-2024",
		Url:         "https://download.ontotext.com/owlim/2b42e0fc-8b01-11ef-afcb-42843b1b6b38/graphdb-10.7.6-dist.zip",
	},
	{
		Version:     "10.7.5",
		ReleaseDate: "Not officially released. Refer to 10.7.6.",
		Url:         "https://download.ontotext.com/owlim/2b42e0fc-8b01-11ef-afcb-42843b1b6b38/graphdb-10.7.6-dist.zip",
	},
	{
		Version:     "10.7.4",
		ReleaseDate: "19-09-2024",
		Url:         "https://download.ontotext.com/owlim/fe8b8a72-7684-11ef-a177-42843b1b6b38/graphdb-10.7.4-dist.zip",
	},
	{
		Version:     "10.7.3",
		ReleaseDate: "18-08-2024",
		Url:         "https://download.ontotext.com/owlim/a364898c-5bdb-11ef-97f6-42843b1b6b38/graphdb-10.7.3-dist.zip",
	},
	{
		Version:     "10.7.2",
		ReleaseDate: "08-08-2024",
		Url:         "https://download.ontotext.com/owlim/0f0b93b6-557d-11ef-8c5b-42843b1b6b38/graphdb-10.7.2-dist.zip",
	},
	{
		Version:     "10.7.1",
		ReleaseDate: "22-07-2024",
		Url:         "https://download.ontotext.com/owlim/ab3e4b6a-45d9-11ef-ab7f-42843b1b6b38/graphdb-10.7.1-dist.zip",
	},
	{
		Version:     "10.7.0",
		ReleaseDate: "04-07-2024",
		Url:         "https://download.ontotext.com/owlim/5e33af66-3a28-11ef-bdf5-42843b1b6b38/graphdb-10.7.0-dist.zip",
	},
	{
		Version:     "10.6.4",
		ReleaseDate: "07-06-2024",
		Url:         "https://download.ontotext.com/owlim/0a7fade4-2406-11ef-9727-42843b1b6b38/graphdb-10.6.4-dist.zip",
	},
	{
		Version:     "10.6.3",
		ReleaseDate: "19-04-2024",
		Url:         "https://download.ontotext.com/owlim/6568897e-fe40-11ee-a47f-42843b1b6b38/graphdb-10.6.3-dist.zip",
	},
	{
		Version:     "10.6.2",
		ReleaseDate: "08-03-2024",
		Url:         "https://download.ontotext.com/owlim/5fc32688-dc98-11ee-ac5d-42843b1b6b38/graphdb-10.6.2-dist.zip",
	},
	{
		Version:     "10.6.1",
		ReleaseDate: "20-02-2024",
		Url:         "https://download.ontotext.com/owlim/e4cbc400-cff2-11ee-aa1c-42843b1b6b38/graphdb-10.6.1-dist.zip",
	},
	{
		Version:     "10.6.0",
		ReleaseDate: "14-02-2024",
		Url:         "https://download.ontotext.com/owlim/84ac68ce-cb3b-11ee-b99d-42843b1b6b38/graphdb-10.6.0-dist.zip",
	},
	{
		Version:     "10.5.1",
		ReleaseDate: "22-01-2024",
		Url:         "https://download.ontotext.com/owlim/eb59afb0-b93d-11ee-9072-42843b1b6b38/graphdb-10.5.1-dist.zip",
	},
	{
		Version:     "10.5.0",
		ReleaseDate: "14-12-2023",
		Url:         "https://download.ontotext.com/owlim/cee4d728-9a9e-11ee-b0a1-42843b1b6b38/graphdb-10.5.0-dist.zip",
	},
	{
		Version:     "10.4.4",
		ReleaseDate: "18-03-2024",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.4.4/graphdb-10.4.4-dist.zip",
	},
	{
		Version:     "10.4.3",
		ReleaseDate: "08-12-2023",
		Url:         "https://download.ontotext.com/owlim/77edd30e-95a6-11ee-8161-42843b1b6b38/graphdb-10.4.3-dist.zip",
	},
	{
		Version:     "10.4.2",
		ReleaseDate: "20-11-2023",
		Url:         "https://download.ontotext.com/owlim/8f46c1be-8796-11ee-a50e-42843b1b6b38/graphdb-10.4.2-dist.zip",
	},
	{
		Version:     "10.4.1",
		ReleaseDate: "01-11-2023",
		Url:         "https://download.ontotext.com/owlim/ed77a766-7806-11ee-afce-42843b1b6b38/graphdb-10.4.1-dist.zip",
	},
	{
		Version:     "10.4.0",
		ReleaseDate: "17-10-2023",
		Url:         "https://download.ontotext.com/owlim/4827d51e-6c36-11ee-9511-42843b1b6b38/graphdb-10.4.0-dist.zip",
	},
	{
		Version:     "10.3.3",
		ReleaseDate: "14-09-2023",
		Url:         "https://download.ontotext.com/owlim/12f42f78-5246-11ee-9e53-42843b1b6b38/graphdb-10.3.3-dist.zip",
	},
	{
		Version:     "10.3.2",
		ReleaseDate: "01-09-2023",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.3.2/graphdb-10.3.2-dist.zip",
	},
	{
		Version:     "10.3.1",
		ReleaseDate: "01-08-2023",
		Url:         "https://download.ontotext.com/owlim/13956f1e-3766-11ee-910e-42843b1b6b38/graphdb-10.3.1-dist.zip",
	},
	{
		Version:     "10.3.0",
		ReleaseDate: "17-07-2023",
		Url:         "https://download.ontotext.com/owlim/06e1cb10-2496-11ee-b9aa-42843b1b6b38/graphdb-10.3.0-dist.zip",
	},
	{
		Version:     "10.2.5",
		ReleaseDate: "01-09-2023",
		Url:         "https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.2.5/graphdb-10.2.5-dist.zip",
	},
	{
		Version:     "10.2.4",
		ReleaseDate: "07-08-2023",
		Url:         "https://download.ontotext.com/owlim/d331a62c-32ce-11ee-a90c-42843b1b6b38/graphdb-10.2.4-dist.zip",
	},
	{
		Version:     "10.2.3",
		ReleaseDate: "12-07-2023",
		Url:         "https://download.ontotext.com/owlim/c239472e-20c9-11ee-9b9d-42843b1b6b38/graphdb-10.2.3-dist.zip",
	},
	{
		Version:     "10.2.2",
		ReleaseDate: "07-06-2023",
		Url:         "https://download.ontotext.com/owlim/16090cd4-0506-11ee-91b1-42843b1b6b38/graphdb-10.2.2-dist.zip",
	},
	{
		Version:     "10.2.1",
		ReleaseDate: "25-04-2023",
		Url:         "http://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.2.1/graphdb-10.2.1-dist.zip",
	},
	{
		Version:     "10.2.0",
		ReleaseDate: "28-02-2023",
		Url:         "http://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/10.2.0/graphdb-10.2.0-dist.zip",
	},
}

/*
GetVersionIndex checks if the selected version is distributed by the CLI.
If found - returns the index of the elements corresponding to the version.
If not found - returns -1
*/
func GetVersionIndex(selectedVersion string) int {
	for i, version := range Versions {
		if selectedVersion == version.Version {
			return i
		}
	}

	return -1
}
