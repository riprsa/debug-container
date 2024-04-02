FROM golang:1.22.1-alpine3.19 as build
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /app

FROM scratch
COPY --from=build /app /app
ENTRYPOINT ["/app"]