## ifconfig

设网卡为ens33

ifconfig ens33 up/down  启停网卡
ifconfig eth0 192.168.1.18 netmask 255.255.255.0 设置ip和掩码

route add default gw 192.168.1.1：临时设定网关
route -n 查看路由表
