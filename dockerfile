FROM golang:1.22.1 AS build

COPY . .

RUN go build -o /app/build ./cmd/main.go

FROM ubuntu:22.04 as production

RUN apt-get update && apt-get install -y curl

WORKDIR /app

COPY --from=build /app/build .

CMD /app/build