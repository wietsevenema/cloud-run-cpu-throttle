FROM golang:1.13 AS gobuilder

WORKDIR /app
COPY . .

ENV GOOS linux
ENV GOARCH amd64
RUN go build -o main cpu-throttle

FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=gobuilder /app/main /app/main

ENTRYPOINT ["./main"]

