FROM golang:1.24-alpine AS build_stage

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -v -o /usr/local/bin/loier ./main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=build_stage /usr/local/bin/loier /app/loier

EXPOSE 8080

ENTRYPOINT [ "/app/loier" ]