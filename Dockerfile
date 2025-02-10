FROM golang:1.23-alpine

WORKDIR /loier

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

CMD ["go", "run", "main.go"]