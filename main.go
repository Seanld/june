package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~seanld/houston"
	"golang.org/x/time/rate"
)

func main() {
	r := houston.BlankRouter()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("ERROR: Failed to get working directory.")
	}

	srvPath := flag.String("path", "static", "Path to directory, from which static content is served")
	crtPathFlag := flag.String("cert", "", "Path to TLS certificate file")
	keyPathFlag := flag.String("key", "", "Path to TLS private key file")
	noRateLimitingFlag := flag.Bool("nolimit", false, "Disable rate-limiting")
	noLoggingFlag := flag.Bool("nolog", false, "Disable connection logging")
	bucketSizeFlag := flag.Int("bs", 2, "Rate-limiting bucket size")
	maxRateFlag := flag.Int("trickle", 2, "Rate-limiting trickle rate")

	flag.Parse()

	staticPath := filepath.Join(dir, *srvPath)
	r.Sandbox("/", staticPath)

	srv := houston.NewServer(&r, &houston.ServerConfig{
		CertificatePath: *crtPathFlag,
		KeyPath: *keyPathFlag,

		EnableLimiting: !*noRateLimitingFlag,
		BucketSize: *bucketSizeFlag,
		MaxRate: rate.Limit(*maxRateFlag),

		EnableLog: !*noLoggingFlag,
		LogFilePath: "june.log",
	})

	fmt.Println("Running...")

	srv.Start()
}
