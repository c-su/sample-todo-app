FROM golang:latest

RUN go get github.com/oxequa/realize
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/mattn/go-sqlite3

RUN mkdir /go/src/work
WORKDIR /go/src/work
COPY . .

EXPOSE 8080

ENTRYPOINT [ "realize", "start", "--run", "--no-config" ]