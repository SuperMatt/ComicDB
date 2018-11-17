package server

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

//Settings ...
type Settings struct {
	AWS  AwsSettings
	HTTP HTTPSettings
}

//NewSettings ...
func NewSettings(args *flag.FlagSet) (s Settings, err error) {
	s.AWS, err = NewAwsSettings(args)
	if err != nil {
		return s, err
	}
	s.HTTP, err = NewHTTPSettings(args)
	return s, err
}

//AwsSettings ...
type AwsSettings struct {
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Path            string
}

func getFlagValue(a *flag.FlagSet, f string) (v string, err error) {
	v = a.Lookup(f).Value.String()
	if v == "" {
		return v, fmt.Errorf("Blank value found for variable %s", f)
	}

	return v, nil
}

//NewAwsSettings ...
func NewAwsSettings(args *flag.FlagSet) (a AwsSettings, err error) {
	a.AccessKeyID, err = getFlagValue(args, "AWS_ACCESS_KEY_ID")
	if err != nil {
		return a, err
	}
	a.SecretAccessKey, err = getFlagValue(args, "AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return a, err
	}
	a.Bucket, err = getFlagValue(args, "aws-bucket")
	if err != nil {
		return a, err
	}
	a.Path, err = getFlagValue(args, "aws-path")
	if err != nil {
		return a, err
	}

	return a, nil
}

//HTTPSettings ...
type HTTPSettings struct {
	BindIP      string
	BindPort    string
	BindAddress string
}

//NewHTTPSettings ...
func NewHTTPSettings(args *flag.FlagSet) (h HTTPSettings, err error) {
	l, err := getFlagValue(args, "listen")
	if err != nil {
		return h, err
	}

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
