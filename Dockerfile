FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /go-side

EXPOSE 8888

CMD [ "/go-side" ]