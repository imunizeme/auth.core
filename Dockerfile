FROM golang:1.8-alpine

RUN mkdir -p /go/src/github.com/imunizeme/auth.core
COPY  ./ /go/src/github.com/imunizeme/auth.core
WORKDIR /go/src/github.com/imunizeme/auth.core
CMD ["go", "run", "main.go"]
