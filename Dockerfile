FROM golang:1.16.3-alpine3.13 AS GO_BUILD
COPY server /server
WORKDIR /server
RUN GOOS=linux go build -o /go/bin/server

FROM alpine:3.13.5
COPY --from=GO_BUILD /go/bin/server ./
ENTRYPOINT ./server --port 8080
