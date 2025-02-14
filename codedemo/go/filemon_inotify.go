package main

import (
        "crypto/sha256"
        "flag"
        "fmt"
        "io"
        "os"
        "path/filepath"
        "strings"

        "github.com/fsnotify/fsnotify"
)

// 检查路径是否需要排除
func shouldExclude(path string, excludePaths []string) bool {
        for _, exclude := range excludePaths {
                if strings.HasPrefix(path, exclude) {
                        return true
                }
        }
        return false
}

func main() {
        // 命令行参数
        var watchPath string
        var exclude string
        flag.StringVar(&watchPath, "path", "", "监控的目录或文件路径，用逗号分隔")
        flag.StringVar(&exclude, "exclude", "", "要排除的子目录，用逗号分隔")
        flag.Parse()

        if watchPath == "" {
                fmt.Println("请使用 -path 指定需要监控的目录或文件路径")
                os.Exit(1)
        }

        //解析路径
        watchPaths := strings.Split(watchPath, ",")

        // 解析排除路径
        excludePaths := []string{}
        if exclude != "" {
                excludePaths = strings.Split(exclude, ",")
        }

        // 初始化 fsnotify
        watcher, err := fsnotify.NewWatcher()
        if err != nil {
                fmt.Printf("无法初始化文件监控器：\n", err)
                os.Exit(1)
        }
        defer watcher.Close()
        var filesHash = make(map[string]string)

        for _, path := range watchPaths {
                if info, err := os.Stat(path); err == nil {
                        if info.IsDir() {
                                // 添加监控对象
                                err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
                                        if err != nil {
                                                return err
                                        }
                                        if shouldExclude(path, excludePaths) {
                                                return nil
                                        }
                                        if info.IsDir() {
                                                err := watcher.Add(path)
                                                if err != nil {
                                                        fmt.Printf("无法添加监控目录 %s: %s\n", path, err)
                                                }
                                        } else {
                                                hash, err := calculateHash(path)
                                                if err != nil {
                                                        fmt.Printf("计算哈希出错, %v\n", err)
                                                }
                                                filesHash[path] = hash
                                        }
                                        return nil
                                })
                                if err != nil {
                                        fmt.Printf("监控初始化失败：%v\n", err)
                                        os.Exit(1)
                                }
                        } else {
                                hash, err := calculateHash(path)
                                if err != nil {
                                        fmt.Printf("计算哈希出错, %v\n", err)
                                }
                                filesHash[path] = hash
                                err = watcher.Add(path)
                                if err != nil {
                                        fmt.Printf("无法添加监控文件 %s: %v\n", path, err)
                                }
                        }
                }
        }

        fmt.Printf("开始监控：%v\n", watchPaths)

        // 监控事件
        go func() {
                for {
                        select {
                        case event, ok := <-watcher.Events:
                                if !ok {
                                        return
                                }
                                if _, ok := filesHash[event.Name]; !ok {
                                        continue
                                }
                                if event.Op&fsnotify.Chmod == fsnotify.Chmod {
                                        continue
                                }
                                // 处理文件事件
                                //if event.Op&fsnotify.Create == fsnotify.Create {
                                //      continue
                                //}
                                if shouldExclude(event.Name, excludePaths) {
                                        continue
                                }
                                if _, ok := filesHash[event.Name]; !ok {
                                        continue
                                }

                                if event.Has(fsnotify.Write) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) || event.Has(fsnotify.Create){
                                        hash, err := calculateHash(event.Name)
                                        if err != nil {
                                                fmt.Printf("%s产生事件，但计算哈希出错, %v\n", event.Name, err)
                                        }
                                        if filesHash[event.Name] != hash {
                                                fmt.Printf("文件%s哈希异常，请确认.\n", event.Name)
                                                os.Exit(1)
                                        }
                                }

                                //if event.Op&fsnotify.Write == fsnotify.Write {
                                //      log.Printf("Modify：%s\n", event.Name)
                                //}
                                //if event.Op&fsnotify.Remove == fsnotify.Remove {
                                //      log.Printf("Delete：%s\n", event.Name)
                                //}
                                //if event.Op&fsnotify.Rename == fsnotify.Rename {
                                //      log.Printf("Rename：%s\n", event.Name)
                                //}

                        case err, ok := <-watcher.Errors:
                                if !ok {
                                        return
                                }
                                fmt.Println("监控器错误：", err)
                        }
                }
        }()

        // 阻塞主线程
        done := make(chan bool)
        <-done
        os.Exit(0)
}

func calculateHash(filePath string) (string, error) {
        file, err := os.Open(filePath)
        if err != nil {
                return "", err
        }
        defer file.Close()

        hasher := sha256.New()
        if _, err := io.Copy(hasher, file); err != nil {
                return "", err
        }

        return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
