FROM golang:1.12

WORKDIR /app

COPY go.mod go.sum ./
COPY run.sh ./

RUN go mod download

COPY . .

RUN go build -o main.go .
RUN chmod +x ./run.sh && ./run.sh

EXPOSE 5000
