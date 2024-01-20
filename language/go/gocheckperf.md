## 工具 
- **pprof**

## 代码中导入	
- 在代码中使用pprof监控：
``` go
import (
	_ "net/http/pprof"
	"net/http"
)

// 在main函数中
go func() {
  log.Println(http.ListenAndServe("localhost:8080", nil))
}()
```

## 查看资源占用信息
使用命令 ``go tool pprof http://10.19.201.231:9009/debug/pprof/XXX`` 获取对应资源情况，web命令可以在web页面以火焰图形式展示。

- allocs：过去所有内存分配的采样

- block：导致同步基元上出现阻塞的堆栈跟踪

- cmdline：当前程序的命令行调用

- goroutine：所有当前goroutine的堆栈跟踪

- heap：对活动对象的内存分配进行采样。可以指定gc GET参数以在获取堆样本之前运行gc。

- mutex：资源锁竞争的堆栈跟踪

- profile：CPU采样占用

- trace：对当前程序执行情况的跟踪，可以在seconds GET参数中指定持续时间。获取trace文件后，使用go tool trace命令来调查跟踪。
