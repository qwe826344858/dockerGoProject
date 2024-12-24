# 使用 golang 作为基础镜像
FROM golang:1.23.4

# 设置工作目录
WORKDIR /usr/src/app/dockerGoProject

# 创建工作目录
RUN mkdir "/usr/src/app/dockerGoProject/log"

# 复制Ao服务到容器中
COPY DockerGoProjectAoServer /usr/src/app/dockerGoProject

ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=.gitlab.com,.gite.com

RUN echo "start DockerGoProjectAoServer"
# 运行AO服务
# CMD ["go","run", "/usr/src/app/dockerGoProject/DockerGoProjectAo/server/server.go",">>","/usr/src/app/dockerGoProject/std_out.log","2>&1","&"]
CMD ["sh","-c","./DockerGoProjectAoServer >> /usr/src/app/dockerGoProject/log/docker_dockergoprojectao_stdout.log 2>&1"]

