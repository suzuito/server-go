package usecase

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ReverseProxy interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type RoundTripperImpl struct {
}

func (r *RoundTripperImpl) RoundTrip(req *http.Request) (*http.Response, error) {
	le := GetContextLogEntry(req)
	if le != nil {
		le.TargetURI = req.URL.String()
		le.TargetMethod = req.Method
		le.TargetStartedAt = time.Now()
	}
	defer func() {
		if le != nil {
			le.TargetResponsedAt = time.Now()
		}
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
	if le != nil {
		le.TargetStatusCode = res.StatusCode
	}
	return res, nil
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	r := httputil.NewSingleHostReverseProxy(target)
	r.Transport = &RoundTripperImpl{}
	return r
}
