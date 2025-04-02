FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o dns-server ./cmd/server/main.go

EXPOSE 50051

CMD ["./dns-server"]