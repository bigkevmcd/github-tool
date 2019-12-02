FROM golang:latest AS build
COPY . /go/build
WORKDIR /go/build

RUN ls -l && go build ./cmd/github-tool

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /root/
COPY --from=build /go/build/github-tool .
CMD ["./github-tool"]
