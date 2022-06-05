FROM golang:1.18 as builder 

WORKDIR /

COPY . /

RUN go build -o ./app .

CMD ["/app"]
