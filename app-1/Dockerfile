FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd main && go build -o app1
RUN cd ..

EXPOSE 8080

CMD ["/app/main/app1"]
