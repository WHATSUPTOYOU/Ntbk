## 服务端
修改/etc/samba/smb.conf文件，添加如下结构共享文件配置：
[myshare]
comment=my share directory
path=/home/gao
browseable=yes
public=yes
writable=yes

smbpasswd -a user 增加smb用户

---

## 客户端
列出共享文件夹：
smbclient -L 198.168.0.1 -U username%password

连接smb
smbclient //192.168.0.1/tmp  -U username%password


