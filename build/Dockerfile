FROM golang:1.12 as schd-extender-builder

COPY / /go/src/github.com/k8s-scheduler-extender

WORKDIR /go/src/github.com/k8s-scheduler-extender
RUN CGO_ENABLED=0 GO111MODULE=off GOOS=linux GOARCH=amd64 go build -i -ldflags '-w -s' -o schd-extender cmd/main.go

FROM alpine:3.9
RUN apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk add --update ca-certificates && update-ca-certificates
COPY --from=schd-extender-builder /go/src/github.com/k8s-scheduler-extender/schd-extender /usr/local/bin/
CMD ["sh"]