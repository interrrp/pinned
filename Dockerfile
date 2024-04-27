FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -ldflags="-s -w" -o pinned .

CMD [ "./pinned" ]
