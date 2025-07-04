package authentication

// Token defines the token object, used for marshaling the session token
// for app usage
type Token struct {
	AuthToken string
}

var AuthToken = Token{}
