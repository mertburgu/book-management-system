FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go && ls -l main

FROM golang:1.22
WORKDIR /root/
COPY --from=build /app/main .

COPY wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

EXPOSE 8082

CMD ["sh", "-c", "wait-for-it db:5432 -- ./main"]
