FROM golang:latest
WORKDIR /go/src/oxylus
RUN go get -d -v github.com/labstack/echo
RUN go get -d -v github.com/labstack/echo/middleware
RUN go get -d -v github.com/satori/go.uuid
RUN go get -d -v github.com/timshannon/bolthold
RUN go get -d -v gopkg.in/mgo.v2
COPY . /go/src/oxylus
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o oxylus .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/oxylus .
CMD ["./oxylus"]  