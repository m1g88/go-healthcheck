FROM golang:1.16

WORKDIR /go/src/go-healthcheck
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-healthcheck"]

# ENTRYPOINT [ "go-healthcheck" ]