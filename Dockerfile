FROM golang:latest

WORKDIR /comment
COPY . /comment
RUN go build .
EXPOSE 10008
ENTRYPOINT ["./comment"]
