FROM golang:1.10.0-alpine AS build

WORKDIR /go/src/github.com/avegao/iot-api

RUN apk add --no-cache --update \
    git \
    glide

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

RUN glide install

COPY ./ ./

ARG VCS_REF="unknown"
ARG BUILD_DATE="unknown"

RUN go test ./... -cover &&\
    go install \
        -ldflags "-X main.buildDate=$BUILD_DATE -X main.commitHash=$VCS_REF"

########################################################################################################################

FROM alpine:3.7

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENV GRPC_VERBOSITY ERROR

RUN addgroup iot-api && \
    adduser -D -G iot-api iot-api

USER iot-api

WORKDIR /app

COPY --from=build /go/bin/iot-api /app/iot-api

EXPOSE 8080/tcp

LABEL com.avegao.iot.api.vcs_ref=$VCS_REF \
      com.avegao.iot.api.build_date=$BUILD_DATE \
      maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["./iot-api"]
