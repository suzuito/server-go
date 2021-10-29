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
	// GCSのpublicリポジトリは
	// HostヘッダがURLと違っている場合、以下の403エラーが返ってくる。
	//  <?xml version='1.0' encoding='UTF-8'?>
	//  <Error>
	//    <Code>UserProjectAccountProblem</Code>
	//    <Message>User project billing account not in good standing.</Message>
	//    <Details>The billing account for the owning project is disabled in state closed</Details>
	//  </Error>
	req.Host = req.URL.Host
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
