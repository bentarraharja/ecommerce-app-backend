FROM golang:1.21.5-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o server .

CMD [ "/app/server" ]
