FROM golang:1.18 as builder 

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

COPY ./app /app/

RUN go mod download

RUN go build -o ./bin/app ./src

FROM golang:1.18 

COPY --from=builder /app/bin/app /

EXPOSE 8080

CMD ["/app"]
