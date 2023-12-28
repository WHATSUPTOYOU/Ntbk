- 字符串比较
	-  
	- 如果字符串可能为空，单括号可能出错，需使用双括号语法
```bash
#!/bin/bash

string1="MyString"
string2="MyString"

if [ "$string1" == "$string2" ]
then
    echo "Equal Stringis"
else
    echo "Strings not equal"
fi
```

<br>

- 字符串截取
	- 
	- \# 删除左边字符，保留右边字符: ```echo ${var#*/}```，*/ 表示从左边开始删除第一个 / 号及左边的所有字符，```echo ${var##*/}```表示从左边开始删除最后（最右边）一个 / 号及左边的所有字符。
	- % 删除右边字符，保留左边字符：```echo ${var%/*}```，%/\* 表示从右边开始，删除第一个 / 号及右边的字符，```echo ${var%%/*}```，表示从右边开始，删除最后（最左边）一个 / 号及右边的字符。
	- 从左边第几个字符开始，及字符的个数，```echo ${var:0:5}```表示从左边第一个字符开始，截取5个字符，不指定```:5```表示直到结束。
	- 从右边第几个字符开始，及字符的个数，```echo ${var:0-7:3}```表示右边算起第七个字符开始，3 表示字符的个数，不指定```:3```表示直到结束。
	
	<br>
	
- 字符串替换
	- 
	
1. `${parameter//pattern/string}`，使用string替换所有pattern
2. 设置IFS分隔符变量：
```bash

#!/bin/bash
 
string="hello,shell,split,test"  
 
#对IFS变量 进行替换处理
OLD_IFS="$IFS"
IFS=","
array=($string)
IFS="$OLD_IFS"
 
for var in ${array[@]}
do
   echo $var
done
```
3. 使用tr命令：
```bash

#!/bin/bash
 
string="hello,shell,split,test"  
array=(`echo $string | tr ',' ' '` )  
 
for var in ${array[@]}
do
   echo $var
done
```

- bash括号
	- 
	- 单小括号：命令替换，数组初始化等
	- 双小括号：高级数学表达式,只要括号中的运算符、表达式符合C语言运算规则，都可用在$((exp))中,如要进行0-4的循环，可以采用```for((i=0;i<5;i++))```，不使用双括号则为```for i in {0..4}``` 或 ```for i in `seq 0 4` ```，再如可以直接使用```if((\$i<5))``` 或 ```if [ $i -lt 5 ]```来表示i < 5的判断。
		符号	| 描述
		:--: | :--:|
		val++	|后增
		val--	|后减
		++val	|先增
		--val	|先减
		！	|逻辑求反
		～	|位求反
		**	|幂求反
		<<	|左位移
		\>>	|右位移
		&&	|逻辑和
    - 单中括号：和test是等同的，两者都是用于字符串比较的，不可用于整数比较，整数比较只能使用-eq，-gt这种形式。无论是字符串比较还是整数比较都不支持大于号小于号。[ ]中的逻辑与和逻辑或使用-a 和-o 表示。或者在一个array 结构的上下文中，中括号用来引用数组中每个元素的编号。
    - 双中括号：1. 使用[[ ... ]]条件判断结构，而不是[ ... ]，能够防止脚本中的许多逻辑错误。比如，&&、||、<和> 操作符能够正常存在于[[ ]]条件判断结构中，但是如果出现在[ ]结构中的话，会报错。比如可以直接使用`if [[ $a != 1 && $a != 2 ]]`, 如果不适用双括号, 则为`if [ $a -ne 1] && [ $a != 2 ]`或者`if [ $a -ne 1 -a $a != 2 ]`。
    - 大括号： 
		- 大括号拓展。(通配(globbing))将对大括号中的文件名做扩展。在大括号中，不允许有空白，除非这个空白被引用或转义。第一种：对大括号中的以逗号分割的文件列表进行拓展。如 touch {a,b}.txt 结果为a.txt b.txt。第二种：对大括号中以点点（..）分割的顺序文件列表起拓展作用，如：touch {a..d}.txt 结果为a.txt b.txt c.txt d.txt。
			```bash
			# ls {ex1,ex2}.sh    
			ex1.sh  ex2.sh    
			# ls {ex{1..3},ex4}.sh    
			ex1.sh  ex2.sh  ex3.sh  ex4.sh    
			# ls {ex[1-3],ex4}.sh    
			ex1.sh  ex2.sh  ex3.sh  ex4.sh 
			```
		- 特殊替换:1、\${var:-string}和\${var:=string}:若变量var为空，则用在命令行中用string来替换\${var:-string}，否则变量var不为空时，则用变量var的值来替换\${var:-string}；对于\${var:=string}的替换规则和\${var:-string}是一样的，所不同之处是\${var:=string}若var为空时，用string替换\${var:=string}的同时，把string赋给变量var： \${var:=string}很常用的一种用法是，判断某个变量是否赋值，没有的话则给它赋上一个默认值。
		2、\${var:+string}的替换规则和上面的相反，即只有当var不是空的时候才替换成string，若var为空时则不替换或者说是替换成变量 var的值，即空值。(因为变量var此时为空，所以这两种说法是等价的)
		3、{var:?string}替换规则为：若变量var不为空，则用变量var的值来替换\${var:?string}；若变量var为空，则把string输出到标准错误中，并从脚本中退出。我们可利用此特性来检查是否设置了变量的值。

			补充扩展：在上面这五种替换结构中string不一定是常值的，可用另外一个变量的值或是一种命令的输出。
			
			```bash
			${var:-string} 
			${var:+string} 
			${var:=string}
			${var:?string}
			```
			
- IF
	- 
	- 常见选项

	选项(操作符)|描述
	:---- | :----: |
	! EXPRESSION | 检查EXPRESSION是否为假。
	-n STRING|检查STRING的长度是否大于零。
	-z STRING|检查STRING的长度是否为零(即为空)
	STRING1 == STRING2|检查STRING1是否等于STRING2。
	STRING1 != STRING2|检查STRING1是否不等于STRING2。
	INTEGER1 -eq INTEGER2|检查INTEGER1在数值上是否等于INTEGER2。
	INTEGER1 -gt INTEGER2|检查INTEGER1在数值上是否大于INTEGER2。
	INTEGER1 -lt INTEGER2|检查INTEGER1在数值上是否小于INTEGER2。
	-d FILE|检查FILE是否存在并且它是一个目录。
	-e FILE|检查FILE是否存在。
	-r FILE|检查FILE是否存在，并授予读取权限。
	-s FILE|检查FILE是否存在并且其大小大于零(表示它不为空)。
	-w FILE|检查FILE是否存在并授予写权限。
	-x FILE|检查FILE是否存在并授予执行权限。

	<br>
