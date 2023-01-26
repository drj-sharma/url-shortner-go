# BUILD stage
FROM golang:1.18.8-alpine3.15 as builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

COPY . /src

WORKDIR /src
RUN apk add git
RUN apk add make
RUN make build-app

# RUN stage
FROM golang:1.18.8-alpine3.15 AS app

COPY . /app

WORKDIR /app
COPY --from=builder /src/server /app

EXPOSE 8000

CMD [ "/app/server" ]
