#################################################
# Server builder
#################################################
FROM golang:1.22.4-alpine3.20 AS server-builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server main.go

#################################################
# Final image
#################################################
FROM alpine:3.18

WORKDIR /app
COPY --from=server-builder /app/server .
ENV PORT=80
EXPOSE $PORT
CMD ["./server"]
