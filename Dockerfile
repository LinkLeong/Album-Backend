FROM golang:1.20 AS builder

WORKDIR /app

COPY . .
# CGO_ENABLED=0 一定不要忘记,会因为依赖库的问题导构建后无法运行
RUN CGO_ENABLED=0 go build -v -o /go/bin/api

# final (target) stage
# FROM scratch
FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/api /go/bin/api
RUN mkdir -p /images
# 声明volume挂载点
VOLUME ["/images"]
LABEL VERSION="0.3"
WORKDIR /images

EXPOSE 8081

ENTRYPOINT ["/go/bin/api","-d","/images"]
