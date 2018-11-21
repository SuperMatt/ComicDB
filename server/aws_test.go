package server

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

func TestS3Session(t *testing.T) {
	var s Settings

	s.AWS.Profile = "digitalocean"
	s.AWS.Endpoint = "https://ams3.digitaloceanspaces.com/"
	s.AWS.Region = "us-west-1"

	sess := NewS3(s)

	var l s3.ListBucketsInput
	out, err := sess.ListBuckets(&l)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
