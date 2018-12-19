FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/sundogrd/content
COPY . /usr/local/go/src/github.com/sundogrd/content
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/content ./content


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/sundogrd/content/build/content /bin/content
CMD ["content", "up"]
