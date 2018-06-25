FROM golang

RUN mkdir -p /go/src/goDemo
RUN mkdir -p /go/bin/

WORKDIR /go/src/goDemo
ADD . /go/src/goDemo

RUN go install goDemo

EXPOSE 8080

ENTRYPOINT ["/go/bin/goDemo"]
