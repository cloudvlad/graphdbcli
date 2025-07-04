// Package statics /*
//
// Contains propertie that are related to CLI global properties
// and GraphDB related properties.
package statics

// DefaultConfPropertiesPath specifies the default path to the properties file
// in a GraphDB instance
var DefaultConfPropertiesPath = []string{"conf", "graphdb.properties"}
var DefaultLicensePath = []string{"conf", "graphdb.license"}
var DefaultExecutablePath = []string{"bin", "graphdb"}

// IsTuiDisabled sets the default value for usage of TUI.
var IsTuiDisabled bool = false

// TUIStatusIndicatorWidth specified width of the symbol (emoji  or create)
// Used for correctly padding the initial part, as some symbols, more specifically emojis,
// are rendered more than 1 symbol.
var TUIStatusIndicatorWidth = 2

var NotTUIStatusIndicatorAdditionalPadding = 0
