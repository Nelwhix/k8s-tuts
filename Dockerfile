FROM golang:1.22.5

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /k8s-go

EXPOSE 3000

CMD ["/k8s-go"]