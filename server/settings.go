package server

import (
	"flag"
	"net"
	"net/http"
	"os"
)

//Settings ...
type Settings struct {
	AWS  AwsSettings
	HTTP HTTPSettings
}

//NewSettings ...
func NewSettings(args *flag.FlagSet) (s Settings, err error) {
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
}

func getFlagValue(a *flag.FlagSet, f string) (v string) {
	v = a.Lookup(f).Value.String()
	return v
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
	a.Region = getFlagValue(args, "aws-region")
	a.Bucket = getFlagValue(args, "aws-bucket")
	a.Path = getFlagValue(args, "aws-path")
	a.Endpoint = getFlagValue(args, "aws-endpoint")
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

func httpGetSettings(w http.ResponseWriter, s *Settings) {
	SendJSONData(w, s)
}
