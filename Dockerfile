FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

ADD . .

# 安装 pandoc
RUN apk add --no-cache pandoc



RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go get -u github.com/swaggo/swag/cmd/swag


RUN go build main.go

RUN chmod +x main

# swagger init
#RUN ls  $GOPATH/bin/




#RUN  #$GOPATH/bin/swag init



EXPOSE 8080

CMD ["./main"]


