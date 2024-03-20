FROM golang:1.19

WORKDIR /server

COPY . .

RUN apt-get update 
RUN apt-get install sqlite3

EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]