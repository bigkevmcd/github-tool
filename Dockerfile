FROM golang:latest AS build
COPY . /go/build
WORKDIR /go/build

RUN go build ./cmd/github-tool

FROM gcr.io/distroless/base-debian11
WORKDIR /root/
COPY --from=build /go/build/github-tool /usr/local/bin
ENTRYPOINT ["/usr/local/bin/github-tool"]
