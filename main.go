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

	staticPath := filepath.Join(dir, "static")

	r.Sandbox("/", staticPath)

	crtPathFlag := flag.String("cert", "", "Path to TLS certificate file")
	keyPathFlag := flag.String("key", "", "Path to TLS private key file")
	noRateLimitingFlag := flag.Bool("nolimit", false, "Disable rate-limiting")
	noLoggingFlag := flag.Bool("nolog", false, "Disable connection logging")
	bucketSizeFlag := flag.Int("bs", 2, "Rate-limiting bucket size")
	maxRateFlag := flag.Int("trickle", 2, "Rate-limiting trickle rate")

	flag.Parse()

	srv := houston.NewServer(&r, &houston.ServerConfig{
		CertificatePath: *crtPathFlag,
		KeyPath: *keyPathFlag,

		EnableLimiting: !*noRateLimitingFlag,
		BucketSize: *bucketSizeFlag,
		MaxRate: rate.Limit(*maxRateFlag),

		EnableLog: !*noLoggingFlag,
	})

	fmt.Println("Running...")

	srv.Start()
}
