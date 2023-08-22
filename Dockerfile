from golang:1.21

WORKDIR /app/
RUN go get -u github.com/swaggo/swag/cmd/swag

CMD ["/bin/sh", "-c", "tail -f /dev/null"]
