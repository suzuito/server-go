package usecase

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/suzuito/blog1-server-go/entity"
)

type ReverseProxy interface {
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}

type ReverseProxyFactory interface {
	NewReverseProxy(transport http.RoundTripper) ReverseProxy
}

type ReverseProxyFactoryImpl struct {
	Target *url.URL
}

func (r *ReverseProxyFactoryImpl) NewReverseProxy(transport http.RoundTripper) ReverseProxy {
	pxy := httputil.NewSingleHostReverseProxy(r.Target)
	pxy.Transport = transport
	return pxy
}

type RoundTripperImpl struct {
	entry *entity.LogEntry
}

func (r *RoundTripperImpl) RoundTrip(req *http.Request) (*http.Response, error) {
	r.entry.TargetURI = req.URL.String()
	r.entry.TargetMethod = req.Method
	r.entry.TargetStartedAt = time.Now()
	defer func() {
		r.entry.TargetResponsedAt = time.Now()
	}()
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	r.entry.TargetStatusCode = res.StatusCode
	return res, nil
}

func NewRoundTripperImpl(le *entity.LogEntry) *RoundTripperImpl {
	return &RoundTripperImpl{
		entry: le,
	}
}
