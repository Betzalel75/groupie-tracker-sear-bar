FROM golang:alpine

LABEL maintainers="eibounda, fatidiop, jefaye"

WORKDIR /app

COPY . .

RUN bash
RUN go build -o groupie-tracker

EXPOSE 8080

CMD ["./groupie-tracker"]