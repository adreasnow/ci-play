FROM golang:latest

WORKDIR /app

ADD views /app/views
ADD main.go /app/

RUN go mod init test
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

RUN go build main.go

ENV DATABASE_URL=postgres://forecast:forecast@db:5432/forecast?sslmode=disable

EXPOSE 8080/tcp

ENTRYPOINT ["/app/main"]