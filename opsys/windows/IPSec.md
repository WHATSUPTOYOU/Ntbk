# Windows netsh IPSec 命令

## 策略建立
- Netsh ipsec static add policy name = Michael's安全策略   //建立策略
- Netsh ipsec static add filteraction name = 阻止 action =block  //建立动作
- Netsh ipsec static add filterlist name =可访问的终端列表  //建立过滤器列表
- Netsh ipsec static add filter filterlist = 可访问的终端列表 srcaddr= //添加过滤器
- Netsh ipsec static add rule name =可访问的终端策略规则 Policy = Michael's安全策略 filterlist =可访问的终端列表 filteraction = 阻止 //建立规则
- netsh ipsec static set policy name = Michael assign = y //激活使用

## 策略查看
- netsh ipsec static show filterlist name=EDRFilterPermitList1 verbose //查看过滤器列表详细信息
- netsh ipsec static show filterlist all verbose //查看所有过滤器
- netsh ipsec static show rule all policy=EDRPolicy

## 策略删除
- netsh ipsec static delete policy name=EDRPolicy
- netsh ipsec static delete filterlist name=EDRFilterBlockList1
- netsh ipsec static delete filterlist name=EDRFilterPermitList1