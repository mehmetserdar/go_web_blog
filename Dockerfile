FROM golang:alpine3.12
WORKDIR /go/src/goweb
COPY . .
CMD ["/go/src/goweb/go_web_blog"]