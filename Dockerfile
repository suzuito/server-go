FROM golang:1.17.7

RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin