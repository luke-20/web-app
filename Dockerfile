FROM golang:1.17

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build


FROM alpine  
WORKDIR /root/
COPY --from=0 /app/main ./
CMD ["./main", "-p 80:80"]  
