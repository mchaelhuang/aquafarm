FROM golang:1.20.3-alpine as Build
COPY . .
RUN GOPATH= go build -o /bin/aquafarm-rest ./cmd/aquafarm-rest
ENTRYPOINT [ "/bin/aquafarm-rest" ]