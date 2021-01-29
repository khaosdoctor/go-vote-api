FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/github.com/khaosdoctor/go-vote-api

COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api

FROM scratch
COPY --from=builder /go/bin/api api
EXPOSE 8080
ENTRYPOINT ["./api"]
