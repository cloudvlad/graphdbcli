// Package authentication defines the authentication objects.
package authentication

// BasicAuth defines the structure used for marshaling credentials
// in a structured manner.
type BasicAuth struct {
	Username string
	Password string
}

var BasicCredentials = BasicAuth{}
