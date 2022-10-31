FROM golang:1.19.2-alpine
WORKDIR /app
COPY . .
RUN go build ./cmd/apiserver/main.go
CMD [ "/app/main", "server" ]