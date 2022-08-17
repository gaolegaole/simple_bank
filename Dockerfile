# Build stage
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY app.env .
COPY wait-for.sh .
COPY start.sh .
COPY db/migration /db/migration
# 如果需要可以增加wait-for.sh等待postgres准备就绪，参考：https://www.bilibili.com/video/BV1dy4y1u7Sq?p=25&spm_id_from=pageDriver&vd_source=b9449ebf8fd2ce5cbdaaa53c27980a04
RUN ["chmod", "+x" ,"/app/wait-for.sh", "/app/start.sh"]
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
