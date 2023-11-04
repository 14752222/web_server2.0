FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

ADD . .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build main.go

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]


