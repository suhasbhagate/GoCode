FROM golang:1.18.4-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8181
RUN go build -o main .
CMD ["/app/main"]