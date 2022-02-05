FROM golang:1.17 AS FirstStage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o main


FROM alpine  
WORKDIR /root
COPY --from=FirstStage /app/main ./
ENTRYPOINT ["./main", "-p 80:80"]  
