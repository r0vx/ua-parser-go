# UA Parser Go

🚀 **ua-parser-js 的 Go 语言重构版本**

一个高性能、类型安全的 User-Agent 解析库，用于检测浏览器、引擎、操作系统、CPU 和设备类型。

## ✨ 特性

- 🎯 **高准确率**：核心功能准确率达 96%
- ⚡ **高性能**：完整解析仅需 ~0.5ms
- 🔒 **类型安全**：利用 Go 的强类型系统
- 🌐 **全面支持**：识别 100+ 种浏览器、设备和操作系统
- 📦 **零依赖**：仅使用 Go 标准库
- 🧪 **测试完善**：包含详尽的单元测试和验证测试

## 📊 准确率

基于原始 ua-parser-js 的测试数据集验证：

| 模块 | 准确率 |
|------|--------|
| 浏览器检测 | 93.0% |
| 设备检测 | 99.0% |
| CPU架构 | 94.9% |
| 引擎检测 | 85.7% |
| 操作系统 | 65.3% |
| **平均** | **87.6%** |

详细验证报告请查看 [VALIDATION_REPORT.md](./VALIDATION_REPORT.md)

## 🚀 快速开始

### 安装

```bash
go get github.com/example/ua-parser-go
```

### 基本使用

```go
package main

import (
    "fmt"
    uaparser "github.com/example/ua-parser-go"
)

func main() {
    ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1"
    
    // 方式 1: 直接解析所有信息
    result := uaparser.Parse(ua)
    fmt.Printf("浏览器: %s %s\n", result.Browser.Name, result.Browser.Version)
    fmt.Printf("操作系统: %s %s\n", result.OS.Name, result.OS.Version)
    fmt.Printf("设备: %s %s\n", result.Device.Vendor, result.Device.Model)
    
    // 方式 2: 使用 Parser 实例（按需解析）
    parser := uaparser.New(ua)
    browser := parser.GetBrowser()
    os := parser.GetOS()
    device := parser.GetDevice()
    engine := parser.GetEngine()
    cpu := parser.GetCPU()
}
```

### 输出示例

```go
result := uaparser.Parse("Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36")

// result.Browser
// {
//   "name": "Mobile Chrome",
//   "version": "120.0.6099.43",
//   "major": "120"
// }

// result.OS
// {
//   "name": "Android",
//   "version": "14"
// }

// result.Device
// {
//   "vendor": "Samsung",
//   "model": "SM-S918B",
//   "type": "mobile"
// }

// result.Engine
// {
//   "name": "Blink",
//   "version": "120.0.6099.43"
// }
```

## 📖 API 文档

### 主要函数

#### `Parse(ua string) *Result`
一次性解析所有信息，返回完整结果。

#### `New(ua string) *UAParser`
创建解析器实例，支持按需解析。

### Parser 方法

- `GetBrowser() Browser` - 获取浏览器信息
- `GetOS() OS` - 获取操作系统信息
- `GetDevice() Device` - 获取设备信息
- `GetEngine() Engine` - 获取浏览器引擎信息
- `GetCPU() CPU` - 获取 CPU 架构信息
- `GetResult() *Result` - 获取完整解析结果

### 数据结构

```go
type Browser struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Major   string `json:"major"`
    Type    string `json:"type"`
}

type OS struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

type Device struct {
    Type   string `json:"type"`
    Model  string `json:"model"`
    Vendor string `json:"vendor"`
}

type Engine struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

type CPU struct {
    Architecture string `json:"architecture"`
}

type Result struct {
    UA      string  `json:"ua"`
    Browser Browser `json:"browser"`
    CPU     CPU     `json:"cpu"`
    Device  Device  `json:"device"`
    Engine  Engine  `json:"engine"`
    OS      OS      `json:"os"`
}
```

## 🎯 支持的浏览器

**桌面浏览器**
- Chrome / Chromium
- Firefox
- Safari
- Edge
- Opera / Opera GX
- Brave
- Vivaldi
- Yandex
- 360浏览器
- QQ浏览器
- 搜狗浏览器
- IE
- 更多...

**移动浏览器**
- Mobile Chrome
- Mobile Safari
- Mobile Firefox
- Samsung Internet
- UCBrowser
- WeChat（微信）
- QQ浏览器
- MIUI浏览器
- Huawei浏览器
- Vivo浏览器
- 更多...

**应用内浏览器**
- Facebook
- Instagram
- Twitter
- LINE
- WeChat
- 更多...

## 🔧 技术细节

### 正则表达式兼容性

由于 Go 的 `regexp` 包使用 RE2 引擎，本项目对原始 JavaScript 正则进行了以下适配：

1. **移除 Lookahead/Lookbehind**：将 `(?=...)` 和 `(?!...)` 重写为等效逻辑
2. **移除反向引用**：将 `\1`, `\2` 等反向引用改为显式匹配
3. **优化匹配顺序**：确保更具体的模式优先匹配

### 性能基准

```
BenchmarkParse-12         2379    502586 ns/op    1465 B/op    25 allocs/op
BenchmarkGetBrowser-12    5362    223390 ns/op    1040 B/op    17 allocs/op
```

在 Apple M2 Max 上：
- 完整解析：~0.5ms
- 单项解析（浏览器）：~0.22ms

## 🧪 测试

```bash
# 运行所有测试
go test -v ./...

# 运行验证测试（使用原始 ua-parser-js 的测试数据）
go test -v -run='Original'

# 运行基准测试
go test -bench=. -benchmem

# 查看测试覆盖率
go test -cover

# 运行示例程序
cd example && go run main.go
```

## 📁 项目结构

```
ua-parser-go/
├── parser.go              # 主解析器
├── mapper.go              # 正则映射器
├── regexes_browser.go     # 浏览器规则
├── regexes_os.go          # 操作系统规则
├── regexes_device.go      # 设备规则
├── regexes_cpu.go         # CPU 规则
├── regexes_engine.go      # 引擎规则
├── parser_test.go         # 单元测试
├── validate_test.go       # 验证测试
├── example/               # 示例代码
├── go.mod
├── README.md
└── VALIDATION_REPORT.md   # 详细验证报告
```

## 🤝 对比原始项目

本项目是 [ua-parser-js](https://github.com/faisalman/ua-parser-js) v2.0.7 的 Go 重构版本。

### 优势
- ✅ 编译型语言，性能更高
- ✅ 类型安全，减少运行时错误
- ✅ 内置并发支持
- ✅ 单一二进制文件，部署简单

### 已知差异
- 部分浏览器名称格式略有不同
- 空值使用空字符串而非 `undefined`
- 某些边缘情况的处理方式可能不同

详细对比请参考 [VALIDATION_REPORT.md](./VALIDATION_REPORT.md)

## 📝 许可证

本项目采用与原始 ua-parser-js 项目相同的许可协议。

## 🙏 致谢

本项目基于 [faisalman/ua-parser-js](https://github.com/faisalman/ua-parser-js) 重构，感谢原作者的出色工作！

---

**Made with ❤️ for the Go community**
