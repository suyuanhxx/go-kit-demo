package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	//stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	//kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8001", "HTTP listen address")
		proxy  = flag.String("proxy", "", "")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	var svc UserService
	svc = userService{}
	svc = proxyingMiddleware(context.Background(), *proxy, logger)(svc)

	getUserByUserNameHandler := httptransport.NewServer(
		makeGetUserByUserNameEndpoint(svc),
		decodeGetUserByUserNameRequest,
		encodeResponse,
	)

	http.Handle("/getUserByUserName", getUserByUserNameHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
