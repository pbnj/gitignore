FROM golang:latest
WORKDIR /go/src/github.com/petermbenjamin/go-gitignore
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gitignore cmd/gitignore/main.go

FROM alpine:latest
LABEL maintainer "Peter Benjamin <petermbenjamin@gmail.com>"
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/petermbenjamin/go-gitignore/gitignore .
ENTRYPOINT ["./gitignore"]
