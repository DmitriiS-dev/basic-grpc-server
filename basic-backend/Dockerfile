FROM golang:1.26-alpine
WORKDIR /usr/local/app

# Install git (might bet required for modules sometimes):
RUN apk add --no-cache git

# Installing app dependencies
COPY go.mod go.sum ./

# COPY Contents - run them and set the output path
COPY . .

# Build binary
RUN go build -o server ./cmd/project

# Create non root user:
RUN adduser -D appuser


# Move binary & permissions
RUN chown appuser:appuser /usr/local/app/server
USER appuser

EXPOSE 50051

CMD ["./server"]

