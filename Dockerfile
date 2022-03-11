FROM golang:1.16 as base-image
# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download

ADD . .
RUN go build -o bin/app main.go


FROM alpine:latest

WORKDIR /

COPY --from=base-image /server/bin .
COPY --from=base-image /server/db/migrations ./db/migrations

#Fills the gap between docker image server and the person who runs the container
EXPOSE 8091
ENTRYPOINT ["./app"]