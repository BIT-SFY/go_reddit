# 导入一个golang的镜像
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on  \
    CGO_ENABLE=0    \
    GOOS=windows  \
    GOARCH=amd64

# 转换工作目录到: /build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件 .是指编译当前目录下的所有文件
RUN go build -o app .

# 转换工作目录到 /dist目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里 .指当前路径
RUN cp /build/app .

# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
CMD [ "/build/app" ]