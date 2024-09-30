# 操作系统相关

## Linux 命令
 <details>
	<summary>常用Linux命令</summary>

  - [tmux](./opsys/cmd/tmux.md)
  - [dd](./opsys/cmd/dd.md)
  - [ip](./opsys/cmd/ip.md)
  - [ln](./opsys/cmd/ln.md)
  - [samba](./opsys/cmd/samba.md)
  - [系统信息](./opsys/cmd/sysinfo.md)
  - [docker](./opsys/cmd/docker.md)
  - [ifconfig](./opsys/cmd/ifconfig.md)
  - [more/less/head/tail](https://blog.csdn.net/qq_15256443/article/details/81664081)
  - [iptables](https://cloud.tencent.com/developer/article/1628661)
  - [shell脚本语法](./opsys/cmd/shellgram.md)
</details>

## 系统组件

1. [iptables介绍](./opsys/linuxsys/iptables.md)

## 常见问题

1. [忘记root登录密码重置](./opsys/common/pwdforget.md)
2. [Ubuntu vi方向键变为其他乱码](https://blog.csdn.net/a12355556/article/details/120512771)
3. [各管理工具镜像源](./opsys/common/mirrors.md)

## Windows
1. [Windows MSRC查询系统发布补丁](https://api.msrc.microsoft.com/cvrf/v2.0/swagger/index)

## 实用命令
-  shell输出json格式数据
`echo '{"name":"zhangsan", "age":"18"}' | python -m json.tool`
- python 搭建当前目录http服务
`python -m http.server`
- 磁盘清理时查看当前目录下的各个文件（夹）大小
`du -h --max-depth=1 /path/to/directory`
- 取出标准输出的第n列数据做额外处理, 如：
`netstat -tupan |awk '{print $4}'|grep 10.19` // 获取连接消息地址的10.19相关连接
`docker rm $(docker ps -a|awk '{print $1}')` // 也可作为多行变量处理


<br>
<br>

# 代码Demo

## Golang
1. [发送数据至kafka](./codedemo/go/send2kafka.md)
2. [文件遍历Demo(filewalk)](./codedemo/go/filewalk.md)
3. [http Demo](./codedemo/go/http.md)
4. [图片相似度计算](./codedemo/go/imageSim.md)
5. [mysql操作](./codedemo/go/mysqlConn.md)

## C/C++
1. [Windows最小化所有窗口,并显示置顶窗口](./codedemo/c/minWin.md)
2. [忘记密码时的密码尝试(根据/etc/passwd文件)](./codedemo/c/pwdtest.md)

## Python
1. [基于selenium实现网页模拟登陆配置](./codedemo/python/browserSimu.md)

## Rust

## Java

## Bash
1. [awk脚本编写示例](./codedemo/bash/awkscript.md)
2. 

<br>
<br>

# 语言

1. [golang静态分析代码检测优化](./language/go/gostatic.md)
2. [使用go工具链进行fuzzing](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p22)
3. [golang性能分析](./language/go/gocheckperf.md)

<br>
<br>

# 杂谈

1. [计算机编码ASCII/Unicode/UTF-8/GBK](./misc/encoding.md)
2. [shadow文件加密方式解析](https://blog.csdn.net/zwbill/article/details/79322374)
3. [docker网络综述](./misc/dockernet.md)
4. [证书、密钥相关认证](./misc/cert.md)

<br>
<br>

# 实用在线工具
1. [代码生成工具，用于高亮插入word](http://www.codeinword.com/)
2. [画图或ppt制作选取配色](https://color.adobe.com/zh/explore)
3. [画图或ppt制作选取配色2](https://coolors.co/palettes/trending)

<br>
<br>

# 实用Github项目
1. [开源版WAF，雷池](https://github.com/chaitin/SafeLine)
2. 
