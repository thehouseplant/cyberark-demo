FROM golang:1.12

RUN go get -u github.com/cyberark/conjur-api-go/conjurapi
RUN go build main.go

EXPOSE 5000
CMD [ "./main" ]