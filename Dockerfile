# 下载基础镜像 这相当于电脑windows系统 所有一切都要在该系统上操作
FROM ubuntu:latest


# WORKDIR 指定容器中的工作目录,默认文件都是在哪保存执行
# WORKDIR /webapp/


# 暴露镜像的端口供主机做映射,不过一般都是直接docker run ... -p 宿主机端口:容器端口  ... 设置了  可以设置多个
EXPOSE 7999
EXPOSE 8888


# COPY 将主机的文件复制到镜像内      其中ADD和COPY一样 但是ADD会对压缩文件（tar, gzip, bzip2, etc）做提取和解压操作
# COPY如果目的位置不存在，Docker会自动创建所有需要的目录结构，但是它只是单纯的复制，并不会去做文件提取和解压工作
COPY gin_test .
COPY gintest.env .



# VOLUME 用来向基于镜像创建的容器添加卷。
# 比如你可以将mongodb镜像中存储数据的data文件指定为主机的某个文件。(容器内部建议不要存储任何数据)
# VOLUME 主机目录 容器目录
# VOLUME /data/db /data/configdb


# 镜像内部的执行... 下载数据软件等...           并不是在宿主机上执行 而是到镜像中执行的命令
# 比如这些命令RUN、COPY、ADD、EXPOSE、WORKDIR、ONBUILD、USER、VOLUME
# 容器内部执行apt更新
# RUN ["apt-get" "update"]


# 容器启动时执行指令 ENTRYPOINT
# CMD docker run -it -d -p 7999:7999 ubuntu


