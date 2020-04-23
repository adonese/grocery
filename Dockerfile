FROM golang:latest

COPY . /app
WORKDIR /app
RUN go build
CMD ["/app/grocery"]
EXPOSE 6661