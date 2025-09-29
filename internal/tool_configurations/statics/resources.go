package statics

// Resource is an object that describes things that are compatible with GraphDB
type Resource struct {
	Name    string // Name that is identifying the resource
	Desc    string
	IpfsCID string // IpfsCID is the distribution file locator
}

// Resources contains public available, pre-uploaded, data that may be used with graphDB
var Resources = []Resource{
	{
		Name:    "starwars-data.ttl",
		Desc:    "Fetches a starwars dataset",
		IpfsCID: "QmUotMH5KxqncztDjfXAtgHKT4gBtLouVDdvEEDxE8SndW",
	},
	{
		Name:    "sss",
		Desc:    "Fetches a sss dataset",
		IpfsCID: "QmSv62HH51wCyoSdpBeG5Dq3nkQAxvCgcHKPSooEy7qqzg",
	},
}
