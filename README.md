#	DOCKER

##	install

###		install from DaoCloud

		curl -sSL https://get.daocloud.io/docker | sh

###		run deamon

		sudo docker -d &

###		配置Docker加速器

		echo "DOCKER_OPTS=\"\$DOCKER_OPTS --registry-mirror=http://60026668.m.daocloud.io\"" | sudo tee -a /etc/default/docker
		
		sudo service docker restart

##	link

Docker Container连接

1.	构建Docker镜像，服务名称：spring:server 	Dockerfile路径：.

		docker build -t spring:server .

2.	启动服务spring:server，服务名：spring_server

		docker run -d --name spring_server spring:server

3.	启动新容器ubuntu:12.04，连接服务，（映射别名：server）

		docker run -i -t --link spring_server:server ubuntu:12.04 /bin/bash

4.	查看环境变量

		env

服务spring:server的环境变量以（别名server）SERVER开始。

##	net

Docker 网络

1.	host

		docker run --net host --name spring_s1 -i -t spring:v1

##	run

1.	EXPOSE

Dockerifle:

		EXPOSE 9000:9000 80:80

配合命令

		docker run -p 9000:9000 -p 80:80 C:latest

## stop

ctrl+p+q : quietly quit

ctrl+c : exit, stop the container and exit

## ID

docker ps -a -q