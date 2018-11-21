package main

import (
	"flag"
	"os"

	"gitlab.com/SuperMatt/ComicDB/server"
)

func main() {
	appFlags := flag.NewFlagSet("server", flag.ExitOnError)
	appFlags.String("AWS_ACCESS_KEY_ID", "", "AWS_ACCESS_KEY_ID")
	appFlags.String("AWS_SECRET_ACCESS_KEY", "", "AWS_SECRET_ACCESS_KEY")
	appFlags.String("aws-profile", "", "AWS profile name")
	appFlags.String("aws-endpoint", "", "AWS endpoint URI")
	appFlags.String("aws-bucket", "", "AWS bucket name")
	appFlags.String("aws-path", "", "Path to data file")
	appFlags.String("listen", ":8080", "Listen Address")
	appFlags.Parse(os.Args[1:])

	server.StartServer(appFlags)
}
