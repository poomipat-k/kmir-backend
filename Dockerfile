# Build Stage

FROM golang:1.22.4-alpine3.20 as buildStage

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN env GOOS=linux CGO_ENABLED=0 go build -o /myApp ./

# Deploy Stage

FROM alpine:latest

WORKDIR /

COPY --from=buildStage /myApp /app

EXPOSE 8080

ENTRYPOINT [ "/app" ]