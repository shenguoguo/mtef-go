FROM golang:alpine

WORKDIR $GOPATH/mathtype2latex
ADD . $GOPATH/mathtype2latex

RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o  app.bin  main.go

EXPOSE 9537
ENTRYPOINT ["./app.bin"]
