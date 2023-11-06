FROM golang:1.18

ADD . /hotels-api

WORKDIR /hotels-api
RUN go mod init hotels-api
RUN go mod tidy
RUN go build -o hotels-api
RUN chmod +x /hotels-api

EXPOSE 8060

ENTRYPOINT ["./hotels-api"]