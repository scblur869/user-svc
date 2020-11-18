FROM golang:alpine

WORKDIR /go/src/app
COPY user-svc .

CMD ["./user-svc"]