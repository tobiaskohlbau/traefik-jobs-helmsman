FROM golang AS builder

WORKDIR /app
COPY go.mod /app
RUN go mod download
COPY main.go /app
RUN CGO_ENABLED=0 go build

FROM busybox
COPY --from=builder /app/logger /logger
CMD [ "/logger" ]
