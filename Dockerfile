# With this file you can build the docker image. We load the language
# golang:alpine from docker hub load, change the workdir & build the
# application. After that we create a prodution image that based on
# alpine:latest. Copy the files from the build image & start the application
FROM golang:alpine AS build

RUN apk --no-cache add \
    git \
    make

WORKDIR /tmp/src

COPY . .
RUN make build

FROM alpine:latest AS production

WORKDIR /app

COPY --from=build /tmp/src/compver .

ENTRYPOINT ["/app/compver"]