# Build Stage
FROM golang:latest as build

ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/nw-union/hidel-wiki
COPY . .
RUN env CGO_ENABLED=0 go install

# Release Stage
FROM gcr.io/distroless/base

COPY --from=build /go/bin/hidel-wiki /hidel-wiki
ENV PORT=8080
EXPOSE 8080

CMD ["/hidel-wiki"]
