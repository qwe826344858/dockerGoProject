# 使用 golang 作为基础镜像
FROM golang:1.23.4

# 设置工作目录
WORKDIR /usr/src/app/dockerGoProject

# 复制Ao服务到容器中
COPY DockerGoProjectAoServer /usr/src/app/dockerGoProject

ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=.gitlab.com,.gite.com

# 运行 main.py
# CMD ["go","run", "/usr/src/app/dockerGoProject/DockerGoProjectAo/server/server.go",">>","/usr/src/app/dockerGoProject/std_out.log","2>&1","&"]
CMD ["./DockerGoProjectAoServer",">>","std_out.log","2>&1"]

