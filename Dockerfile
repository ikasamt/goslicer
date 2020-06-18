FROM golang:1.14.4

WORKDIR /go/src



RUN export GO111MODULE=on && \
   go get -u github.com/ikasamt/zapp-jam && \
   go get -u github.com/cheekybits/genny
