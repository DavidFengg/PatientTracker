FROM golang:latest 

RUN mkdir -p /app
WORKDIR /app

ADD . /app

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/dgrijalva/jwt-go

RUN go get github.com/davidfengg/restAPI/controllers
RUN go get github.com/davidfengg/restAPI/database
RUN go get github.com/davidfengg/restAPI/login
RUN go get github.com/davidfengg/restAPI/models
RUN go get github.com/davidfengg/restAPI/route

RUN go build -o main .

CMD ["./main"]