// Package statics /*
//
// Contains the released versions, their URLs and other
// information about them
package statics

// Version is mapped to a released version of GraphDB.
type Version struct {
	Version     string // Version that is following the Semantic versions standards
	ReleaseDate string // ReleaseDate that the version was published
	Sha256sum   string // Sha256sum is the hash of the file
}

// Versions contains all released versions.
var Versions = []Version{
	{
		Version:     "11.1.2",
		ReleaseDate: "15-10-2025",
		Sha256sum:   "36a92fff05328051c39f1b5cdebc5762c75ced36f0a29ea271fc0fdfa743c0c4",
	},
	{
		Version:     "11.1.1",
		ReleaseDate: "17-09-2025",
		Sha256sum:   "dae1e7118dadd2207ce42f178d24063b9f4fd59ac1891ea2cf1b7e33993fbd7b",
	},
	{
		Version:     "11.1.0",
		ReleaseDate: "21-08-2025",
		Sha256sum:   "29fbf76624b3bdfa9562c61987fe54110e4af53d628a029c8c543931cd7c2e5a",
	},
	{
		Version:     "11.0.2",
		ReleaseDate: "07-07-2025",
		Sha256sum:   "7abac3ad5b47f4243fbda871f4e4695d0c978c5d4808beade81bf06cfcf6f30a",
	},
	{
		Version:     "11.0.1",
		ReleaseDate: "07-05-2025",
		Sha256sum:   "294ecbff5339a112a569a9ceb6efb0bd2e10755d067a0b9e22136fff3eaaf2f4",
	},
	{
		Version:     "11.0.0",
		ReleaseDate: "08-04-2025",
		Sha256sum:   "408106abc4d1f75db6c854940b671d3caf760624812bd8a045c1d33748f5090d",
	},
	{
		Version:     "10.8.11",
		ReleaseDate: "24-09-2025",
		Sha256sum:   "bc487769d220024ef8f4fc8c54dec00704f0ee181ceaaa683f6902157bc07d5b",
	},
	{
		Version:     "10.8.10",
		ReleaseDate: "21-08-2025",
		Sha256sum:   "02df04cf134855cae9e4b927e7d2d5b6df8760ea15f7cb61b9122409496dd518",
	},
	{
		Version:     "10.8.9",
		ReleaseDate: "14-07-2025",
		Sha256sum:   "ef36698f1939a4ed0fdc8a74d30240ca946c5996bad8c59ec7c4eba1df635423",
	},
	{
		Version:     "10.8.8",
		ReleaseDate: "05-06-2025",
		Sha256sum:   "b1bc91fbe01128569d9e680c0eec62d9a67baf704edf1554f578a1debd6ae2b7",
	},
	{
		Version:     "10.8.7",
		ReleaseDate: "27-05-2025",
		Sha256sum:   "35b46bbf95fc15c97730739dfc2a2932a2c10cb39866982449dc9245d2a4e2b6",
	},
	{
		Version:     "10.8.6",
		ReleaseDate: "16-05-2025",
		Sha256sum:   "8968a6d4c5b9f2b709c54cd167341b61a65022de1a356b922ac8b77434d24423",
	},
	{
		Version:     "10.8.5",
		ReleaseDate: "08-04-2025",
		Sha256sum:   "7c3e0a2838f944fa887ddebaf3d7cac6801564d2701dea10144a97a51e28e21f",
	},
	{
		Version:     "10.8.4",
		ReleaseDate: "25-02-2025",
		Sha256sum:   "9a14dee2624e6fdce7c6ef797a89583a5f7a0c4447fc72603822a69f676c4f5d",
	},
	{
		Version:     "10.8.3",
		ReleaseDate: "28-01-2025",
		Sha256sum:   "1f4f1549d16a60fa5f90184b791070ce43a20255c1b29d194e8c3795fb2773ab",
	},
	{
		Version:     "10.8.2",
		ReleaseDate: "13-12-2024",
		Sha256sum:   "460cdbc897c7001a8302016e4a27bad451dc2619716ed8208078e9a8db66a84e",
	},
	{
		Version:     "10.8.1",
		ReleaseDate: "02-12-2024",
		Sha256sum:   "3c03fc01312077f41214321e0e708d7c63f9c380fcbbf9b99c27923fa1a0c040",
	},
	{
		Version:     "10.8.0",
		ReleaseDate: "11-11-2024",
		Sha256sum:   "0753363dfb80d7a12432f46620b0d2e5d36eb44ade1f08ed4bef4cb4cf11a1b5",
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
