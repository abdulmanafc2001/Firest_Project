# build stage
FROM golang:1.20.10-alpine3.17 AS builder
WORKDIR /app
ADD . .
RUN go build -o main main.go

# run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY . .
COPY app.env .

EXPOSE 3000
CMD [ "/app/main" ]