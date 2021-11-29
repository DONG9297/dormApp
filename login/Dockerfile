FROM golang:1.17 as builder
# 配置代理
ENV GOPROXY https://goproxy.cn
# 设置go缓存
ENV GO111MODULE=on
# 设置编码格式
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
# go工作目录
WORKDIR /go/release
# 增加所有的项目到该工作路径下
ADD . .
RUN go mod tidy
# 运行命令
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

# 解决镜像太大的问题
FROM alpine:latest
# 安装 /bin/bash 方便调试
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories
RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash
# 安装证书 解决http请求
RUN apk add --update ca-certificates
RUN update-ca-certificates

# 统一亚洲时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezon

# 设置目录
WORKDIR /root/
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#把app文件从“builder”中拷贝到本级的当前目录
COPY --from=builder /go/release/app .
# 暴露端口
EXPOSE 10712
# 项目默认启动命令
CMD ["./app"]