FROM golang:1.20-alpine

RUN apk update \
    && apk --no-cache --update add build-base \
    && apk add libffi-dev

WORKDIR /usr/src/infoliticas

COPY . .

CMD [ "go run *.go" ]