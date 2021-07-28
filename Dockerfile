# Prepare
FROM golang:1.14-alpine as baseimg

RUN apk --no-cache upgrade && apk --no-cache add git make build-base

ARG ENV
ARG BUILD_TAGS="-tags musl"

# First only download the dependencies, so the dependencies can be cahced before we copy the code
COPY go.mod go.sum /app/

WORKDIR /app/

# Installation step. Without -s the command ends up exposing the GitHub token
RUN go mod download

COPY . ./

# Build step. Without -s the command ends up exposing the GitHub token
RUN rm -rf ./bin

RUN go build -o bin/scrapper cmd/main.go

# Use alpine image for optimazation
FROM alpine

COPY --from=baseimg /app/bin/scrapper /opt/

RUN apk --no-cache upgrade && apk --no-cache add git build-base make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /opt/

CMD ["./scrapper"]
