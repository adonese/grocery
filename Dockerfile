FROM golang:latest

COPY . /app
WORKDIR /app
RUN go build
EXPOSE 6661
CMD ["/app/grocery"]
