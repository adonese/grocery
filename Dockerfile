FROM golang:latest

COPY . /app
WORKDIR /app
RUN go build
CMD ["/grocery"]
EXPOSE 6661