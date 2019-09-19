FROM golang:latest 

RUN mkdir -p /app
WORKDIR /app

ADD . /app

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/go-sql-driver/mysql

RUN go build ./main.go

CMD ["./main"]