#build go server
FROM golang:1.25

WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 9001
CMD [ "go", "run", "main.go" ]