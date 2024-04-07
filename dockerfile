FROM golang:1.22.1 AS build

COPY . .

RUN go build -o /app/build ./cmd/main.go

FROM ubuntu:22.04 as production

WORKDIR /app

COPY --from=build /app/build .

CMD /app/build