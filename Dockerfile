FROM golang:latest

LABEL maintainer="Zach Hanson <zachary_hanson@comcast.com>"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . /app

# RUN cd cmd/toyrobot && go build -o main .

RUN go build -o /app/main cmd/toyrobot/main.go

EXPOSE 8080

# CMD ["/app/cmd/toyrobot/main"]

CMD ["/app/main"]