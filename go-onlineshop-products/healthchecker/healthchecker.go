package healthchecker

import (
	"github.com/heptiolabs/healthcheck"
	"net/http"
)

func AddHealthChecker() {
	health := healthcheck.NewHandler()
	go http.ListenAndServe("0.0.0.0:8086", health)
}
