FROM golang:latest 

RUN mkdir -p /app
WORKDIR /app

ADD . /app

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/dgrijalva/jwt-go

RUN go get github.com/davidfengg/PatientTracker/controllers
RUN go get github.com/davidfengg/PatientTracker/database
RUN go get github.com/davidfengg/PatientTracker/login
RUN go get github.com/davidfengg/PatientTracker/models
RUN go get github.com/davidfengg/PatientTracker/route

RUN go build -o main .

CMD ["./main", "-ip=docker.for.mac.localhost:3306"]