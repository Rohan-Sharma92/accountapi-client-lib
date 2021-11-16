FROM golang:1.16

ENV CGO_ENABLED 0
ENV GO111MODULE=on
RUN set -xe \
    && mkdir /app

COPY go.* /app/

WORKDIR /app

RUN set -xe \
    && go mod download

CMD ["go", "test", "-v", "./..."]
