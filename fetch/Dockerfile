FROM golang:1.13.0

WORKDIR /go/src/app

COPY . .

RUN go get -d -v .
RUN go install -v .

CMD ["app"]