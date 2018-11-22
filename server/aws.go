package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//NewS3 ...
func (s *Server) NewS3() {
	sess := session.Must(session.NewSession())
	sess.Config.Endpoint = &s.AWS.Endpoint
	sess.Config.Region = &s.AWS.Region
	s3Session := s3.New(sess)
	s.AWS.Session = s3Session
}
