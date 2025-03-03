FROM golang:1.23-alpine

WORKDIR /api

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o bin/api main.go

EXPOSE 3333

CMD ["/api/bin/api", "3333"]
