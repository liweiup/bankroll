FROM golang:alpine as builder
# 声明工作目录
WORKDIR /go/src/bankroll
# 拷贝整个server项目到工作目录
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
# go generate 编译前自动执行代码
# go env 查看go的环境变量
# go build -o server . 打包项目生成文件名为server的二进制文件
RUN go generate && go env && go build -o server .

# 声明镜像来源为alpine:latest
FROM alpine:latest
# 镜像编写者及邮箱
LABEL MAINTAINER="321327476@qq.com"
# 声明工作目录
WORKDIR /go/src/bankroll
# 把/go/src/bankroll整个文件夹的文件到当前工作目录
COPY --from=0 /go/src/bankroll ./
#
#COPY --from=0 /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip
#ENV ZONEINFO /opt/zoneinfo.zip
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

EXPOSE 8010
# 运行打包好的二进制 并用-c 指定config.docker.yaml配置文件
ENTRYPOINT ./server -c config-beta.yaml


