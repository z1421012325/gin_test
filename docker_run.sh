#!/bin/bash
# docker build 根据当前目录的Dockerfile文件内容制作镜像 . 代表当前目录(可以绝对路径)  -t表示设置镜像名字
docker build -t myweb .

# 加载镜像启动容器 根据镜像名字设置启动模式 -itd 表示后台命令行模式... -p在宿主机和容器开启的端口连接一起
# /bin/bash 表示在容器内部执行一些命令但不进入 -c 如果需要找容器内部执行命令 必须有-c 后面这是执行命令 启动容器内的./gin_test
docker run -itd -p 1111:7999 --name="myweb_docker_one" myweb /bin/bash -c ./gin_test

# 展示docker容器的执行状况
docker ps



# 持续监控docker容器内部的输出日志 -t 监控日志 --tail 只显示多少条
# docker logs -t --tail=20 myweb_docker_one

# 进入容器内部 docker exec -it 容器名字 bash 执行命令行模式 如何有其他需求可以进入内部执行,不过一般都在创建容器阶段执行了
# docker exec -it myweb_docker_one bash

# 进入容器内部执行 接上面的命令
# ./gin_test












# 上传镜像到hub.docker
#前提 docker 是login 过后的  退出登录状态 logout
#docker将容器转为镜像 将镜像上传至hub.docker仓库
#
#将容器封装为镜像
#docker commit -a="作者信息" -m="提交信息" 容器id 设置打包的镜像名字
#
#将打包的镜像设置tag 属性(上传仓库)
#docker tag 刚刚打包设置的镜像名字 仓库名字(比如我的是z1421012325/gowebs 我的仓库名是这样的 如果没有设置仓库,那么用户名/xxx 自动创建一个用户名/xxx的仓库)
#
#docker images 能看到有一个仓库名的镜像和被打包成设置的镜像的两个id一样 不要慌 都是一个
#docker push 在tag设置的仓库名