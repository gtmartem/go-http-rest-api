FROM golang:1.13.7-alpine3.11 AS building_container
COPY . /apiserver
WORKDIR /apiserver
RUN go build -o server cmd/main.go

FROM alpine:latest
EXPOSE 8080
COPY --from=building_container /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=building_container /apiserver/ /
WORKDIR /
