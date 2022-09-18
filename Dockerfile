FROM golang:1.18.5-alpine3.16 AS build
ARG TODOCHECK_VERSION=custom_version

RUN apk add --no-cache \
    make
ADD . /usr/src/todocheck
WORKDIR /usr/src/todocheck 

# Build
RUN go build -ldflags "-X main.version=$TODOCHECK_VERSION"
RUN cp todocheck /usr/local/bin/todocheck

FROM alpine:3.16
COPY --from=build /usr/local/bin/todocheck /usr/local/bin/todocheck
ENTRYPOINT ["/usr/local/bin/todocheck"]


