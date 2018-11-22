package server

import (
	"net/http"
)

func httpShowAll(w http.ResponseWriter, s *Server) {
	db, err := s.ReadDBFile()
	if err != nil {
		SendInternalServerError(w, err)
		return
	}

	SendJSONData(w, db)
}
