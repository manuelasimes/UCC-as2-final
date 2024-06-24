FROM golang:1.21.1

ADD . /containers-api

WORKDIR /containers-api
RUN go mod tidy
RUN go build -o containers-api
RUN chmod +x /containers-api

EXPOSE 8040

ENTRYPOINT ["./containers-api"]