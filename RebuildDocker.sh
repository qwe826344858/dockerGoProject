# shell

sudo docker rm dockergoprojectao
sudo docker rmi zoneslee/dockergoproject:v1.0

# 打包成二进制可执行文件
go mod tidy
go build -o DockerGoProjectAoServer ./DockerGoProjectAo/server/server.go


sudo docker build -t zoneslee/dockergoproject:v1.0 .

#sudo docker run --name dockergoprojectao -p 40001:40001 zoneslee/dockergoproject:v1.0 &
# 容器日志路径挂载宿主机日志目录中:(-v)  容器指定端口:-p
sudo docker run --name dockergoprojectao -p 40001:40001 -v /home/lighthouse/LogInfo/DockerGoLog:/usr/src/app/dockerGoProject/log zoneslee/dockergoproject:v1.0
# 进入容器内部
# sudo docker run  -it --name dockergoprojectao zoneslee/dockergoproject:v1.0 /bin/bash
