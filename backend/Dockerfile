FROM golang:1.14-alpine
COPY ./ /workspace
WORKDIR /workspace
RUN apk update \
    && apk add --no-cache mysql-client\
    && go get github.com/pilu/fresh
CMD [ "fresh", "main.go" ]