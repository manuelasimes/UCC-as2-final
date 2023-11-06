FROM golang:1.18

ADD . /user-res-api

WORKDIR /user-res-api
RUN go mod init user-res-api
RUN go mod tidy
RUN go build -o user-res-api
RUN chmod +x /user-res-api

EXPOSE 8070

ENTRYPOINT ["./user-res-api"]