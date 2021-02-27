FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o app calc/calc_server/server.go
#RUN go build -o app calc/calc_client/client.go
#RUN go build -o app calc/avg_server/server.go
#RUN go build -o app calc/avg_client/server.go