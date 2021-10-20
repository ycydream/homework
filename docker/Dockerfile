FROM golang:1.16-alpine AS build

WORKDIR /build

ENV GOPROXY https://goproxy.cn
COPY http.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build http.go

FROM scratch

COPY --from=build /build/http /

ENTRYPOINT ["/http"]


