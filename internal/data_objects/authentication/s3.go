package authentication

type S3 struct {
	AccessKeyID     string
	SecretAccessKey string
}

var S3Authentication = S3{}
