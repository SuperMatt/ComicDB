package server

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//StartServer is the entry point for the server. This takes a *flag.FlagSet for its options
func StartServer(f *flag.FlagSet) {
	serverSettings, err := NewSettings(f)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		httpGetSettings(w, &serverSettings)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
