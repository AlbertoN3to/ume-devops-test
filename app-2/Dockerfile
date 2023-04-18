FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN cd main && go build -o app2
RUN cd ..

EXPOSE 9090

CMD ["/app/main/app2"]
