# setup base image
FROM golang:1.22-alpine as base


# build dependency
FROM base AS dependency

WORKDIR /build

RUN apk add --update make

COPY go.mod .
COPY go.sum .

RUN go mod download

# build image
FROM dependency as builder


COPY . .

RUN make build

FROM golang:1.22-alpine

COPY --from=builder /build/build/basic-web /usr/local/bin/basic-web
COPY --from=builder /build/.env /usr/local/bin/.env

RUN chmod +x /usr/local/bin/basic-web

WORKDIR /usr/local/bin

ENTRYPOINT ["basic-web"]

CMD [ "web" ]
