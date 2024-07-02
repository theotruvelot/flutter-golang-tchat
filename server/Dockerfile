FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server cmd/main.go

EXPOSE 1234
CMD ["/server"]

