package util

import (
	"crypto/tls"
	"net/http"
	"time"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var httpClient = http.Client{Transport: tr, Timeout: 10 * time.Second}
