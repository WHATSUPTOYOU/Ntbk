
##### 基于密钥的ssh登录
1. 登录客户端生成密钥对
```
	ssh-keygen
```
2. 安装公钥配置（假设为root配置免密登录）
```
	修改/root/.ssh/authorized_keys，添加生成的id_rsa.pub
	$ chmod 600 authorized_keys
	$ chmod 700 /root/.ssh/
```
3. 添加私钥为默认使用私钥，配置~/.ssh/config，将私钥放置在~/.ssh路径下（Windows/Linux通用）

##### 服务证书生成
1. 生成ca的密钥
```
	openssl genrsa -des3 -out ca.key 2048

```
2. 使用ca密钥生成ca证书
```
	openssl req -new -x509 -days 3650 -key ca.key -out ca.crt

```
3. 生成服务端密钥
```
	openssl genrsa -out server.key 2048
```
4. 生成csr文件
```
	openssl req -new -out server.csr -key server.key

```
5. 生成服务端证书
```
	openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650

```