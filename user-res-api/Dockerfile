FROM golang:1.20

ADD . /user-res-api

WORKDIR /user-res-api
RUN go mod init user-res-api
RUN go mod tidy
RUN go mod download golang.org/x/crypto

RUN go build -o user-res-api
RUN chmod +x /user-res-api

EXPOSE 8070

ENTRYPOINT ["./user-res-api"]