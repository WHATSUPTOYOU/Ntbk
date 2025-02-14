# grep/sed/awk
## 说明
这三个命令与linux文本操作相关，grep一般用于基本的文本查找匹配，sed适合编辑匹配到的文本，awk适合对文本进行处理，执行复杂的处理指令或格式调整。

## grep
一般常见的使用方式：
``grep XXX file //查找指定文件的指定内容行，XXX可以是正则``
``grep -r "XXX" /path/to/search //递归查找包含指定字符串的文件内容``

## sed
一般常见的使用方式：
``sed -n "/XXX/p" file //只打印匹配到XXX的行，去掉-n打印所有``
``sed (-i) "s/XXX/OOO/(g)" file //替换XXX为OOO，-i表示是否写入原文件，不带g表示仅替换一行中第一个匹配到的内容``
``sed -n "2p" file //打印第二行``

## awk
一般常见的使用方式：
``awk -F: '{print $1,$2}' file  //以冒号分隔，打印每一行的前两个分割结果``

## 例
> 使用实例
> 1. 查找目录下包含指定字符(忽略大小写)的所有文件，将文件中的指定字符串替换为另一个字符串
> 实现：`grep -r -i "XXX" .| awk -F: {print $1}|sort |uniq | xargs -I FILE sed 's/XXX/OOO/g' FILE`

## TODO
awk是一个非常强大的文本处理工具，需要时间阅读