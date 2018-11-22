package server

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
)

//Server ...
type Server struct {
	AWS  AwsSettings
	HTTP HTTPSettings
}

//NewServer ...
func NewServer(args *flag.FlagSet) (s Server, err error) {
	s.AWS = NewAwsSettings(args)
	s.HTTP, err = NewHTTPSettings(args)
	return s, err
}

//AwsSettings ...
type AwsSettings struct {
	Endpoint        string
	Profile         string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Path            string
	Region          string
	Session         *s3.S3
}

func getFlagValue(a *flag.FlagSet, f string) string {
	v := a.Lookup(f)
	if v == nil {
		return ""
	}

	return v.Value.String()
}

func setEnvIfNotSet(k, v string) {
	if os.Getenv(k) == "" {
		os.Setenv(k, v)
	}
}

//NewAwsSettings ...
func NewAwsSettings(args *flag.FlagSet) (a AwsSettings) {
	a.AccessKeyID = getFlagValue(args, "AWS_ACCESS_KEY_ID")
	a.SecretAccessKey = getFlagValue(args, "AWS_SECRET_ACCESS_KEY")
	a.Profile = getFlagValue(args, "aws-profile")
	r := getFlagValue(args, "aws-region")
	a.Region = r
	if r == "" {
		a.Region = "us-west-1"
	}
	a.Bucket = getFlagValue(args, "aws-bucket")
	a.Path = getFlagValue(args, "aws-path")
	a.Endpoint = getFlagValue(args, "aws-endpoint")

	a.SetAWSEnvironmentVariables()

	return a
}

//SetAWSEnvironmentVariables ...
func (a AwsSettings) SetAWSEnvironmentVariables() {
	setEnvIfNotSet("AWS_ACCESS_KEY_ID", a.AccessKeyID)
	setEnvIfNotSet("AWS_SECRET_ACCESS_KEY", a.SecretAccessKey)
	setEnvIfNotSet("AWS_PROFILE", a.Profile)
	setEnvIfNotSet("AWS_REGION", a.Region)
}

//HTTPSettings ...
type HTTPSettings struct {
	BindIP      string
	BindPort    string
	BindAddress string
}

//NewHTTPSettings ...
func NewHTTPSettings(args *flag.FlagSet) (h HTTPSettings, err error) {
	l := getFlagValue(args, "listen")

	iface, port, err := net.SplitHostPort(l)
	if err != nil {
		return h, err
	}

	h.BindIP = iface
	h.BindPort = port
	h.BindAddress = l

	return h, nil
}

func httpGetSettings(w http.ResponseWriter, s *Server) {
	SendJSONData(w, s)
}
