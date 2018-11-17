package server

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//StartServer is the entry point for the server. This takes a *flag.FlagSet for its options
func StartServer(*flag.FlagSet) {
	http.HandleFunc("/settings", httpGetSettings)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
