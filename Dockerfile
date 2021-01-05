### BUILD FRONTEND
FROM node:12-alpine as frontend

WORKDIR /client

COPY client/package.json .

RUN yarn install

COPY client/ .

RUN yarn build

### BUILD BACKEND
FROM golang:1.14 as builder

WORKDIR /server

COPY server/ .

RUN mkdir -p /client

COPY --from=frontend /client/build /client/build

WORKDIR /

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY main.go main.go

RUN go get github.com/GeertJohan/go.rice github.com/GeertJohan/go.rice/rice && \
    rice embed-go && \
    CGO_ENABLED=0 GOOS=linux go build -o app .

### READY FOR RUN
FROM alpine:3.6

RUN apk add --no-cache ca-certificates

COPY --from=builder /app /go/bin/app

EXPOSE 8080

CMD ["/go/bin/app"]