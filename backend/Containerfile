FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o is-towel-day .

FROM scratch

COPY --from=builder /app/is-towel-day /is-towel-day

EXPOSE 8080

CMD ["/is-towel-day"]
