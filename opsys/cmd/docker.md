
## docker

- docker pull image:tag //拉取镜像
- docker images (-a) //查看镜像
 - docker run -it image:tag bash //用bash命令行新建并运行容器
- docker ps //列出当前正在运行的容器
- docker ps -a //所有容器
- docker start/stop containerid//启动或停止
- docker rm 容器//移除容器
- docker attach id //进入容器，标准输入输出
- docker exec id //进入容器执行命令
- docker stats //监控容器资源信息
- docker cp 容器ID:容器内的文件路径 宿主机路径//从容器内拷贝到宿主机
- docker tag docker.io/centos docker.io/centos:v1 //给镜像打标签
- docker inspect ID //查看容器网络信息
- docker的四种网络模式
-- None --- 不为容器进行任何网络配置，容器不能访问外部网络，内部存在回路地址,这个Docker容器没有网卡、IP、路由等信息，只有lo 网络接口。需要我们自己为Docker容器添加网卡、配置IP等。
-- Container --- 将容器的网络栈合并到一起，可与其他容器共享IP地址和端口范围等。而不是和宿主机共享,两个容器除了网络方面，其他的如文件系统、进程列表等还是隔离的。
-- Host --- 与主机共享网络。
-- Bridge --- 默认网络模式，通过主机和容器的端口映射（iptable转发）来通信。桥接是在主机上，一般叫docker0。
-  exit退出并停止容器
- docker commit [CONTAINER ID] [IMAGE NAME]   #容器ID  创建的镜像名 //保存为镜像
-v   参数可以指定映射目录，使得容器中和本机中的文件共享
-  Ctrl+p+q只退出容器，不停止容器

例：Docker运行数据库并本地连接:
1.docker run --name mydb_test -p 3305:3306 -e MYSQL_ROOT_PASSWORD=asdasd -d mariadb:latest  //3305:3306 将本地的3305端口映射到docker中3306端口
2.本地连接数据库3305端口即可

docker修改运行时要执行的命令：
将容器导出为image，重新通过定义执行/bin/bash进入

启动容器并进入容器bash：
docker start  ->  docker  exec  -it  ...  /bin/bash（docker执行的命令不是/bin/bash时，不能用attach）

docker 指定仓库镜像时报错：Error response from daemon: Get https://registry-1.docker.io/v2/: x509: certificate signed by unknown authority
解决：需要下载证书并添加到docker配置下
Step 1: openssl s_client -showcerts -connect ${DOMAIN}:${PORT}</dev/null2>/dev/null|openssl x509 -outform PEM >ca.crt    //获取证书

What I Ran: openssl s_client -showcerts -connect registry-1.docker.io:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >ca.crt

Step 2: sudo cp ca.crt /etc/docker/certs.d/${DOMAIN}/ca.crt  //复制证书，certs.d目录可能需要创建

What I Ran: sudo cp ca.crt /etc/docker/certs.d/registry-1.docker.io/ca.crt

Step 3: cat ca.crt | sudo tee -a /etc/ssl/certs/ca-certificates.crt  //  添加到ssl连接证书中

Step 4: sudo service docker restart


docker   配置容器走本地代理：
修改~/.docker/config.json，添加如下：
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://proxy.example.com:8080",
     "httpsProxy": "http://proxy.example.com:8080",
     "noProxy": "localhost,127.0.0.1,.example.com"
   }
 }
}
重启docker服务

docker 运行参数：
https://haicoder.net/docker/docker-run.html

docker compose 配置文件说明：
https://www.jianshu.com/p/748416621013

清除所有容器：
- docker stop $(docker ps -q)
- docker rm $(docker ps -aq)

ifconfig ens33 up/down  启停网卡
ifconfig eth0 192.168.1.18 netmask 255.255.255.0 设置ip和掩码

route add default gw 192.168.1.1：临时设定网关
route -n 查看路由表
