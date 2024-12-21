# Dockerfile for custom base image
FROM python:3.11

# 设置工作目录
WORKDIR /usr/src/app/

# 将依赖文件复制到工作目录
COPY requirements.txt ./


RUN pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/
RUN unset http_proxy
RUN unset https_proxy

# 安装依赖
RUN pip install --no-cache-dir -r requirements.txt

# 设定一个环境变量，指示基础镜像安装了依赖
ENV DEPENDENCIES_INSTALLED=True

# 之后你可以继续添加其他配置

