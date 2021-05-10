FROM golang:1.12.6-stretch

# Setup
WORKDIR /build
COPY go.* ./
COPY *.go ./

# Install older version of GopherJS
RUN go get github.com/gopherjs/gopherjs@1.12.3+go1.12

# Special sauce part 1
RUN go mod vendor

# Special sauce part 2 - copy to GOPATH
RUN mkdir -p /go/src/hcl2-parser
RUN cp -r * /go/src/hcl2-parser
WORKDIR /go/src/hcl2-parser

# Finally we can build
RUN gopherjs build parser.go -o index.js -m

COPY build/copy-out.sh .
ENTRYPOINT ["/go/src/hcl2-parser/copy-out.sh"]