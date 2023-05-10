FROM golang:1.20-alpine3.17

RUN mkdir /code
WORKDIR /code
COPY cmd/api/* /code/cmd/api/
COPY main.go /code/
COPY go.mod /code/
COPY go.sum /code/
# ssh
RUN go mod download
RUN go build -o main .
RUN rm -rf cmd

EXPOSE 8000 2222
ENTRYPOINT ["./main"]
CMD []