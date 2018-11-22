package server

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//StartServer is the entry point for the server. This takes a *flag.FlagSet for its options
func StartServer(f *flag.FlagSet) {
	s, err := NewServer(f)
	if err != nil {
		log.Fatal(err)
	}

	s.NewS3()

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		httpGetSettings(w, &s)
	})

	http.HandleFunc("/showall", func(w http.ResponseWriter, r *http.Request) {
		httpShowAll(w, &s)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
