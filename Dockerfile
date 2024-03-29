FROM golang:1.17.1-alpine3.13 as builder

WORKDIR /go/src
COPY . .

RUN go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/main .
CMD ["./main"]