FROM golang:1.19-buster as build

WORKDIR /go/src/helloworld
ADD . /go/src/helloworld

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o /go/bin/helloworld ./main.go

# Runtime Container
FROM gcr.io/distroless/static-debian11:nonroot

USER nonroot

COPY --from=build /go/bin/helloworld /

CMD ["/helloworld"]
