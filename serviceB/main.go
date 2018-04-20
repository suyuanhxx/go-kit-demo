package main

import (
	"flag"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	var (
		listen = flag.String("listen", ":8001", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

}
