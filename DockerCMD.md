#	DOCKER

##	link

Docker Container连接

1.	构建Docker镜像，服务名称：spring:server 	Dockerfile路径：.

docker build -t spring:server .

2.	启动服务spring:server，服务名：spring_server

##docker run -d --name spring_server spring:server##

3.	启动新容器ubuntu:12.04，连接服务，（映射别名：server）

##docker run -i -t --link spring_server:server ubuntu:12.04 /bin/bash##

4.	查看环境变量

env

服务spring:server的环境变量以（别名server）SERVER开始。

