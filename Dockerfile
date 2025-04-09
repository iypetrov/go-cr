FROM golang:1.24.2 AS build-stage
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main .

FROM gcr.io/distroless/base-debian12 AS run-stage
CMD ["/bin/main"]
