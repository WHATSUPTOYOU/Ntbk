# 各管理工具的镜像切换

## - apt -
- 路径：/etc/apt/source.list
- 常用源
	- 阿里
		```
		deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
		deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
		deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
		deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
		deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
		deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
		deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
		deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
		deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
		deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
		```
	- 中科大源
		```
		deb https://mirrors.ustc.edu.cn/ubuntu/ bionic main restricted universe multiverse
		deb https://mirrors.ustc.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
		deb https://mirrors.ustc.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
		deb https://mirrors.ustc.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
		deb https://mirrors.ustc.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
		deb-src https://mirrors.ustc.edu.cn/ubuntu/ bionic main restricted universe multiverse
		deb-src https://mirrors.ustc.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
		deb-src https://mirrors.ustc.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
		deb-src https://mirrors.ustc.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
		deb-src https://mirrors.ustc.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
		```
	- 清华源
		```
		deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
		deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
		deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
		deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
		deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
		deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
		deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
		deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
		deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
		deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-proposed main restricted universe multiverse
		```
		
		<br>
		
## - yum -
- 路径：/etc/yum.repos.d/CentOS-Base.repo
- 常用源
	 - 网易
		```
		yum -y install wget
		rm -rf /etc/yum.repo.d/*
		wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.163.com/.help/CentOS7-Base-163.repo
		yum clean all && yum makecache
		```
	- 阿里云
		```
		yum -y install yum-utils
		yum-config-manager --add-repo http://mirrors.aliyun.com/repo/Centos-7.repo
		mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo_bak
		mv /etc/yum.repos.d/Centos-7.repo /etc/yum.repos.d/CentOS-Base.repo
		yum clean all
		yum makecache
		yum repolist
		```
	- 中科大
		```
		yum -y install wget
		rm -rf /etc/yum.repos.d/*
		wget -O /etc/yum.repos.d/CentOS-Base.repo 'https://lug.ustc.edu.cn/wiki/_export/code/mirrors/help/centos?codeblock=3'
		yum clean all && yum makecache
		```


## - python pip -
- 路径：~/.pip/pip.conf
- 配置格式
	```
	[global]
	index-url = http://af.hikvision.com.cn/artifactory/api/pypi/pypi/simple/
	trusted-host = af.hikvision.com.cn
	```
- 常用源
	```
	阿里云 http://mirrors.aliyun.com/pypi/simple/
	中国科技大学 https://pypi.mirrors.ustc.edu.cn/simple/
	豆瓣(douban) http://pypi.douban.com/simple/
	清华大学 https://pypi.tuna.tsinghua.edu.cn/simple/
	```
- 临时切换
```pip install markdown -i https://pypi.tuna.tsinghua.edu.cn/simple```
	<br>
	
##  - go -
- 配置
	-  go env -w GOPROXY="source"
-  常用源
	```
	阿里云：https://mirrors.aliyun.com/goproxy
	微软：https://goproxy.io
	七牛云：https://goproxy.cn
	GoCenter：https://gocenter.io
	```
	<br>
	
## - rust cargo -
- 配置文件路径：~/.cargo/config
- 配置格式及常用源
	```
	[source.crates-io]
	registry = "https://github.com/rust-lang/crates.io-index"
	# 指定镜像
	#replace-with = 'tuna' # 如：tuna、sjtu、ustc，或者 rustcc

	# 注：以下源配置一个即可，无需全部

	# 中国科学技术大学
	[source.ustc]
	registry = "git://mirrors.ustc.edu.cn/crates.io-index"

	# 上海交通大学
	[source.sjtu]
	registry = "https://mirrors.sjtug.sjtu.edu.cn/git/crates.io-index"

	# 清华大学
	[source.tuna]
	registry = "https://mirrors.tuna.tsinghua.edu.cn/git/crates.io-index.git"

	# rustcc社区
	[source.rustcc]
	registry = "https://code.aliyun.com/rustcc/crates.io-index.git"
	```
