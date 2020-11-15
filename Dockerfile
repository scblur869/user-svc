FROM golang:latest

WORKDIR /go/src/app
COPY user-svc .

CMD ["./user-svc"]