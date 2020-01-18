FROM golang:1.13.6 AS build

WORKDIR /go/src/davidchristie/identity

COPY . .

RUN CGO_ENABLED=0 go build -o /identity

FROM alpine:3.11.2

COPY --from=build /identity /

ENTRYPOINT [ "/identity" ]
