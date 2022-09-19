FROM golang:1.18.5-alpine3.16 AS build
ARG TODOCHECK_VERSION=custom_version

ADD . /usr/src/todocheck
WORKDIR /usr/src/todocheck 

# Build
RUN go build -ldflags "-X main.version=$TODOCHECK_VERSION"

FROM alpine:3.16
COPY --from=build /usr/src/todocheck/todocheck /usr/local/bin/todocheck
ENTRYPOINT ["/usr/local/bin/todocheck"]
