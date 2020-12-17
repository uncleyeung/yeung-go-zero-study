FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY rpc/add/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/add rpc/add/add.go


FROM centos:7

#RUN apk update --no-cache
#RUN apk add --no-cache ca-certificates
#RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/add /app/add
COPY --from=builder /app/etc /app/etc

CMD ["./add", "-f", "etc/add.yaml"]
