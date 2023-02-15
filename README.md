## Go标准应用模板项目

### 依赖包
```bash
# swagger :undo:
go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# LICENSE 文件
go install github.com/nishanths/license/v5@latest
license -list # 查看支持的代码协议
$ license -n 'Illus1onnn <1161524405@qq.com>' -o LICENSE mit # 项目根目录下执行

# Gin Http 框架
go get github.com/gin-gonic/gin
# UUID
go get github.com/google/uuid
# Cobra 命令行工具
github.com/spf13/cobra v1.6.1
# Viper 配置
github.com/spf13/viper v1.15.0
```

### 编译和运行 Zcar 应用
```bash
# windows
go build -o _output/zcar.exe -v cmd/zcar/main.go
# linux
go build -o _output/zcar -v cmd/zcar/main.go

# 运行
# windows
go run .\\cmd\\zcar\\main.go --config configs/
```