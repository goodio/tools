# env

自动设置环境变量或者获取配置文件内容（配置文件默认为 .env）

## 使用
```bash
// .env 文件内容
HOST=localhost
```
加载到环境变量
```golang

func main() {
  _, err := env.Load()
  if err != nil {
    fmt.Println(err)
  }
  
  // 获取.env文件HOST的内容
  host := os.Getenv("HOST")
  fmt.Println(host)
  
  // 获取系统PATH
  path := os.Getenv("PATH")
  
  fmt.Println(path)
  
}
```


