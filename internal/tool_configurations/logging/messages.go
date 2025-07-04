package logging

type Message struct {
	Internal string
	External string
}

// ErrorMessages contain messages that are used for logging and user output.
var ErrorMessages = map[int]Message{
	001: {
		Internal: "no license path specified",
		External: "Specify a path to a license file",
	},
	002: {
		Internal: "the selected license file is not a regular file",
		External: "Select a properly formatted license file",
	},
	003: {
		Internal: "The provided license file cannot be opened",
		External: "The provided license file cannot be opened",
	},
	004: {
		Internal: "issues occur with fetching the user's home directory",
		External: "Unable to get user home directory",
	},
	005: {
		Internal: "the licenses directory could not be opened and the license file was not stored",
		External: "Error occurred while storing the license file",
	},
	006: {
		Internal: "The licenses directory could not be opened and the note file for the license was not stored",
		External: "Error occurred while storing the note for the license",
	},
}
