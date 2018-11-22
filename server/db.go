package server

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/s3"
	"gitlab.com/SuperMatt/ComicDB/comicdb"
)

//ReadDBFile ...
func (s *Server) ReadDBFile() (*comicdb.ComicDB, error) {
	var l s3.GetObjectInput
	l.Bucket = &s.AWS.Bucket
	l.Key = &s.AWS.Path

	obj, err := s.AWS.Session.GetObject(&l)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	var db comicdb.ComicDB

	err = json.Unmarshal(b, &db)
	if err != nil {
		return nil, err
	}

	return &db, nil

}
