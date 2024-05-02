

# Build stage
FROM golang:1.22.2-alpine3.19 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]