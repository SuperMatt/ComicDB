package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//NewS3 ...
func NewS3(s Settings) *s3.S3 {
	sess := session.Must(session.NewSession())
	return s3.New(sess)
}
