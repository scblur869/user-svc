FROM golang:latest

WORKDIR /go/src/app
COPY registration-svc .

CMD ["./registration-svc"]