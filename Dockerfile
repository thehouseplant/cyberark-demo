FROM golang:1.12

WORKDIR /app

COPY go.mod go.sum ./
COPY run.sh ./

RUN go mod download

COPY . .

RUN go build -o main.go .

EXPOSE 5000

CMD [ "sh" "./run.sh" ]