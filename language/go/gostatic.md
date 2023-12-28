# Golang静态分析

## Go编译过程

Go的编译过程和其他静态语言类似：
词法分析 ➡ 语法分析 ➡ 语义分析 ➡ IR生成 ➡代码优化 ➡ 机器码生成

## 词法分析
 词法分析阶段实现对源代码的词法提取，获取构成的tokens，Go提供了词法分析的扫描器 go/scanner

代码示例：

*源码：*
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, world!")
}
```

*词法分析结果：*
```go
hello.go:1:1    package "package"
hello.go:1:9    IDENT   "main"
hello.go:1:14   ;       "\n"
hello.go:3:1    import  "import"
hello.go:3:8    STRING  "\"fmt\""
hello.go:3:14   ;       "\n"
hello.go:5:1    func    "func"
hello.go:5:6    IDENT   "main"
hello.go:5:10   (       ""
hello.go:5:11   )       ""
hello.go:5:13   {       ""
hello.go:6:2    IDENT   "fmt"
hello.go:6:5    .       ""
hello.go:6:6    IDENT   "Println"
hello.go:6:13   (       ""
hello.go:6:14   STRING  "\"Hello, world!\""
hello.go:6:29   )       ""
hello.go:6:31   ;       "\n"
hello.go:7:1    }       ""
hello.go:7:3    ;       "\n"
```

## 语法分析

语法分析阶段将词法分析获得的token通过语法规则转换为抽象语法树AST，语法树主要节点类型如下：
```
 Decl，声明 
 	GenDecl，类型声明（import，constant，type，变量） 
 	FuncDecl，函数声明 
 Stmt，语句 
 	IfStmt、ForStmt、ReturnStmt，流程控制语句  
 	BlockStmt，代码块 
 	ExprStmt，表达式语句 
 Expr，表达式 
 	BinaryExpr，二元表达式（X、操作符、Y） 
 	CallExpr，调用函数
  	…
```

## Golang提供的分析接口

Golang提供的工具包汇总如下：
- go/scanner，词法分析
- go/token，token定义
- go/parser，语法分析
- go/ast，AST结构定义
- golang.org/x/tools/go/packages，一组包检查和分析
- golang.org/x/tools/go/ssa，SSA分析
- golang.org/x/tools/go/callgraph，调用关系算法和工具
- golang.org/x/tools/go/analysis，静态分析工具

## golangci-lint集成分析工具

- golangci-lint 集成了很多关于go的代码分析工具
	- Ref： https://golangci-lint.run/
    - 可配置使用的分析工具及对应参数配置
         - 如默认开放的linters：errcheck/gosimple/govet等（检测结果）
    - 可自定义linter
        - Ref：https://golangci-lint.run/contributing/new-linters/

### lint示例 - bodyclose

bodyclose集成在golangci-lint中，检查是否存在未关闭的http response body。

实现方式为：遍历所有引用了Response的包中的所有Instruction，分为下面三种情况：
-  直接使用，ssa.FieldAddr 类型引用
   - 通过定义变量的传递直接进行关闭判断
- 指针引用，ssa.Store 类型引用
   - 形如变量被某个闭包函数引用，会产生一个ssa.Store的引用，我们需要判断被引用后是否调用了Close方法
- 函数调用，ssa.Call /ssa.Defer 类型引用
  - resp获取后，如果被用于参数传递给了其他函数，会产生一个ssa.Call的引用，那就递归到对应的函数中进行判断


## 总结

1. 在编译的各个阶段，都可以实现不同粒度的检测分析工作，一般在AST或SSA IR上进行
2. Golang 当前的代码检测工具golangci-lint集成了大量的检测工具（工具覆盖编写规范、代码逻辑、安全性问题等范畴），支持用户自定义配置，选择需要的检测工具及工具配置，或者添加其他的检测工具
3. Golang 提供了可调用接口和检测框架帮助用户实现自定义的代码分析
4. 为追求检测的效率，有时会需要放弃一些精度，需要根据实际业务需求选择或开发合适的检测方案
