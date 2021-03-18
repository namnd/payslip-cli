FROM golang:1.15-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . /app

WORKDIR src/cmd/payslip-cli

RUN go build -o main.go .

CMD ["./main.go"]
