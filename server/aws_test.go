package server

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

func TestS3Session(t *testing.T) {
	var s Settings

	s.AWS.Profile = "digitalocean"
	s.AWS.Endpoint = "https://ams3.digitaloceanspaces.com/"
	s.AWS.Region = "us-west-1"
	s.AWS.Bucket = "matty-comicdb"
	s.AWS.Path = "db.json"

	s.AWS.SetAWSEnvironmentVariables()

	sess := NewS3(s)

	var l s3.GetObjectInput
	l.Bucket = &s.AWS.Bucket
	l.Key = &s.AWS.Path
	out, err := sess.GetObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
